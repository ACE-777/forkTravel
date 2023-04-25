package main

import (
	"flag"
	"fmt"
	"forkTravel/src/utils"
	"forkTravel/src/utils/database"
	_ "github.com/ClickHouse/clickhouse-go"
)

func getHosts() string {
	flagPort := flag.String("port", "8000", "a port address")
	flag.Parse()
	return *flagPort
}

func main() {
	fmt.Println("Start Service on 8000 port")

	database := database.DBConnect{Ip: "127.0.0.1", Port: "9000", Password: "", User: "default", Database: "tour"}

	err := database.Open()
	err = database.Connection.Ping()
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Printf("Can not connect to ClickHouse: %s:%s", database.Ip, database.Port)
		panic(err)
	}

	fmt.Printf("Success Connect to ClickHouse: %s:%s", database.Ip, database.Port)

	server := utils.New(&database)
	server.Start(8000)
}
