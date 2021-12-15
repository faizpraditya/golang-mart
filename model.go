package main

import "time"

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
	Qty    int    `db:"qty"`
}

type DetailTransactions struct {
	Id            string    `db:"id"`
	Purchase_id   string    `db:"purchase_id"`
	Product_id    string    `db:"product_id"`
	Price         int       `db:"price"`
	Qty           int       `db:"qty"`
	Purchase_date time.Time `db:"purchase_date"`
	Customer_id   string    `db:"customer_id"`
	Amount        int       `db:"amount"`
}

type Transactions struct {
	Id            string    `db:"id"`
	Purchase_date time.Time `db:"purchase_date"`
	Customer_id   string    `db:"customer_id"`
}
