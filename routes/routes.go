package routes

import (
	"luiz/shop/controllers"
	"net/http"
)

func Router() {
	http.HandleFunc("/", controllers.Index)
}
