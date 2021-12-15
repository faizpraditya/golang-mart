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
		defer productMenu(db)
		return
	case 2:
		defer transactionMenu(db)
		return
	case 3:
		defer reportMenu(db)
		return
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
	fmt.Println("4. Kembali ke menu utama")
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
		inId := scanner.Text()
		fmt.Print("Nama produk: ")
		scanner.Scan()
		inNama := scanner.Text()
		fmt.Print("Stok produk: ")
		scanner.Scan()
		inStock, _ := strconv.Atoi(scanner.Text())
		fmt.Print("Harga produk: ")
		scanner.Scan()
		inPrice, _ := strconv.Atoi(scanner.Text())
		AddProduct(db, inId, inNama, inStock, inPrice)
		return
	case 2:
		defer MainMenu(db)
		fmt.Println("HAPUS PRODUK")
		fmt.Print("Id produk: ")
		scanner.Scan()
		DeleteProduct(db, scanner.Text())
		return
	case 3:
		defer MainMenu(db)
		fmt.Println("DETAIL PRODUK")
		fmt.Print("Id produk: ")
		scanner.Scan()
		DetailProduct(db, scanner.Text())
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

func transactionMenu(db *sqlx.DB) {
	fmt.Println("TRANSAKSI PENJUALAN")
	fmt.Println("1. Tambah Transaksi Penjualan")
	fmt.Println("2. Detail Transaksi Penjualan")
	fmt.Println("3. Kembali ke menu utama")
	fmt.Print("Pilih Menu: ")
	scanner.Scan()
	menuTransaction, _ := strconv.Atoi(scanner.Text())
	transactionMenuController(menuTransaction, db)
}

func transactionMenuController(menu int, db *sqlx.DB) {
	switch menu {
	case 1:
		defer MainMenu(db)
		fmt.Println("TAMBAH TRANSAKSI PENJUALAN")
		fmt.Print("Id transaksi: ")
		scanner.Scan()
		inTId := scanner.Text()
		fmt.Print("Id produk: ")
		scanner.Scan()
		inPId := scanner.Text()
		fmt.Print("Id customer: ")
		scanner.Scan()
		inCId := scanner.Text()
		fmt.Print("Jumlah produk: ")
		scanner.Scan()
		inQty, _ := strconv.Atoi(scanner.Text())
		AddDetailTransaction(db, inTId, inPId, inQty, inCId)
		// AddTransaction(db, inTId, inCId)
		return
	case 2:
		defer MainMenu(db)
		fmt.Println("DETAIL TRANSAKSI PENJUALAN")
		fmt.Print("Id transaksi: ")
		scanner.Scan()
		DetailTransaction(db, scanner.Text())
		return
	case 3:
		defer MainMenu(db)
		return
	default:
		defer productMenu(db)
		fmt.Println("Wrong input")
		return
	}
}

func reportMenu(db *sqlx.DB) {
	fmt.Println("LAPORAN PENJUALAN")
	fmt.Println("1. Semua laporan")
	fmt.Println("2. Kembali ke menu utama")
	fmt.Print("Pilih Menu: ")
	scanner.Scan()
	menuReport, _ := strconv.Atoi(scanner.Text())
	reportMenuController(menuReport, db)
}

func reportMenuController(menu int, db *sqlx.DB) {
	switch menu {
	case 1:
		defer MainMenu(db)
		SalesReport(db)
		return
	case 2:
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

func CheckStock(db *sqlx.DB, id string) int {
	product := Products{}
	err := db.Get(&product, GET_PRODUCT, id)

	if err != nil {
		log.Fatal(err)
	}

	return product.Stock
}

func GetProductPrice(db *sqlx.DB, id string) int {
	product := Products{}
	err := db.Get(&product, GET_PRODUCT, id)

	if err != nil {
		log.Fatal(err)
	}

	return product.Price
}

func DeleteProduct(db *sqlx.DB, id string) {
	if CheckStock(db, id) > 0 {
		log.Println("Stok masih ada")
	} else {
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
}

func DetailProduct(db *sqlx.DB, id string) {
	product := Products{}
	err := db.Get(&product, GET_PRODUCT, id)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(product)
}

func AddDetailTransaction(db *sqlx.DB, id string, product_id string, qty int, customer_id string) {
	if CheckStock(db, product_id) > 0 {
		tx := db.MustBegin()

		AddTransaction(db, id, customer_id)

		detailTransaction := DetailTransactions{Id: id, Purchase_id: id, Product_id: product_id, Price: GetProductPrice(db, product_id), Qty: qty}
		_, err := db.NamedExec(INSERT_DETAIL_TRANSACTION, detailTransaction)
		if err != nil {
			log.Fatal(err)
		} else {
			check_error(err, "insert detail transaction")
		}

		product := Products{Id: product_id, Qty: qty}
		result, _ := tx.NamedExec(UPDATE_STOCK_DECREASE, product)
		r, _ := result.RowsAffected()
		if r == 0 {
			log.Println("kosong bos")
			tx.Rollback()
		}

		tx.Commit()
	} else {
		log.Println("Stok habis")
	}
}

func AddTransaction(db *sqlx.DB, id string, customer_id string) {
	transaction := Transactions{Id: id, Customer_id: customer_id}
	_, err := db.NamedExec(INSERT_TRANSACTION, transaction)
	if err != nil {
		log.Fatal(err)
	} else {
		check_error(err, "insert transaction")
	}
}

func DetailTransaction(db *sqlx.DB, id string) {
	transaction := DetailTransactions{}
	err := db.Get(&transaction, GET_DETAIL_TRANSACTION, id)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(transaction)
}

func SalesReport(db *sqlx.DB) {
	transactions := []DetailTransactions{}
	err := db.Select(&transactions, SALES_REPORT)

	if err != nil {
		log.Fatal(err)
	}

	log.Println(transactions)
	log.Println("Total Omzet: ", TotalOmzet(db))
}

func TotalOmzet(db *sqlx.DB) int {
	omzet := DetailTransactions{}
	err := db.Get(&omzet, TOTAL_OMZET)

	if err != nil {
		log.Fatal(err)
	}

	return omzet.Amount
}

func check_error(err error, s string) {
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("Successfully " + s + " to database")
	}
}
