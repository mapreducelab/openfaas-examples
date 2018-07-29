package function

import (
	"fmt"
	"log"

	"database/sql"
	_ "github.com/lib/pq"
)

// Read more about postgres driver at https://godoc.org/github.com/lib/pq
// Handle a serverless request
func Handle(req []byte) string {
	const (
		host     = "postgresql"
		port     = 5432
		user     = "postgres"
		password = "spark123"
		dbname   = "postgres"
		sslmode  = "disable"
	)
	q := "SELECT name,age FROM \"user\""

	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%d", user, password, dbname, sslmode, host, port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(q)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var names []string
	var ages []int
	for rows.Next() {
		var name string
		var age int
		err := rows.Scan(&name, &age)
		if err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
		ages = append(ages, age)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	return fmt.Sprintf("Names: %v, Ages: %v", names, ages)
}
