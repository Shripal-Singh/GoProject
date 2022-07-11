package main

import (
	"flag"
	"fmt"
	"itemkart/database"
	h "itemkart/handlers"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang/glog"
	"github.com/joho/godotenv"
)

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}

func init() {
	flag.Usage = usage
	// NOTE: This next line is key you have to call flag.Parse() for the command line
	// options or "flags" that are defined in the glog module to be picked up.
	flag.Parse()
}
func main() {
	err := godotenv.Load("localconfig.env")
	if err != nil {
		glog.Fatal("DSome error occured to load config file :", err)
	}

	DBCONNECTION := os.Getenv("DBCONNECTION")
	PORT := os.Getenv("PORT")

	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}
	if os.Getenv("DBCONNECTION") != "" {
		DBCONNECTION = os.Getenv("DBCONNECTION")
	}

	db, err := database.GetConnection(DBCONNECTION)
	if err != nil {
		glog.Fatal("Database Error:", err)
	}

	//Data store in File
	//cdb := &filedb.FileDB{}

	//Test API
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Customer API
	cdb := &database.ContactDB{Client: db}
	contactHandler := &h.CustmerHandler{ICustomer: cdb}

	r.GET("/customer/:id", contactHandler.GetCustomerByID())
	r.POST("/customer", contactHandler.CreateCustomer())
	r.POST("/customer/:id", contactHandler.UpdateCustomer())
	r.DELETE("/customer/:id", contactHandler.DeleteCustomer())

	//Product API
	cdbp := &database.ContactDBP{Client: db}
	ProductHandler := &h.ProductHandler{IProduct: cdbp}

	r.GET("/product/:id", ProductHandler.GetProductByID())
	r.POST("/product", ProductHandler.CreateProduct())
	r.POST("/product/:id", ProductHandler.UpdateProduct())
	r.DELETE("/product/:id", ProductHandler.DeleteProduct())

	r.Run(PORT)
	fmt.Println("This is a ItemKart service.....")
}
