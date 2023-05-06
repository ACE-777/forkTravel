package adapter

import (
	"context"
	"fmt"
	"forkTravel/src/utils/database"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"log"
	"time"
)

type Tour struct {
	Countries string `json:"countries" db:"countries"`
	Mountain  string `json:"mountain" db:"mountain"`
	Sea       string `json:"sea" db:"sea"`
	Excursion string `json:"excursion" db:"excursion"`
	Health    string `json:"health" db:"health"`
}

type UserDatabase struct {
	database *database.DBConnect
}

func CreateUserDatabaseAdapter(database *database.DBConnect) *UserDatabase {
	adapter := &UserDatabase{database: database}
	return adapter
}

func (adapter *UserDatabase) GetAllTours() (tours []*Tour, err error) {
	ctx := context.Background()
	rows, err := adapter.database.Connection.Query(ctx, "SELECT countries, mountain, sea, excursion, health FROM tour.countries")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var version string
		var version2 string
		var version3 string
		var version4 time.Time
		var version5 int64
		if err := rows.Scan(&version, &version2, &version3, &version4, &version5); err != nil {
			log.Fatal(err)
		}
		fmt.Println(version, version2, version3, version4, version5)
		fmt.Println("====================================")
	}

	//rows, err := adapter.database.Connection.Query("SELECT countries, mountain, sea, excursion, health FROM tour.countries")
	//if err != nil {
	//	return nil, err
	//}
	//
	//defer func(rows *sql.Rows) { _ = rows.Close() }(rows)
	//
	//tours = make([]*Tour, 0)
	//
	//for rows.Next() {
	//	tour := &Tour{}
	//	err = rows.Scan(&tour.Countries, &tour.Mountain, &tour.Sea, &tour.Excursion, &tour.Health)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	tours = append(tours, tour)
	//}

	return
}
