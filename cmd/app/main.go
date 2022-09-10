package main

import "github.com/migmatore/bakery-shop-api/internal/app"

func main() {
	app.Run("8181", "postgresql://migmatore:root@localhost:5432/testdb")

}
