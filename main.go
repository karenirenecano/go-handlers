package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

//Product struct
type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand string `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productJSON := `[
		{
			"productId": 123,
			"manufacturer": "Big Box Company",
			"sku": "4561qHJK",
			"upc": "41465444566",
			"pricePerUnit": "12.99",
			"quantityOnHand": "28",
			"productName": "Gizmo"
		},
		{
			"productId": 124,
			"manufacturer": "Company 2",
			"sku": "5555",
			"upc": "1233333",
			"pricePerUnit": "566",
			"quantityOnHand": "300",
			"productName": "Succulents"
		}
	]`

	err := json.Unmarshal([]byte(productJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

func getNextID() int {
	highestID := -1
	for _, product := range productList {
		if highestID < product.ProductID {
			highestID = product.ProductID
		}
	}

	return highestID + 1
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		productsJSON, err := json.Marshal(productList)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(productsJSON)

	case http.MethodPost:
		var newProduct Product
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = json.Unmarshal(bodyBytes, &newProduct)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if newProduct.ProductID != 0 {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		newProduct.ProductID = getNextID()
		productList = append(productList, newProduct)
		w.WriteHeader(http.StatusCreated)
		return
	}
}

func main() {
	http.HandleFunc("/products", productsHandler)
	http.ListenAndServe(":5000", nil)
}
