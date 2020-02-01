package main

import (
	"alisa-millionaire-lite/server/middlewares"
	"alisa-millionaire-lite/server/routers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r = routers.GetRouters(r)
	r.Use(middlewares.Cors)

	http.Handle("/", r)
	http.ListenAndServe(":5000", nil)
}
