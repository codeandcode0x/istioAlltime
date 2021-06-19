package service

import (
	"ticket-manager/model"
)

type OrderService struct {
	DAO model.OrderDAO
}

/**
var dao model.OrderDAO
func init() {
	var m model.BaseModel
	dao = &model.Order{BaseModel: m}
}
*/

// get order service
func (u *OrderService) getSvc() *OrderService {
	var m model.BaseModel
	return &OrderService{
		DAO: &model.Order{BaseModel: m},
	}
}

// find all orders
func (u *OrderService) FindAllOrders() (interface{}, error) {
	var orders []model.Order
	err := u.getSvc().DAO.FindAll(&orders)
	return orders, err
}

// create order
func (u *OrderService) CreateOrder(order *model.Order) error {
	err := u.getSvc().DAO.Create(order)
	return err
}

// create order
func (u *OrderService) FindOrderByPages(currentPage, pageSize int) ([]model.Order, error) {
	var orders []model.Order
	err := u.getSvc().DAO.FindByPages(&orders, currentPage, pageSize)
	return orders, err
}
