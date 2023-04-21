package main

import (
	"database/sql"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"

	_ "github.com/ClickHouse/clickhouse-go"
)

func getHosts() string {
	flagPort := flag.String("port", "8000", "a port address")
	flag.Parse()
	return *flagPort
}

type Tour struct {
	ID      string `json:"id"`
	Country string `json:"country"`
	Sea     string `json:"sea"`
	Ex      string `json:"ex"`
	Health  string `json:"health"`
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("pong"))
	if err != nil {
		fmt.Println(err)
	}

	log.Println("pong")
	return
}

func handlerGetTour(w http.ResponseWriter, r *http.Request) {
	nameOfCountry := r.URL.Query().Get("name")
	fmt.Println(nameOfCountry)

	dsn := "tcp://localhost:9000?username=default&password=&database=tour"
	conn, err := sql.Open("clickhouse", dsn)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	rows, err := conn.Query("SELECT countries, moun,sea,excursione,health FROM tour.table_name")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []Tour

	for rows.Next() {
		var user Tour
		err = rows.Scan(&user.ID, &user.Country, &user.Sea, &user.Ex, &user.Health)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	usersJSON, err := json.Marshal(users)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(usersJSON)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func main() {
	http.HandleFunc("/", handlerPing)

	http.HandleFunc("/tour/", handlerGetTour)

	log.Fatal(http.ListenAndServe(":"+getHosts(), nil))
}
