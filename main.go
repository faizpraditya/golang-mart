package main

func main() {
	db := connectDB()
	MainMenu(db)
	// DeleteProduct(db, "p002")
	// AddProduct(db, "p002", "pena", 7, 3000)
	// DetailProduct(db)
}
