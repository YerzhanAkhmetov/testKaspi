package service

import model "TestBroker/internal/domain/ord/entity"

type OrderService interface {
	GetAllOrders() ([]model.Order, error)
}
