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

func MainMenu(db *sqlx.DB) {
	fmt.Println("ENIGMA-MART")
	fmt.Println("1. Produk Barang")
	fmt.Println("2. Transaksi Penjualan")
	fmt.Println("3. Laporan Penjualan")
	fmt.Println("4. Exit")
	fmt.Print("Pilih Menu: ")
	scanner.Scan()
	menu, _ := strconv.Atoi(scanner.Text())
	mainMenuController(menu, db)
}

func mainMenuController(menu int, db *sqlx.DB) {
	switch menu {
	case 1:
		productMenu(db)
	case 2:
	case 3:
	case 4:
		fmt.Println("Exit from program")
		os.Exit(1)
	default:
		defer MainMenu(db)
		fmt.Println("Wrong input")
		return
	}
}

func productMenu(db *sqlx.DB) {
	fmt.Println("PRODUK BARANG")
	fmt.Println("1. Tambah Produk")
	fmt.Println("2. Hapus Produk")
	fmt.Println("3. Detail Produk")
	fmt.Println("4. Back to main menu")
	fmt.Print("Pilih Menu: ")
	scanner.Scan()
	menuProduct, _ := strconv.Atoi(scanner.Text())
	productMenuController(menuProduct, db)
}

func productMenuController(menu int, db *sqlx.DB) {
	switch menu {
	case 1:
		defer MainMenu(db)
		fmt.Println("TAMBAH PRODUK")
		fmt.Print("Id produk: ")
		scanner.Scan()
		inID := scanner.Text()
		fmt.Print("Nama produk: ")
		scanner.Scan()
		inNama := scanner.Text()
		fmt.Print("Stok produk: ")
		scanner.Scan()
		inStock, _ := strconv.Atoi(scanner.Text())
		fmt.Print("Harga produk: ")
		scanner.Scan()
		inPrice, _ := strconv.Atoi(scanner.Text())
		AddProduct(db, inID, inNama, inStock, inPrice)
		return
	case 2:
		defer MainMenu(db)
		fmt.Println("HAPUS PRODUCT")
		fmt.Print("Id produk: ")
		scanner.Scan()
		DeleteProduct(db, scanner.Text())
		return
	case 3:
		defer MainMenu(db)
		DetailProduct(db)
		return
	case 4:
		defer MainMenu(db)
		return
	default:
		defer productMenu(db)
		fmt.Println("Wrong input")
		return
	}
}

func AddProduct(db *sqlx.DB, id string, name string, stock int, price int) {
	product := Products{Id: id, Name: name, Stock: stock, Price: price}
	_, err := db.NamedExec(INSERT_PRODUCT, product)
	if err != nil {
		log.Fatal(err)
	} else {
		check_error(err, "insert")
	}
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

func DetailProduct(db *sqlx.DB) {
	products := []Products{}
	err := db.Select(&products, DETAIL_PRODUCT)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(products)
}

func check_error(err error, s string) {
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Successfully " + s + " to database")
	}
}
