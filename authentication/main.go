package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/mtnmunuklu/bavul/authentication/repository"
	"github.com/mtnmunuklu/bavul/authentication/service"
	"github.com/mtnmunuklu/bavul/db"
	"github.com/mtnmunuklu/bavul/pb"
	"github.com/mtnmunuklu/bavul/security"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// Contains some variables(port, local) for authentication service.
var (
	port  int
	local bool
)

// Init initializes the specify options for authentication service.
func init() {
	flag.IntVar(&port, "port", 9001, "authentication service port")
	flag.BoolVar(&local, "local", true, "run authentication service local")
	flag.Parse()
}

// Main starts the authentication service.
func main() {
	if local {
		err := godotenv.Load()
		if err != nil {
			log.Panicln(err)
		}
	}

	cfg := db.NewConfig()
	conn, err := db.NewConnection(cfg)
	if err != nil {
		log.Panicln(err)
	}
	defer conn.Close()

	userRepository := repository.NewUserRepository(conn)
	authService := service.NewAuthService(userRepository)

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	cert_path := os.Getenv("CERT_PATH")
	tlsCredentials, err := security.LoadServerTLSCredentials(cert_path)
	if err != nil {
		log.Fatal("cannot load TLS credentials: ", err)
	}
	grpcServer := grpc.NewServer(grpc.Creds(tlsCredentials))
	pb.RegisterAuthServiceServer(grpcServer, authService)

	log.Printf("Authentication service running on [::]:%d\n", port)

	grpcServer.Serve(listen)
}
