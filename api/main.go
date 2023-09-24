package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/mtnmunuklu/bavul/api/handlers"
	"github.com/mtnmunuklu/bavul/api/routes"
	"github.com/mtnmunuklu/bavul/pb"
	"github.com/mtnmunuklu/bavul/security"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port     int
	local    bool
	authAddr string
	vulnAddr string
)

func init() {
	flag.BoolVar(&local, "local", true, "run api service local")
	flag.IntVar(&port, "port", 9000, "api service port")
	flag.StringVar(&authAddr, "auth_addr", "localhost:9001", "authentication service address")
	flag.StringVar(&vulnAddr, "vuln_addr", "localhost:9002", "vulnerability service address")

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

	// for vulnerability service
	vulnConn, err := grpc.Dial(vulnAddr, grpc.WithTransportCredentials(tlsCredentials))
	if err != nil {
		log.Panicln(err)
	}
	defer authConn.Close()

	vulnSvcClient := pb.NewAuthServiceClient(vulnConn)
	vulnHandlers := handlers.NewAuthHandlers(vulnSvcClient)
	vulnRoutes := routes.NewAuthRoutes(vulnHandlers)

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New())

	routes.Install(app, authRoutes)
	routes.Install(app, vulnRoutes)

	log.Printf("API service running on [::]:%d\n", port)

	log.Fatal(app.Listen(fmt.Sprintf(":%d", port)))
}
