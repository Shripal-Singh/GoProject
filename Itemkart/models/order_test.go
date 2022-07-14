package models_test

import (
		"itemkart/models"
		"testing"
)

func TestValidateOrderPass(t *testing.T) {
	contact := &models.Order{ProductId: 101, CustomerId: 1001}
	err := contact.Validate()
	//time.Sleep(time.Second * 31)
	if err != nil {
		t.Error("Validate must not return an error")
		t.Fail()
	}
}

func TestValidateOrderFail1(t *testing.T) {
	contact := &models.Order{ProductId: 101}
	err := contact.Validate()
	if err == nil {
		t.Error("Validate must return an error")
		t.Fail()
	}
}
func TestValidateOrderFail2(t *testing.T) {
	contact := &models.Order{CustomerId: 1001}
	err := contact.Validate()
	if err == nil {
		t.Error("Validate must return an error")
		t.Fail()
	}
}

