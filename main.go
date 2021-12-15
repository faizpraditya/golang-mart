package main

func main() {
	db := connectDB()
	// mainMenu()
	DeleteProduct(db, "p002")
}
