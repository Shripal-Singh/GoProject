package models

import "fmt"

type Product struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"index"`
	Category    string `json:"category"`
	Description string `json:"description"`
}

func (c *Product) Validate() error {
	if c.Name == "" {
		return fmt.Errorf("invalid data for the filed:Name")
	}
	if c.Category == "" {
		return fmt.Errorf("invalid data for the filed:Category")
	}
	if c.Description == "" {
		return fmt.Errorf("invalid data for the filed:Description")
	}
	return nil
}
