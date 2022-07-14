package models_test

import (
	"itemkart/models"
	"testing"
)

func TestValidateProductPass(t *testing.T) {
	contact := &models.Product{Category: "redme", Description: "redme 9", Name: "mobile"}
	err := contact.Validate()
	if err != nil {
		t.Error("Validate must not return an error")
		t.Fail()
	}
}

func TestValidateProductFail(t *testing.T) {
	contact := &models.Product{Category: "redme", Description: "redme 9"}
	err := contact.Validate()
	if err == nil {
		t.Error("Validate must return an error")
		t.Fail()
	}
}
