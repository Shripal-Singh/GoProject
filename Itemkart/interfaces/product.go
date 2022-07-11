package interfaces

import "itemkart/models"

type IProduct interface {
	Create(*models.Product) (interface{}, error)
	Get(string) (*models.Product, error)
	Update(string, *models.Product) (interface{}, error)
	Delete(string) (interface{}, error)
}
