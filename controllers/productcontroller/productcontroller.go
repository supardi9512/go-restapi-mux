package productcontroller

import (
	"encoding/json"
	"fmt"
	"go-restapi-mux/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var products []models.Product

	if err := models.DB.Find(&products).Error; err != nil {
		fmt.Println(err)
		return
	}

	response, _ := json.Marshal(products)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
