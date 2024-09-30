package repository

import model "TestBroker/internal/domain/ord/entity"

type OrderRepository interface {
	GetOrders() ([]model.Order, error)
}
