package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var connStr = ""

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {

	connStr = os.Getenv("CONN_STR")

	// Open the database connection using the full connection string
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")

	for {

		fmt.Println("Printing users in 'Users' table:")

		// Query data from the users table
		rows, err := db.Query("SELECT id, name, age FROM users")
		if err != nil {
			log.Fatalf("Error querying database: %v", err)
		}

		// Iterate over the rows and print the results
		for rows.Next() {
			var user User
			err := rows.Scan(&user.ID, &user.Name, &user.Age)
			if err != nil {
				log.Fatalf("Error scanning row: %v", err)
			}
			fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
		}

		// Check for errors after iterating through the rows
		err = rows.Err()
		if err != nil {
			log.Fatalf("Error during row iteration: %v", err)
		}

		fmt.Println("")
		time.Sleep(time.Second * 3)
	}
}
