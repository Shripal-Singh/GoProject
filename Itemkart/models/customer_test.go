package models_test

import (
	"itemkart/models"
	"testing"
)

func TestValidateCustomerPass(t *testing.T) {
	contact := &models.Customer{Email: "shri@gmail.com", Address: "noida", Name: "shri"}
	err := contact.Validate()
	//time.Sleep(time.Second * 31)
	if err != nil {
		t.Error("Validate must not return an error")
		t.Fail()
	}
}

func TestValidateCustomerFail(t *testing.T) {
	contact := &models.Customer{Email: "shri@gmail.com", Address: "noida"}
	err := contact.Validate()
	if err == nil {
		t.Error("Validate must return an error")
		t.Fail()
	}
}
