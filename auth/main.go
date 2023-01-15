package main

import (
	"demo-k8s-auth/pkg/db"
	"log"
)

func main() {
	log.Println("Hello, world!")
	conn, err := db.Connect("192.168.1.100", "test", "test", "users", 5432)
	if err != nil {
		log.Println(err)
	}
	defer db.Close(conn)
	err = db.Ping(conn)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Connected to database successfully!")
	}
}
