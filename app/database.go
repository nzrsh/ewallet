package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDB() error {
	var err error
	db, err = sql.Open("sqlite3", "./database/ewallet.db")
	if err != nil {
		log.Println("Ошибка при открытии существующей базы данных, создаем новую")
		file, err := os.Create("./database/ewallet.db")
		//Тут должна вызываться функция заполнения новой бд
		if err != nil {
			log.Println(err)
			return err
		}
		defer file.Close()

		db, err = sql.Open("sqlite3", "./database/ewallet.db")
		if err != nil {
			log.Println("Ошибка при открытии новой базы данных")
			return err
		}
	}
	log.Println(db)
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var address string
		var balance float64
		err := rows.Scan(&id, &address, &balance)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("ID: %d, Name: %s, Age: %f\n", id, address, balance)
	}
	return nil
}
