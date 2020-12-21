package main

import (
	"flag"
	"fmt"
)

func main() {
	host := flag.String("host", "localhost", "Put your database host")
	username := flag.String("username", "root", "go-db")
	password := flag.String("", "root", "go-db-api")

	flag.Parse()

	fmt.Println(*host, *username, *password)
}
