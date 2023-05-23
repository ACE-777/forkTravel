package main

import (
	"context"
	"flag"
	"fmt"
	"forkTravel/src/utils"
	"forkTravel/src/utils/database"
	"log"
)

func getHosts() int {
	flagPort := flag.Int("port", 8000, "a port address")
	flag.Parse()
	return *flagPort
}

func main() {
	fmt.Println("Start Service on 8000 port")

	database := database.DBConnect{Ip: "rc1a-23es9acuj4dk5eqe.mdb.yandexcloud.net", Port: 9440, Password: "72b3CywH", User: "misha777776", Database: "projects"}
	ctx := context.Background()

	err := database.Open()
	err = database.Connection.Ping(ctx)
	if err != nil {
		log.Printf("can not ping ClickHouse: %v", err)
		fmt.Println(err)
	}

	if err != nil {
		log.Printf("Can not connect to ClickHouse: %s:%v", database.Ip, database.Port)
		fmt.Printf("Can not connect to ClickHouse: %s:%v", database.Ip, database.Port)
		panic(err)
	}
	log.Printf("Success Connect to ClickHouse: %s:%v", database.Ip, database.Port)
	fmt.Printf("Success Connect to ClickHouse: %s:%v", database.Ip, database.Port)

	server := utils.New(&database)
	server.Start(getHosts())
}
