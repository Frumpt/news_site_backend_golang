package main

import "fmt"

var configDB string = "host=localhost Router=postgres password=passwordtest dbname=Todos port=5432 sslmode=disable TimeZone=Asia/Shanghai"

var input string

func main() {
	for input != "exit" {
		fmt.Scan(&input)

		switch input {
		case "create":
			createTable()
		case "drop":
			dropTable()
		case "exit":
			fmt.Println("Exit")
		default:
			fmt.Println("Wrong input")
		}
	}

}
