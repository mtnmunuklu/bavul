package main

import (
	"fmt"
	"os"
	"os/exec"
)

func runCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Cleaning and testing
	// fmt.Println("Cleaning and testing...")
	// err := runCommand("go clean --cache && go test -v -cover ./security/... ./authentication/...")
	// if err != nil {
	// 	fmt.Println("Error cleaning and testing:", err)
	// 	return
	// }

	// Build authentication service
	fmt.Println("Building authentication service...")
	err := runCommand("go build -o ../authentication/authsvc ../authentication/main.go")
	if err != nil {
		fmt.Println("Error building authentication service:", err)
		return
	}

	// Build vulnerability service
	fmt.Println("Building vulnerability service...")
	err = runCommand("go build -o ../vulnerability/vulnsvc ../vulnerability/main.go")
	if err != nil {
		fmt.Println("Error building vulnerability service:", err)
		return
	}

	// Build API service
	fmt.Println("Building API service...")
	err = runCommand("go build -o ../api/apisvc ../api/main.go")
	if err != nil {
		fmt.Println("Error building API service:", err)
		return
	}

	// Build web service
	fmt.Println("Building web service...")
	err = runCommand("go build -o ../web/websvc ../web/main.go")
	if err != nil {
		fmt.Println("Error building web service:", err)
		return
	}

	fmt.Println("Build process completed successfully.")
}
