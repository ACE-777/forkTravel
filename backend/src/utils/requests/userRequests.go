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
	_, err := w.Write([]byte("pong"))
	if err != nil {
		fmt.Println(err)
	}

	log.Println("pong")
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

	tmpl, err := template.ParseFiles("C:\\Users\\misha\\GolandProjects\\forkTravel\\frontend\\templates\\handler_get_tour.html")
	if err != nil {
		fmt.Println("errr", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//a := Test{One: "oNEEEEE"}
	//err = tmpl.Execute(w, string(usersJSON))
	err = tmpl.Execute(w, tours)
	if err != nil {
		log.Println(err)
	}

	//fmt.Fprintf(w, "<b>"+string(usersJSON)+"</b>")

	return
}
