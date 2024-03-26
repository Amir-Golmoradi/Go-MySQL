package routes

import (
	"github.com/Amir_Golmoradi/Go_MySQL/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")

}
