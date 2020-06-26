package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/karenirenecano/go-handlers/database"
	"github.com/karenirenecano/go-handlers/product"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
