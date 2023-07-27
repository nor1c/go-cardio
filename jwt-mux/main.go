package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	authController "github.com/nor1c/GoJWTMux/controllers/authController"
	productController "github.com/nor1c/GoJWTMux/controllers/productController"
	"github.com/nor1c/GoJWTMux/middlewares"
	"github.com/nor1c/GoJWTMux/models"
)

func main() {
	models.Connect()

	r := mux.NewRouter()

	r.HandleFunc("/login", authController.Login).Methods("POST")
	r.HandleFunc("/register", authController.Register).Methods("POST")
	r.HandleFunc("/logout", authController.Logout).Methods("GET")

	products := r.PathPrefix("/products").Subrouter()
	products.HandleFunc("", productController.GetProducts).Methods("GET")
	products.Use(middlewares.JWTMiddleware)

	log.Fatal(http.ListenAndServe(":8080", r))
}
