package handlers

import (
	"encoding/json"
	"fmt"
	"itemkart/interfaces"
	"itemkart/models"

	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type CustmerHandler struct {
	ICustomer interfaces.ICustomer
}

func (ch *CustmerHandler) GetCustomerByID() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.ICustomer == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ContactHandler or IContact is nil")
			c.Abort()
			return
		}

		id, ok := c.Params.Get("id")

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id parameter not found",
			})
			glog.Errorln("id parameter not found")
			c.Abort()
			return
		}

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in id param",
			})
			glog.Errorln("id cannot be empty")
			c.Abort()
			return
		}
		contact, err := ch.ICustomer.Get(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error in fetching contact",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if contact == nil {
			c.JSON(http.StatusNoContent, nil)
			glog.Info("No content")
			c.Abort()
			return
		}

		c.JSON(http.StatusOK, *contact)
		glog.Info("Contact successfully fetched:", *contact)
		c.Abort()
	}
}
func (ch *CustmerHandler) CreateCustomer() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.ICustomer == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ContactHandler or IContact is nil")
			c.Abort()
			return
		}

		buf, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the body",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}

		customer := &models.Customer{}
		err = json.Unmarshal(buf, customer)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in body json format",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		err = customer.Validate()

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": err.Error(),
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		// if err := ch.IContact.IfExists(contact.Email); err != nil {

		// 	if err != nil {
		// 		c.JSON(http.StatusBadRequest, models.Response{Status: "fail", Message: err.Error()})
		// 		glog.Errorln(err)
		// 		c.Abort()
		// 		return
		// 	}
		// }
		//contact.Status = "active"
		//contact.LastModified = time.Now().UTC().String()
		id, err := ch.ICustomer.Create(customer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error to store in database",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "Customer successfully created",
		})
		glog.Info("Customer successfully created:", id)
		c.Abort()
	}
}
func (ch *CustmerHandler) DeleteCustomer() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.ICustomer == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ContactHandler or IContact is nil")
			c.Abort()
			return
		}

		id, ok := c.Params.Get("id")

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id parameter not found",
			})
			glog.Errorln("id parameter not found")
			c.Abort()
			return
		}

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id cannot be empty",
			})
			glog.Errorln("id cannot be empty")
			c.Abort()
			return
		}
		result, err := ch.ICustomer.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "error deleting contact",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		if result.(int64) == 0 {
			c.JSON(http.StatusOK, gin.H{
				"status":  "success",
				"message": "no record to delete with the given id:" + id,
			})
			glog.Info("no record to delete with the given id:", id)
			c.Abort()
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": fmt.Sprint(result, " customer record deleted"),
		})
		glog.Info("Contact successfully deleted:", result)
		c.Abort()
	}
}
func (ch *CustmerHandler) UpdateCustomer() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.ICustomer == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ContactHandler or IContact is nil")
			c.Abort()
			return
		}

		buf, err := ioutil.ReadAll(c.Request.Body)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the body",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}

		customer := &models.Customer{}
		err = json.Unmarshal(buf, customer)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in body json format",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		// err = customer.Validate()

		// if err != nil {
		// 	c.JSON(http.StatusBadRequest, gin.H{
		// 		"status":  "fail",
		// 		"message": err.Error(),
		// 	})
		// 	glog.Errorln(err)
		// 	c.Abort()
		// 	return
		// }
		// if err := ch.IContact.IfExists(contact.Email); err != nil {

		// 	if err != nil {
		// 		c.JSON(http.StatusBadRequest, models.Response{Status: "fail", Message: err.Error()})
		// 		glog.Errorln(err)
		// 		c.Abort()
		// 		return
		// 	}
		// }
		//contact.Status = "active"
		//contact.LastModified = time.Now().UTC().String()
		id, ok := c.Params.Get("id")

		if !ok {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id parameter not found",
			})
			glog.Errorln("id parameter not found")
			c.Abort()
			return
		}

		if id == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "id cannot be empty",
			})
			glog.Errorln("id cannot be empty")
			c.Abort()
			return
		}
		ids, err := ch.ICustomer.Update(id, customer)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error to store in database",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"status":  "success",
			"message": "Customer record successfully updated",
		})
		glog.Info("Customer record successfully updated:", ids)
		c.Abort()
	}
}
