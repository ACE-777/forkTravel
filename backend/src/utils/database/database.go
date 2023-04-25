package database

import (
	"database/sql"
	"fmt"
	_ "github.com/ClickHouse/clickhouse-go"
)

type DBConnect struct {
	Ip       string `json:"ip"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`

	Connection *sql.DB
}

func (client *DBConnect) Open() error {

	driver := "clickhouse"

	db, err := sql.Open(driver, fmt.Sprintf("tcp://%s:%s?username=%s&password%s&database%s", client.Ip, client.Port, client.User, client.Password, client.Database))
	if err != nil {
		fmt.Println(err)
	}

	client.Connection = db
	return nil
}

func (client *DBConnect) Close() {
	client.Connection.Close()
}
