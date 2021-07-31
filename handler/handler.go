package handler

import (
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {
	dbConnection := openDbConnection()

	var route Route
	router := mux.NewRouter()
	route.MemberRouterHandling(router, dbConnection)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000" // Default port if not specified
	}

	http.ListenAndServe(":"+port, router)
}
