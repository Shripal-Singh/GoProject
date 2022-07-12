package interfaces

import "itemkart/models"

type IOrder interface {
	Create(*models.Order) (interface{}, error)
	Get(string) (*models.Order, error)
	Update(string, *models.Order) (interface{}, error)
	Delete(string) (interface{}, error)
}
