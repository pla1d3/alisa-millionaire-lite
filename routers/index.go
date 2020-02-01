package routers

import (
	"alisa-millionaire-lite/server/controllers/alisa"

	"github.com/gorilla/mux"
)

func GetRouters(routers *mux.Router) *mux.Router {
	routers.HandleFunc("/alisa", alisa.AlisaHook).Methods("POST")
	return routers
}
