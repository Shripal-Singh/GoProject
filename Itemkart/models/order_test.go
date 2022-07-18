package models_test

import (
		"itemkart/models"
		"testing"
)

func TestOrder(t *testing.T) {
    type test struct {
        input models.Order
        want  string
    }

    tests := []test{
        {input: models.Order{ProductId: 101, CustomerId: 1001},want: "pass"},
		{input: models.Order{}, want: "fail"},
		{input: models.Order{ProductId: 101}, want: "fail"},
		{input: models.Order{CustomerId: 1001}, want: "fail"},
		     
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

// func TestValidateOrderPass(t *testing.T) {
// 	contact := &models.Order{ProductId: 101, CustomerId: 1001}
// 	err := contact.Validate()
// 	//time.Sleep(time.Second * 31)
// 	if err != nil {
// 		t.Error("Validate must not return an error")
// 		t.Fail()
// 	}
// }

// func TestValidateOrderFail1(t *testing.T) {
// 	contact := &models.Order{ProductId: 101}
// 	err := contact.Validate()
// 	if err == nil {
// 		t.Error("Validate must return an error")
// 		t.Fail()
// 	}
// }
// func TestValidateOrderFail2(t *testing.T) {
// 	contact := &models.Order{CustomerId: 1001}
// 	err := contact.Validate()
// 	if err == nil {
// 		t.Error("Validate must return an error")
// 		t.Fail()
// 	}
// }

