package main

import "log"

func main() {
	store, err := NewPostgresStore()
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	server := NewAPIServer(":3000", store)
	server.Run()
}
