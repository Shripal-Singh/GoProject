package database

import (
	"itemkart/models"

	"gorm.io/gorm"
)

type ContactDB struct {
	Client interface{}
}

func (c *ContactDB) Get(id string) (*models.Customer, error) {
	contact := &models.Customer{}
	result := c.Client.(*gorm.DB).First(contact, id)

	if result.Error != nil {
		return nil, result.Error
	}
	return contact, nil
}
func (c *ContactDB) Create(contact *models.Customer) (interface{}, error) {
	c.Client.(*gorm.DB).AutoMigrate(&models.Customer{})
	//if err := c.IfExists(contact.Email); err != nil {
	//	return nil, err
	//}
	result := c.Client.(*gorm.DB).Create(contact)
	if result.Error != nil {
		return nil, result.Error
	}
	return contact.ID, nil
}
func (c *ContactDB) Delete(id string) (interface{}, error) {
	result := c.Client.(*gorm.DB).Delete(&models.Customer{}, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return result.RowsAffected, nil
}
func (c *ContactDB) Update(id string, contact *models.Customer) (interface{}, error) {
	act := &models.Customer{}
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
	if contact.Email != "" {
		act.Email = contact.Email
	}
	if contact.Address != "" {
		act.Address = contact.Address
	}
	result1 := c.Client.(*gorm.DB).Save(&act) //.Where("id = ?", id).Save(&act)

	if result1.Error != nil {
		return nil, result1.Error
	}
	return id, nil
}
