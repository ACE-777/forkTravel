package main

import (
	"context"
	"flag"
	"fmt"
	"forkTravel/src/utils"
	"forkTravel/src/utils/database"
)

func getHosts() int {
	flagPort := flag.Int("port", 8000, "a port address")
	flag.Parse()
	return *flagPort
}

func main() {
	fmt.Println("Start Service on 8000 port")

	database := database.DBConnect{Ip: "", Port: 9440, Password: "", User: "", Database: "tour"}

	ctx := context.Background()

	err := database.Open()
	err = database.Connection.Ping(ctx)
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Printf("Can not connect to ClickHouse: %s:%s", database.Ip, database.Port)
		panic(err)
	}

	fmt.Printf("Success Connect to ClickHouse: %s:%s", database.Ip, database.Port)

	server := utils.New(&database)
	server.Start(getHosts())
}
