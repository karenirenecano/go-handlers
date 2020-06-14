package main

import (
	"net/http"

	"github.com/karenirenecano/go-handlers/product"
)

const apiBasePath = "/api"

func main() {
	product.SetupRoutes(apiBasePath)
	http.ListenAndServe(":5000", nil)
}
