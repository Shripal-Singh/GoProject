package models

import "fmt"

type Customer struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name" gorm:"index"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

func (c *Customer) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("invalid data for the filed:Name")
	}
	if c.Email == "" {
		return fmt.Errorf("invalid data for the filed:Email")
	}
	if c.Address == "" {
		return fmt.Errorf("invalid data for the filed:Address")
	}
	return nil
}
