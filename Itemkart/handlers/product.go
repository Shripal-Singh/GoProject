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

type ProductHandler struct {
	IProduct interfaces.IProduct
}

func (ch *ProductHandler) GetProductByID() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.IProduct == nil {
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
		contact, err := ch.IProduct.Get(id)
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
func (ch *ProductHandler) CreateProduct() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.IProduct == nil {
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

		product := &models.Product{}
		err = json.Unmarshal(buf, product)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in body json format",
			})
			glog.Errorln(err)
			c.Abort()
			return
		}
		err = product.Validate()

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
		id, err := ch.IProduct.Create(product)
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
			"message": "Product successfully created",
		})
		glog.Info("Product successfully created:", id)
		c.Abort()
	}
}
func (ch *ProductHandler) DeleteProduct() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.IProduct == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ProductHandler or IContact is nil")
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
		result, err := ch.IProduct.Delete(id)
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
			"message": fmt.Sprint(result, " Product record deleted"),
		})
		glog.Info("Product successfully deleted:", result)
		c.Abort()
	}
}
func (ch *ProductHandler) UpdateProduct() func(*gin.Context) {
	return func(c *gin.Context) {
		if ch == nil || ch.IProduct == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "fail",
				"message": "Error in the handler",
			})
			glog.Errorln("ProductHandler or IContact is nil")
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

		product := &models.Product{}
		err = json.Unmarshal(buf, product)

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
		ids, err := ch.IProduct.Update(id, product)
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
			"message": "Product record successfully updated",
		})
		glog.Info("Product record successfully updated:", ids)
		c.Abort()
	}
}
