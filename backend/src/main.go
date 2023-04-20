package main

import (
	"database/sql"
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

func handlerPing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write([]byte("pong"))
	if err != nil {
		fmt.Println(err)
	}

	log.Println("pong")
	return
}

func main() {
	conn, err := sql.Open("clickhouse", "tcp://<127.0.0.1>:<8777>?username=<admin>&password=<password>&database=<tour>")
	if err != nil {
		fmt.Printf("can not connect to clickHouse: %v", err)
	}

	defer func() { _ = conn.Close() }()

	http.HandleFunc("/", handlerPing)

	log.Fatal(http.ListenAndServe(":"+getHosts(), nil))
}
