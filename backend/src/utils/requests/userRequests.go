package requests

import (
	"encoding/json"
	"fmt"
	"forkTravel/src/utils/adapter"
	"forkTravel/src/utils/database"
	"html/template"
	"log"
	"net/http"
)

type UserServer struct {
	Database *database.DBConnect
}

func NewUserServer(database *database.DBConnect) *UserServer {
	return &UserServer{Database: database}
}

func (server *UserServer) HandlerHome(w http.ResponseWriter, r *http.Request) {

	//w.Header().Set("Content-Type", "application/json")
	tmpl, err := template.ParseFiles("../frontend/templates/handler_home_tour.html")
	if err != nil {
		fmt.Println("errr", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
	return
}

func (server *UserServer) HandlerGetTours(w http.ResponseWriter, r *http.Request) {
	//nameOfCountry := r.URL.Query().Get("name")
	//fmt.Println(nameOfCountry)

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(server.Database)
	tours, err := userDatabaseAdapter.GetAllTours()
	if err != nil {
		fmt.Printf("Error in func GetALLTours: %v", err)
	}

	usersJSON, err := json.Marshal(tours)
	if err != nil {
		fmt.Printf("Error in marshalalling func GetALLTours: %v", err)
	}

	fmt.Println(string(usersJSON))

	tmpl, err := template.ParseFiles("../frontend/templates/handler_get_tour.html")
	if err != nil {
		fmt.Println("errr", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, tours)
	if err != nil {
		log.Println(err)
	}

	return
}
