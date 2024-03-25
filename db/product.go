package db

import "database/sql"

type DB struct {
	db *sql.DB
}

func New(db *sql.DB) DB {
	return DB{db: db}
}

func (d DB) CreateProductTable() error {
	query := "CREATE TABLE IF NOT EXISTS products(id INT auto_increment primary key,name varchar(15),price float4,in_stock bool)"
	_, err := d.db.Exec(query)
	if err != nil {
		return err
	}

	_, err = d.db.Exec(`INSERT INTO products(name,price,in_stock) values("colgate",20.0,true)`)

	return err
}
