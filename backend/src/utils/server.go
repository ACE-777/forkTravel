package utils

import (
	"fmt"
	"forkTravel/src/utils/database"
	"forkTravel/src/utils/requests"
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

func (server *Server) prepare() {

	userServer := requests.NewUserServer(server.Database)

	userHome := http.HandlerFunc(userServer.HandlerHome)
	http.Handle("/home/", userHome)

	userGetTour := http.HandlerFunc(userServer.HandlerGetTours)
	http.Handle("/tours/", userGetTour)
}

func (server *Server) Start(port int) {
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
