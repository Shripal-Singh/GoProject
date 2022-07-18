package models_test

import (
	"itemkart/models"
	"testing"
)

func TestSplitProduct(t *testing.T) {
    type test struct {
        input models.Product
        want  string
    }

    tests := []test{
        {input: models.Product{Category: "redme", Description: "redme 9", Name: "mobile"},want: "pass"},
		{input: models.Product{}, want: "fail"},
		{input: models.Product{Category: "redme"}, want: "fail"},
		{input: models.Product{Name: "Shri"}, want: "fail"},
		{input: models.Product{Description: "redme 9"}, want: "fail"},
        {input: models.Product{Category: "sredme", Description: "redme 9"}, want: "fail"},
		{input: models.Product{Category: "redme", Name: "Shri"}, want: "fail"},
		{input: models.Product{Description: "redme 9", Name: "Shri"}, want: "fail"},
	     
    }

    for _, tc := range tests {
		contact :=tc.input
         err := contact.Validate()
		 if tc.want=="pass" && err != nil {
			t.Error("Validate must not return an error")
			t.Fail()
		}
		if tc.want=="fail" && err == nil {
			t.Error("Validate must not return an error")
			t.Fail()
		}
    }
}

// func TestValidateProductPass(t *testing.T) {
// 	contact := &models.Product{Category: "redme", Description: "redme 9", Name: "mobile"}
// 	err := contact.Validate()
// 	if err != nil {
// 		t.Error("Validate must not return an error")
// 		t.Fail()
// 	}
// }

// func TestValidateProductFail(t *testing.T) {
// 	contact := &models.Product{Category: "redme", Description: "redme 9"}
// 	err := contact.Validate()
// 	if err == nil {
// 		t.Error("Validate must return an error")
// 		t.Fail()
// 	}
// }
