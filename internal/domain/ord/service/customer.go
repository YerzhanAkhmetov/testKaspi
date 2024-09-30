package service

import (
	model "TestBroker/internal/domain/ord/entity"
	"TestBroker/internal/domain/ord/repository"
)

type OrderServiceImpl struct {
	repo repository.OrderRepository
}

func NewOrderService(repo repository.OrderRepository) *OrderServiceImpl {
	return &OrderServiceImpl{repo: repo}
}

func (s *OrderServiceImpl) GetAllOrders() ([]model.Order, error) {
	return s.repo.GetOrders()
}
