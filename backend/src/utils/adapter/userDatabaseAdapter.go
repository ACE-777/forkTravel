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

type Output struct {
	From       string    `json:"from"`
	Preference []string  `json:"preference"`
	Filters    []string  `json:"filters"`
	Result     []*Result `json:"result"`
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

func (adapter *UserDatabase) getAllToursRecursive(ctx context.Context, preferences []string, currentIndex int, currentResult *Result, results *[]*Result) error {
	if currentIndex == len(preferences) {
		*results = append(*results, currentResult)
		return nil
	}

	preference := preferences[currentIndex]
	query := fmt.Sprintf("SELECT %v, countries FROM projects.countries WHERE isNotNull(%v)", preference, preference)

	rows, err := adapter.database.Connection.Query(ctx, query)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		result := *currentResult

		switch currentIndex {
		case 0:
			if err := rows.Scan(&result.First, &result.FirstCountry); err != nil {
				return err
			}
		case 1:
			if err := rows.Scan(&result.Second, &result.SecondCountry); err != nil {
				return err
			}
		case 2:
			if err := rows.Scan(&result.Third, &result.ThirdCountry); err != nil {
				return err
			}
		case 3:
			if err := rows.Scan(&result.Fourth, &result.FourthCountry); err != nil {
				return err
			}
		}

		err := adapter.getAllToursRecursive(ctx, preferences, currentIndex+1, &result, results)
		if err != nil {
			return err
		}
	}

	return nil
}

func (adapter *UserDatabase) GetAllTours(UserFromCountry, UserPreferences, UserFilters string) (output *Output) {
	ctx := context.Background()
	results := make([]*Result, 0)
	if !strings.Contains(UserPreferences, ",") {
		add_filter := ""
		if UserFilters == "without_visa" {
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

		output = &Output{}
		output.From = UserFromCountry
		output.Preference = replaceOnUsefulValues(strings.Split(UserPreferences, ","))
		filters := strings.Split(UserFilters, ",")
		output.Filters = filters
		output.Result = results

		return
	} else {
		preferences := strings.Split(UserPreferences, ",")
		currentResult := &Result{}
		err := adapter.getAllToursRecursive(ctx, preferences, 0, currentResult, &results)
		if err != nil {
			log.Fatal(err)
		}
	}

	output = &Output{}
	output.From = UserFromCountry
	output.Preference = replaceOnUsefulValues(strings.Split(UserPreferences, ","))
	filters := strings.Split(UserFilters, ",")
	output.Filters = filters
	output.Result = results
	return
}

func replaceOnUsefulValues(preferences []string) []string {
	for i := range preferences {
		preferences[i] = strings.ReplaceAll(preferences[i], "sea", "Пляжный отдых")
		preferences[i] = strings.ReplaceAll(preferences[i], "excursion", "Исторический центр")
		preferences[i] = strings.ReplaceAll(preferences[i], "mountain", "Горнолыжный курорт")
		preferences[i] = strings.ReplaceAll(preferences[i], "health", "Здоровый отдых")
	}

	return preferences
}
