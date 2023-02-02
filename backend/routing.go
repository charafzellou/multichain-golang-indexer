package main

import (
	"database/sql"
	"fmt"
	"net/http"
)

func listenOnPort(dbpointer *sql.DB) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// w.Header().Set("Access-Control-Allow-Origin", "*")
		if r.Method == "POST" {
			err := r.ParseForm()
			if err != nil {
				panic(err)
			}
			receivedString := r.FormValue("address")
			fmt.Println("Received string address: ", receivedString)
			fmt.Fprintf(w, "Received string address: %s", receivedString)

			var exists bool
			err = dbpointer.QueryRow("SELECT EXISTS(SELECT address FROM public.contracts WHERE address = '" + receivedString + "')").Scan(&exists)
			if err != nil {
				panic(err)
			}

			if exists {
				fmt.Println("Contract already existing in DB")
			} else {
				indexContract(dbpointer, receivedString)
				fmt.Println("Contract Indexed!")
			}
		} else {
			fmt.Println("r.Method: ", r.Method)
			panic("Only POST requests are accepted")
		}
	})

	err := http.ListenAndServe(":3200", nil)
	if err != nil {
		panic(err)
	}
}

func indexContract(dbpointer *sql.DB, addr string) {
	insertRowContracts(dbpointer, "contracts", getContract(addr))
	for _, opData := range getOperations(addr) {
		insertRowOperations(dbpointer, "operations", opData)
	}
}
