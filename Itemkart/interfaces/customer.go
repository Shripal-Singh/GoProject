package interfaces

import "itemkart/models"

type ICustomer interface {
	Create(*models.Customer) (interface{}, error)
	Get(string) (*models.Customer, error)
	Update(string, *models.Customer) (interface{}, error)
	Delete(string) (interface{}, error)
}
