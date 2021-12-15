package main

// dibuat public karena akan dipakai di package sqlx
type Customers struct {
	Id      string `db:"id"`
	Name    string `db:"name"`
	Address string `db:"address"`
}

type Products struct {
	Id     string `db:"id"`
	Name   string `db:"name"`
	Stock  int    `db:"stock"`
	Price  int    `db:"price"`
	Status int    `db:"status"`
}
