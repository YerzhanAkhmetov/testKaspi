package app

import (
	"TestBroker/internal/config"
	"TestBroker/internal/domain/ord/repository"
	repoCahe "TestBroker/internal/domain/ord/repository/cache"
	"TestBroker/internal/domain/ord/service"
	"TestBroker/internal/handler/http"
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
)

type App struct {
	httpServer      *http.Server
	ctx             context.Context
	ctxCancel       context.CancelFunc
	customerRepo    repository.OrderRepository
	customerService service.OrderService
	exitCode        int
}

func (a *App) Init() {

	a.ctx, a.ctxCancel = context.WithCancel(context.Background())

	// Инициализация репозиториев и сервисов
	a.customerRepo = repoCahe.NewFileOrderRepository(config.Conf.OrdersFilePath)
	a.customerService = service.NewOrderService(a.customerRepo)

	// Инициализация HTTP хендлеров
	orderHandler := handler.NewOrderHandler(a.customerService)

	// Инициализация HTTP сервера
	a.httpServer = &http.Server{
		Addr:              ":" + config.Conf.HttpPort,
		Handler:           a.setupRoutes(orderHandler),
		ReadHeaderTimeout: 2 * time.Second,
		ReadTimeout:       time.Minute,
		MaxHeaderBytes:    300 * 1024,
	}
}

func (a *App) setupRoutes(orderHandler *handler.OrderHandler) http.Handler {
	r := mux.NewRouter()

	// Определяем HTTP маршруты
	httpHandlers := []struct {
		method  string
		path    string
		handler http.HandlerFunc
	}{
		{
			method:  "GET",
			path:    "/orders",
			handler: orderHandler.GetOrders,
		},
		// Можем добавить дополнительные маршруты здесь
	}

	for _, h := range httpHandlers {
		r.HandleFunc(h.path, h.handler).Methods(h.method)
	}

	return r
}

func (a *App) Start() {
	slog.Info("Starting HTTP server...")
	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("http-server stopped", "error", err)
			a.exitCode = 1
		}
	}()
	slog.Info("HTTP server started at " + a.httpServer.Addr)
}
func (a *App) Listen() {
	signalCtx, signalCtxCancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer signalCtxCancel()

	// Ожидание сигнала завершения
	<-signalCtx.Done()
}
func (a *App) Stop() {
	slog.Info("Shutting down...")

	// stop context
	a.ctxCancel()

	// http-gw server
	{
		ctx, ctxCancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer ctxCancel()

		if err := a.httpServer.Shutdown(ctx); err != nil {
			slog.Error("http-server shutdown error", "error", err)
			a.exitCode = 1
		}
	}

}

func (a *App) Exit() {
	slog.Info("Exit")

	os.Exit(a.exitCode)

}
