package adapter

import (
	"context"
	"fmt"
	"forkTravel/src/utils/database"
	_ "github.com/ClickHouse/clickhouse-go/v2"
	"log"
	"strings"
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

type Result struct {
	First         string `json:"first"`
	FirstCountry  string `json:"firstCountry"`
	Second        string `json:"second"`
	SecondCountry string `json:"secondCountry"`
	Third         string `json:"third"`
	ThirdCountry  string `json:"thirdCountry"`
	Fourth        string `json:"fourth"`
	FourthCountry string `json:"fourthCountry"`
}

type UserDatabase struct {
	database *database.DBConnect
}

func CreateUserDatabaseAdapter(database *database.DBConnect) *UserDatabase {
	adapter := &UserDatabase{database: database}
	return adapter
}

func (adapter *UserDatabase) GetAllTours(UserFromCountry, UserPreferences, UserFilters string) (results []*Result) {
	ctx := context.Background()
	results = make([]*Result, 0)

	if !strings.Contains(UserPreferences, ",") {
		add_filter := ""
		if UserFilters == "visa" {
			add_filter = " AND visa == 'Безвизовая'"
		}
		rows, err := adapter.database.Connection.Query(ctx, fmt.Sprintf("SELECT %v,countries FROM projects.countries WHERE isNotNull(%v)%v", UserPreferences, UserPreferences, add_filter))
		if err != nil {
			log.Fatal(err)
		}

		defer rows.Close()

		for rows.Next() {
			result := &Result{}
			if err := rows.Scan(&result.First, &result.FirstCountry); err != nil {
				log.Fatal(err)

			}
			results = append(results, result)
		}

		return
	} else {
		preferences := strings.Split(UserPreferences, ",")
		query := fmt.Sprintf("SELECT %v, countries FROM projects.countries WHERE isNotNull(%v)", preferences[0], preferences[0])
		rows1, err := adapter.database.Connection.Query(ctx, query)
		if err != nil {
			log.Fatal(err)
		}

		defer rows1.Close()

		for rows1.Next() {
			result := &Result{}
			if err := rows1.Scan(&result.First, &result.FirstCountry); err != nil {
				log.Fatal(err)
			}

			querySecond := fmt.Sprintf("SELECT %v, countries FROM projects.countries WHERE isNotNull(%v)", preferences[1], preferences[1])
			rows2, err := adapter.database.Connection.Query(ctx, querySecond)
			if err != nil {
				log.Fatal(err)
			}

			defer rows2.Close()

			for rows2.Next() {
				tempResult := *result
				if err := rows2.Scan(&tempResult.Second, &tempResult.SecondCountry); err != nil {
					log.Fatal(err)
				}

				results = append(results, &tempResult)
			}

		}

		//query := fmt.Sprintf("SELECT %v, countries FROM projects.countries WHERE isNotNull(%v)", preferences[0], preferences[0])
		//fmt.Println("query", query)
		//rows, err := adapter.database.Connection.Query(ctx, query)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//defer rows.Close()
		//
		//for rows.Next() {
		//	result := &Result{}
		//	values := make([]interface{}, len(preferences))
		//	pointers := make([]interface{}, len(preferences))
		//
		//	for i := range values {
		//		pointers[i] = &values[i]
		//	}
		//
		//	if err := rows.Scan(pointers...); err != nil {
		//		log.Fatal(err)
		//	}
		//
		//	for i, pref := range preferences {
		//		switch pref {
		//		case "mountain":
		//			result.First = values[i].(string)
		//		case "sea":
		//			result.Second = values[i].(string)
		//		case "excursion":
		//			result.Third = values[i].(string)
		//		case "health":
		//			result.Fourth = values[i].(string)
		//		}
		//	}
		//
		//	result.FirstCountry = values[len(values)-1].(string)
		//	results = append(results, result)
		//}

		return

		//tour := strings.Split(UserPreferences, ",")
		//fmt.Println("tour", tour)
		//fmt.Println("len", len(tour))
		//
		//return
	}

	//rows, err := adapter.database.Connection.Query(ctx, "SELECT * FROM projects.countries WHERE mountain != 'Нет'")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//defer rows.Close()
	//
	//tours = make([]*Tour, 0)
	//
	//for rows.Next() {
	//	tour := &Tour{}
	//	if err := rows.Scan(&tour.Countries, &tour.Mountain, &tour.Sea, &tour.Excursion, &tour.Health, &tour.Visa, &tour.Continent, &tour.Info); err != nil {
	//		log.Fatal(err)
	//
	//	}
	//	tour.Countries = ""
	//	tours = append(tours, tour)
	//	fmt.Println("====================================")
	//}

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

	//return
}
