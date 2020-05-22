package main

import (
	"github.com/tajud99n/bookstore_items-api/app"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func main() {
	app.StartApplication()
}
