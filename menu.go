package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var scanner = bufio.NewScanner(os.Stdin)

func mainMenu() {
	fmt.Println("ENIGMA-MART")
	fmt.Println("1. Produkt Barang")
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
