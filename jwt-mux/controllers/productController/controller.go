package productcontroller

import (
	"net/http"

	"github.com/nor1c/GoJWTMux/helpers"
)

func GetProducts(w http.ResponseWriter, r *http.Request) {
	products := []map[string]interface{}{
		{
			"name": "Shoes",
			"qty":  200,
		},
		{
			"name": "T-shirt",
			"qty":  "1000",
		},
	}

	helpers.ResponseJSON(w, http.StatusOK, products)
}
