package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/mtnmunuklu/bavul/api/handlers"
	"github.com/mtnmunuklu/bavul/api/routes"
	"github.com/mtnmunuklu/bavul/pb"
	"github.com/mtnmunuklu/bavul/security"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

var (
	port      int
	local     bool
	authAddr  string
	catAddr   string
	crawlAddr string
	catzeAddr string
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
	cert_path := os.Getenv("CERT_PATH")
	tlsCredentials, err := security.LoadCATLSCredentials(cert_path)
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

	router := mux.NewRouter().StrictSlash(true)
	routes.Install(router, authRoutes)

	log.Printf("API service running on [::]:%d\n", port)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), routes.WithCORS(router)))
}
