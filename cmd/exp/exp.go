package main

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	SSLMode  string
}

func (cfg PostgresConfig) String() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Database, cfg.SSLMode)
}
func main() {
	cfg := PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "baloo",
		Password: "junglebook",
		Database: "lenslocked",
		SSLMode:  "disable",
	}
	db, err := sql.Open("pgx", cfg.String())
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// Create a table...
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS users(
		id SERIAL PRIMARY KEY,
		name TEXT,
		email TEXT UNIQUE NOT NULL
	);
	
	
	CREATE TABLE IF NOT EXISTS orders(
		id SERIAL PRIMARY KEY,
		user_id INT NOT NULL,
		amount INT,
		description TEXT
	);`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Tables created.")

	// name := "test1"
	// email := "test1@test.com"
	// row := db.QueryRow(`
	// INSERT INTO users(name, email)
	// VALUES ($1,$2) RETURNING id;`, name, email)
	// var id int
	// err = row.Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("user created. id = ", id)

	id := 1
	row := db.QueryRow(`
	SELECT name, email
	FROM users
	where id=$1; `, id)
	var name, email string
	err = row.Scan(&name, &email)
	if err != nil {
		panic(err)
	}
	fmt.Printf("user info: name=%s, email=%s \n", name, email)

	userID := 1

	// for i := 1; i <= 5; i++ {
	// 	amount := i * 100
	// 	desc := fmt.Sprintf("fake order #%d", i)

	// 	_, err := db.Exec(`
	// 	INSERT INTO orders(user_id, amount, description)
	// 	VALUES($1, $2, $3)`, userID, amount, desc)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// }
	fmt.Println("created fake orders")

	type Order struct {
		ID          int
		UserID      int
		Amount      int
		description string
	}
	var orders []Order
	rows, err := db.Query(`
		SELECT  id, amount, description FROM orders WHERE user_id=$1`, userID)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var order Order
		order.UserID = userID
		err := rows.Scan(&order.ID, &order.Amount, &order.description)
		if err != nil {
			panic(err)
		}
		orders = append(orders, order)
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}
	fmt.Println("Orders: ", orders)
	// check for an error
}
