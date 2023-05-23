package utils

import (
	"fmt"
	"forkTravel/src/utils/database"
	"forkTravel/src/utils/requests"
	"log"
	"net/http"
)

type Server struct {
	Database *database.DBConnect
}

func New(database *database.DBConnect) *Server {
	webServer := &Server{Database: database}
	webServer.prepare()
	return webServer
}

var mux = http.NewServeMux()

func (server *Server) prepare() {

	userServer := requests.NewUserServer(server.Database)

	userHome := http.HandlerFunc(userServer.HandlerHome)
	mux.HandleFunc("/home/", userHome)

	userResult := http.HandlerFunc(userServer.HandlerResult)
	mux.HandleFunc("/result/", userResult)

	//userGetTour := http.HandlerFunc(userServer.HandlerGetTours)
	//mux.HandleFunc("/tours/", userGetTour)
}

func (server *Server) Start(port int) {
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
	log.Fatal(err)
}
