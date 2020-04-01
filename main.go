// payment-service/main.go

package main

import (
	"context"
	"errors"
	"log"
	"os"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"

	authPB "github.com/SleepingNext/auth-service/proto"
	"github.com/SleepingNext/payment-service/database"
	"github.com/SleepingNext/payment-service/handler"
	paymentPB "github.com/SleepingNext/payment-service/proto"
	"github.com/SleepingNext/payment-service/repository/postgres"
	_ "github.com/lib/pq"
	"github.com/micro/go-micro"
	"github.com/micro/go-plugins/registry/consul"
)

func main() {
	// Take or set the port
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":50053"
	}

	// Create a new registry
	registry := consul.NewRegistry()

	// Create a new service
	s := micro.NewService(
		micro.Name("com.ta04.srv.payment"),
		micro.WrapHandler(AuthWrapper),
		micro.Address(port),
		micro.Registry(registry),
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

func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, res interface{}) error {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in the request")
		}

		token := meta["Token"]
		log.Println("authenticating with token: ", token)

		// Validate the token
		authClient := authPB.NewAuthServiceClient("com.ta04.srv.auth", client.DefaultClient)
		_, err := authClient.ValidateToken(context.Background(), &authPB.Token{
			Token: token,
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, res)
		return err
	}
}
