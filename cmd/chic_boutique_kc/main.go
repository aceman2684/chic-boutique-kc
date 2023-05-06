package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/ae054577/chic-boutique-kc/address"
	"github.com/ae054577/chic-boutique-kc/postgres"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	db := connectDatabase()
	defer db.Close()

	loadSQLfile := postgres.LoadSQLFile

	addressService := address.NewAddressService(db, loadSQLfile)
	address, err := addressService.CreateAddress("1234 Main St", nil, "Olathe", "Kansas", "66062")
	if err != nil {
		panic(err)
	}

	fmt.Println(address)

	router := initializeRouter()
	port := getPort()
	http.ListenAndServe(":"+port, router)
}

func connectDatabase() *sql.DB {
	db, err := postgres.Connect()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return db
}

func initializeRouter() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.SetHeader("Content-Type", "application/json"))

	return router
}

func getPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		fmt.Println("no PORT set")
		os.Exit(1)
	}

	return port
}
