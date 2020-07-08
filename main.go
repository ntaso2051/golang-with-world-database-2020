package main

import (
	"fmt"
	"log"
	"os"

	_"github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type City struct {
	ID		int	`json:"id,otitempty" db:"ID"`
	Name		string	`json:"name,otitempty" db:"Name"`
	CountryCode	string	`json:"countryCode,otitempty" db:"CountryCode"`
	District	string	`json:"district,otitempty" db:"District"`
	Population	int	`json:"population,otitempty" bd:"Population"`
}

func main() {
	db, err := sqlx.Connect("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", os.Getenv("DB_USERNAME"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOSTNAME"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE")))
	if err != nil {
		log.Fatalf("Cannot Connect to Database: %s", err)
	}

	fmt.Println("Connected!")
	city := City{}
	nm := os.Args
	db.Get(&city, "SELECT * FROM city WHERE Name='%s'", nm[1])

	fmt.Printf("%sの人口は%d人です\n", nm[1], city.Population)
}
