package models

import "fmt"

type Order struct {
	ID         uint `json:"id"`
	ProductId  uint `json:"productid"`
	CustomerId uint `json:"customerid"`
}

func (c *Order) Validate() error {
	if c.ProductId == 0 {
		return fmt.Errorf("invalid data for the filed:ProductId")
	}
	if c.CustomerId == 0 {
		return fmt.Errorf("invalid data for the filed:CustomerId")
	}
	
	return nil
}

