package models

type Order struct {
	ID         uint `json:"id"`
	ProductId  uint `json:"productid"`
	CustomerId uint `json:"customerid"`
}
