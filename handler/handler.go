package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	dbConnection := openDbConnection()

	var route Route
	router := mux.NewRouter()
	route.MemberRouterHandling(router, dbConnection)

	http.ListenAndServe(":3000", router)
}
