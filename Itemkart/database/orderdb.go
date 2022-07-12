package database

import (
	"itemkart/models"

	"gorm.io/gorm"
)

type ContactDBO struct {
	Client interface{}
}

func (c *ContactDBO) Get(id string) (*models.Order, error) {
	contact := &models.Order{}
	result := c.Client.(*gorm.DB).First(contact, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}
func (c *ContactDBO) Create(contact *models.Order) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&models.Order{})
	//if err := c.IfExists(contact.Email); err != nil {
	//	return nil, err
	//}
	result := c.Client.(*gorm.DB).Create(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact.ID, nil
}
func (c *ContactDBO) Delete(id string) (interface{}, error) {
	result := c.Client.(*gorm.DB).Delete(&models.Order{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}
func (c *ContactDBO) Update(id string, contact *models.Order) (interface{}, error) {
	act := &models.Order{}
	result := c.Client.(*gorm.DB).First(act, id)

	if result.Error != nil {
		return nil, result.Error
	}
	// c.Client.(*gorm.DB).AutoMigrate(&models.Customer{})

	// act := models.Customer{}
	// i, err := strconv.Atoi(id)
	// if err != nil {
	// 	// ... handle error
	// 	panic(err)
	// }

	act.ProductId = contact.ProductId
	act.CustomerId = contact.CustomerId

	result1 := c.Client.(*gorm.DB).Save(&act) //.Where("id = ?", id).Save(&act)

	if result1.Error != nil {
		return nil, result1.Error
	}
	return id, nil
}
