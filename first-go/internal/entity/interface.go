package entity

type OrderRepositoryInterface interface {
	add(order *Order) error
}