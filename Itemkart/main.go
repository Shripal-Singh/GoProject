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
)

var (
	DBCONNECTION string = "root:Mysql@123@tcp(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"

	//"host=localhost user=root password=Mysql@123 dbname=demo port=3306 sslmode=disable TimeZone=Asia/Shanghai"
	PORT string = ":50080"
)

func usage() {
	//fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n")
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
	if os.Getenv("PORT") != "" {
		PORT = os.Getenv("PORT")
	}
	if os.Getenv("DBCONNECTION") != "" {
		DBCONNECTION = os.Getenv("DBCONNECTION")
	}

	db, err := database.GetConnection(DBCONNECTION)
	//db, err := sql.Open("mysql", "root:Mysql@123@tcp(127.0.0.1:3306)/demo")
	//db.SetMaxOpenConns(10)
	//db.SetMaxIdleConns(5)
	if err != nil {
		glog.Fatal("Database Error:", err)
	}

	cdb := &database.ContactDB{Client: db}
	//cdb := &filedb.FileDB{}
	contactHandler := &h.CustmerHandler{ICustomer: cdb}

	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/customer/:id", contactHandler.GetCustomerByID())
	r.POST("/customer", contactHandler.CreateCustomer())
	r.POST("/customer/:id", contactHandler.UpdateCustomer())
	r.DELETE("/customer/:id", contactHandler.DeleteCustomer())
	r.Run(PORT)
	fmt.Println("This is a ItemKart service.....")
}