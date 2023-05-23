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
		log.Printf("Can not parse template for home: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		log.Printf("Can not execute template: %v", err)
	}
	return
}

func (server *UserServer) HandlerResult(w http.ResponseWriter, r *http.Request) {
	UserFromCountry := r.URL.Query().Get("from")
	UserPreferences := r.URL.Query().Get("preferences")
	UserFilters := r.URL.Query().Get("filters")

	userDatabaseAdapter := adapter.CreateUserDatabaseAdapter(server.Database)
	result := userDatabaseAdapter.GetAllTours(UserFromCountry, UserPreferences, UserFilters)

	usersJSON, err := json.Marshal(result)
	if err != nil {
		fmt.Printf("Error in marshalalling func GetALLTours: %v", err)
	}

	fmt.Println("result", result)
	fmt.Println(string(usersJSON))

	tmpl, err := template.ParseFiles("../frontend/templates/handler_get_result.html")
	if err != nil {
		fmt.Println("errr", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, result)
	if err != nil {
		log.Println(err)
	}

	return
}
