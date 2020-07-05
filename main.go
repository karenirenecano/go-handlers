package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/karenirenecano/go-handlers/database"
	"github.com/karenirenecano/go-handlers/product"
	"github.com/karenirenecano/go-handlers/utils"
)

const apiBasePath = "/api"

func main() {
	utils.SetGlobalEnvVariables()
	database.SetupDatabase()
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
