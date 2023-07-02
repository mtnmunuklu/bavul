package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mtnmunuklu/bavul/api/handlers"
	"github.com/mtnmunuklu/bavul/api/middlewares"
	"github.com/mtnmunuklu/bavul/api/routes"
	"github.com/mtnmunuklu/bavul/pb"
	"github.com/mtnmunuklu/bavul/security"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port     int
	local    bool
	authAddr string
)

func init() {
	flag.BoolVar(&local, "local", true, "run api service local")
	flag.IntVar(&port, "port", 9000, "api service port")
	flag.StringVar(&authAddr, "auth_addr", "localhost:9001", "authentication service address")
	flag.Parse()
}

func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}

	// for secure communication
	certPath := os.Getenv("CERT_PATH")
	tlsCredentials, err := security.LoadCATLSCredentials(certPath)
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}

	// for authentication service
	authConn, err := grpc.Dial(authAddr, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Panicln(err)
	}
	defer authConn.Close()

	autSvcClient := pb.NewAuthServiceClient(authConn)
	authHandlers := handlers.NewAuthHandlers(autSvcClient)
	authRoutes := routes.NewAuthRoutes(authHandlers)

	app := fiber.New()
	app.Use(middlewares.CORS())

	routes.Install(app, authRoutes)

	log.Printf("API service running on [::]:%d\n", port)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
