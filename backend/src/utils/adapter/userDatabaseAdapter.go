package adapter

import (
	"database/sql"
	"forkTravel/src/utils/database"
	_ "github.com/ClickHouse/clickhouse-go"
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
	rows, err := adapter.database.Connection.Query("SELECT countries, mountain, sea, excursion, health FROM tour.countries")
	if err != nil {
		return nil, err
	}

	defer func(rows *sql.Rows) { _ = rows.Close() }(rows)

	tours = make([]*Tour, 0)

	for rows.Next() {
		tour := &Tour{}
		err = rows.Scan(&tour.Countries, &tour.Mountain, &tour.Sea, &tour.Excursion, &tour.Health)
		if err != nil {
			return nil, err
		}

		tours = append(tours, tour)
	}

	return
}
