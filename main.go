package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	ProductDB "product/db"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

type Product struct {
	Id      int     `json:"id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
	InStock bool    `json:"inStock"`
}

func main() {
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1)/products")
	if err != nil {
		fmt.Printf("error while opening sql connection, err : %v", err)

		return
	}

	database := ProductDB.New(db)
	err = database.CreateProductTable()
	if err != nil {
		fmt.Printf("error while creating tables, err : %v", err)

		return
	}

	http.HandleFunc("/product", func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		if id == "" {
			http.Error(w, "empty id", 400)

			return
		}

		//id = path.Base(r.URL.Path)

		idInt, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid id", 400)

			return
		}

		var p Product

		row := db.QueryRow("SELECT id,name,price,in_stock from products where id = ?", idInt)
		err = row.Scan(&p.Id, &p.Name, &p.Price, &p.InStock)
		if err == sql.ErrNoRows {
			http.Error(w, "Entity not found for Id : "+id, 404)

			return
		}

		if err != nil {
			http.Error(w, "Internal server error, err : "+err.Error(), 500)

			return
		}

		resp, err := json.Marshal(p)

		_, _ = w.Write(resp)
		w.WriteHeader(200)

		return
	})

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
