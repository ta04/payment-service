// payment-service/main.go

package main

import (
	"log"

	"github.com/SleepingNext/payment-service/database"
	"github.com/SleepingNext/payment-service/handler"
	paymentPB "github.com/SleepingNext/payment-service/proto"
	"github.com/SleepingNext/payment-service/repository/postgres"
	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
)

func main() {
	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.srv.payment"),
	)

	// Initialize the service
	s.Init()

	// Connect to postgres
	db, err := database.OpenPostgresConnection()
	if err != nil {
		log.Fatalf("failed to connect to postgres: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	// Create a new handler
	h := handler.NewHandler(&postgres.Repository{
		DB: db,
	})

	// Register the handler
	paymentPB.RegisterPaymentServiceHandler(s.Server(), h)

	// Run the service
	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
