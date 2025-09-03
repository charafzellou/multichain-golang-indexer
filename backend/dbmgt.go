package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq" // driver Postgres
)

// Constantes Postgres
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgresuser"
	password = "postgrespwd"
	dbname   = "postgres"
	schema   = "public"
)

func initDB() *sql.DB {
	var dbpointer = connectDB()

	dropTable(dbpointer, "contracts")
	createTableContracts(dbpointer, "contracts")

	dropTable(dbpointer, "operations")
	createTableOperations(dbpointer, "operations")

	return dbpointer
}

func connectDB() *sql.DB {
	connectionString := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbname + " sslmode=disable"
	fmt.Println("connectionString: ", connectionString)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Connexion successful to database: " + dbname)
	}

	return db
}

func createTableContracts(db *sql.DB, tableName string) {
	sqlStatement := "CREATE TABLE IF NOT EXISTS " + tableName + " ( address TEXT NOT NULL, eth_balance FLOAT, creator TEXT, creation_timestamp BIGINT, transactions BIGINT, PRIMARY KEY (address) )"
	fmt.Println("createTableContracts: ", sqlStatement)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func createTableOperations(db *sql.DB, tableName string) {
	sqlStatement := "CREATE TABLE IF NOT EXISTS " + tableName + " ( hash TEXT, timestamp BIGINT, from_address TEXT, to_address TEXT, value FLOAT, input TEXT, success BOOLEAN, PRIMARY KEY (hash) )"
	fmt.Println("createTableOperations: ", sqlStatement)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}

func insertRowContracts(db *sql.DB, tableName string, data []string) {
	// address TEXT, eth_balance FLOAT, creator TEXT, creation_timestamp BIGINT, transactions BIGINT
	if len(data) == 5 {
		sqlStatement := "INSERT INTO " + tableName + " (address, eth_balance, creator, creation_timestamp, transactions) VALUES ('" + data[0] + "', " + data[1] + ", '" + data[2] + "', " + data[3] + ", " + data[4] + ");"
		fmt.Println("insertRowContracts: ", sqlStatement)
		_, err := db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}
	} else {
		panic("Not the right amount of parameters in data[] : expected(5)")
	}
}
func insertRowOperations(db *sql.DB, tableName string, data []string) {
	// hash TEXT, timestamp BIGINT, from_address TEXT, to_address TEXT, value FLOAT, input TEXT, success BOOLEAN
	if len(data) == 7 {
		sqlStatement := "INSERT INTO " + tableName + " (hash, timestamp, from_address, to_address, value, input, success) VALUES ('" + data[0] + "', " + data[1] + ", '" + data[2] + "', '" + data[3] + "', " + data[4] + ", '" + data[5] + "', " + data[6] + ");"
		fmt.Println("insertRowOperations: ", sqlStatement)
		_, err := db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}
	} else {
		panic("Not the right amount of parameters in data[] : expected(7)")
	}

}

func selectRow(db *sql.DB, tableName string, condition string) *sql.Row {
	sqlStatement := "SELECT * FROM public." + tableName + " WHERE " + condition
	fmt.Println("selectRow: ", sqlStatement)
	sRow := db.QueryRow(sqlStatement)
	return sRow
}

func dropTable(db *sql.DB, tableName string) {
	sqlStatement := "DROP TABLE IF EXISTS " + tableName
	fmt.Println("dropTable: ", sqlStatement)
	_, err := db.Exec(sqlStatement)
	if err != nil {
		panic(err)
	}
}
