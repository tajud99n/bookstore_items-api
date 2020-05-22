package app

import (
	"net/http"

	"github.com/tajud99n/bookstore_items-api/controllers"
)

func mapURLs() {
	router.HandleFunc("/items", controllers.ItemsController.Create).Methods(http.MethodPost)

}
