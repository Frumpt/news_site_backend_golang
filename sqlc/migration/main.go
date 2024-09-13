package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	var input string
	for input != "exit" {
		fmt.Println("Welcome to migration service for Postgresql \n select :  \"up\"; \"down\"; \"exit\"")
		_, _ = fmt.Scan(&input)

		switch input {
		case "up":
			up()
		case "down":
			down()
		case "exit":
			fmt.Println("Exit")
		default:
			fmt.Println("Wrong input")
		}
	}
}

func up() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:passwordtest@localhost:5432/Todos?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}
}

func down() {
	m, err := migrate.New(
		"file://migrations",
		"postgres://postgres:passwordtest@localhost:5432/Todos?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Down(); err != nil {
		log.Fatal(err)
	}
}
