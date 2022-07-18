package models_test

import (
	"itemkart/models"
	"testing"
)

func TestCustomer(t *testing.T) {
    type test struct {
        input models.Customer
        want  string
    }

    tests := []test{
        {input: models.Customer{Email: "shri@gmail.com", Address: "noida", Name: "shri"},want: "pass"},
		{input: models.Customer{}, want: "fail"},
		{input: models.Customer{Email: "shri@gmail.com"}, want: "fail"},
		{input: models.Customer{Name: "Shri"}, want: "fail"},
		{input: models.Customer{Address: "Noida"}, want: "fail"},
        {input: models.Customer{Email: "shri@gmail.com", Address: "noida"}, want: "fail"},
		{input: models.Customer{Email: "shri@gmail.com", Name: "Shri"}, want: "fail"},
		{input: models.Customer{Address: "Noida", Name: "Shri"}, want: "fail"},
	     
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

// func TestValidateCustomerPass(t *testing.T) {
// 	contact := &models.Customer{Email: "shri@gmail.com", Address: "noida", Name: "shri"}
// 	err := contact.Validate()
// 	//time.Sleep(time.Second * 31)
// 	if err != nil {
// 		t.Error("Validate must not return an error")
// 		t.Fail()
// 	}
// }

// func TestValidateCustomerFail(t *testing.T) {
// 	contact := &models.Customer{Email: "shri@gmail.com", Address: "noida"}
// 	err := contact.Validate()
// 	if err == nil {
// 		t.Error("Validate must return an error")
// 		t.Fail()
// 	}
// }
