package main

import (
	"log"

	"catetin-backend/internal/app"
	"catetin-backend/internal/config"
	"catetin-backend/internal/database"
	"catetin-backend/internal/modules/auth"
	"catetin-backend/internal/modules/transaction"
)

func main() {
	config.LoadEnv()
	database.Connect()

	authModule := auth.NewModule(database.DB)
	transactionModule := transaction.NewModule(database.DB)

	app := app.New(
		authModule,
		transactionModule,
	)

	app.Setup()

	log.Fatal(app.Fiber.Listen(":3000"))
}
