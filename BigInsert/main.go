package main

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/TheGolurk/BigInsert/random"
	_ "github.com/go-sql-driver/mysql" //mysql driver
)

func main() {

	db, err := sql.Open("mysql", "user_area:user1area2#@(mysql.area.ingenieriaacustica.cl)/db_area?parseTime=true")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	db.SetConnMaxLifetime(0)

	file, err := os.Open("datos.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)
	index := 1

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("ERROR")
			return
		}

		fmt.Printf("Row -> [%d] %v \t", index, record[0])

		url := random.Generate(6)

		urll, email, nom, ap, opt, fe, fc := " ", " ", " ", " ", " ", " ", ""
		stmt, err := db.Prepare("SELECT * FROM usuarios WHERE url = ? OR email = ?")

		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		err = stmt.QueryRow(url, record[0]).Scan(&urll, &email, &nom, &ap, &opt, &fe, &fc)

		switch err {

		case sql.ErrNoRows:

			fmt.Print("No rows were returned!")

			if index != 1 {

				stmtt, err := db.Prepare("INSERT INTO usuarios(url, email) VALUES(?, ?)")
				if err != nil {
					log.Fatal(err)
				}
				defer stmtt.Close() // Prepared statements take up server resources and should be closed after use.

				if _, err := stmtt.Exec(url, record[0]); err != nil {
					log.Fatal(err)
				}
			}

		case nil:

			fmt.Print("Row already exist")

		default:

			panic(err)

		}

		fmt.Println(" ")
		index++

		// fmt.Println("Max Open Conn: ", db.Stats().MaxOpenConnections, " Open Conns: ", db.Stats().OpenConnections)
		// fmt.Println("Cons in use: ", db.Stats().InUse, " Cons Idle: ", db.Stats().Idle)

	}

}
