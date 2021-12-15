package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
)

var scanner = bufio.NewScanner(os.Stdin)

func mainMenu() {
	fmt.Println("ENIGMA-MART")
	fmt.Println("1. Produk Barang")
	fmt.Println("2. Transaksi Penjualan")
	fmt.Println("3. Laporan Penjualan")
	fmt.Println("4. Exit")
	fmt.Print("Pilih Menu: ")
	scanner.Scan()
	menu, _ := strconv.Atoi(scanner.Text())
	menuController(menu)
}

func menuController(menu int) {
	switch menu {
	case 1:
	case 2:
	case 3:
	case 4:
		fmt.Println("Exit from program")
		os.Exit(1)
	default:
		defer mainMenu()
		fmt.Println("Wrong input")
		return
	}
}

func AddProduct(db *sqlx.DB, id string) {

}

func DeleteProduct(db *sqlx.DB, id string) {
	product := Products{Id: id, Status: 0}
	_, err := db.NamedExec(UPDATE_STATUS, product)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatal("Id doesn't exist")
		}
	} else {
		check_error(err, "delete")
	}
}

func check_error(err error, s string) {
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Successfully " + s + " to database")
	}
}
