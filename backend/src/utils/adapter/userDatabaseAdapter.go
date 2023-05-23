package adapter

import (
	"context"
	"fmt"
	"forkTravel/src/utils/database"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"log"
)

type Tour struct {
	Countries string `json:"countries" db:"countries"`
	Mountain  string `json:"mountain" db:"mountain"`
	Sea       string `json:"sea" db:"sea"`
	Excursion string `json:"excursion" db:"excursion"`
	Health    string `json:"health" db:"health"`
	Visa      string `json:"visa" db:"visa"`
	Continent string `json:"continent" db:"continent"`
	Info      string `json:"info" db:"info"`
}

type UserDatabase struct {
	database *database.DBConnect
}

func CreateUserDatabaseAdapter(database *database.DBConnect) *UserDatabase {
	adapter := &UserDatabase{database: database}
	return adapter
}

func (adapter *UserDatabase) GetAllTours(UserFromCountry, UserPreferences, UserFilters string) (tours []*Tour, err error) {
	ctx := context.Background()
	rows, err := adapter.database.Connection.Query(ctx, "SELECT * FROM projects.countries")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	tours = make([]*Tour, 0)

	for rows.Next() {
		tour := &Tour{}
		if err := rows.Scan(&tour.Countries, &tour.Mountain, &tour.Sea, &tour.Excursion, &tour.Health, &tour.Visa, &tour.Continent, &tour.Info); err != nil {
			log.Fatal(err)
		}

		tours = append(tours, tour)
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
