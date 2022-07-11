package database

import (
	"itemkart/models"

	"gorm.io/gorm"
)

type ContactDBP struct {
	Client interface{}
}

func (c *ContactDBP) Get(id string) (*models.Product, error) {
	contact := &models.Product{}
	result := c.Client.(*gorm.DB).First(contact, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}
func (c *ContactDBP) Create(contact *models.Product) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&models.Product{})
	//if err := c.IfExists(contact.Email); err != nil {
	//	return nil, err
	//}
	result := c.Client.(*gorm.DB).Create(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact.ID, nil
}
func (c *ContactDBP) Delete(id string) (interface{}, error) {
	result := c.Client.(*gorm.DB).Delete(&models.Product{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}
func (c *ContactDBP) Update(id string, contact *models.Product) (interface{}, error) {
	act := &models.Product{}
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

	// act.ID = uint(i)
	act.Name = contact.Name
	if contact.Description != "" {
		act.Description = contact.Description
	}
	if contact.Category != "" {
		act.Category = contact.Category
	}
	result1 := c.Client.(*gorm.DB).Save(&act) //.Where("id = ?", id).Save(&act)

	if result1.Error != nil {
		return nil, result1.Error
	}
	return id, nil
}
