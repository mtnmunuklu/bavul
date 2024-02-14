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
	// Copy builted services to docker directory
	fmt.Println("Copy builted services to docker directory")
	copyCommands := []string{
		"cp ../authentication/authsvc ../k8s/docker",
		"cp ../vulnerability/vulnsvc ../k8s/docker",
		"cp ../api/apisvc ../k8s/docker",
		"cp ../api/websvc ../k8s/docker",
	}

	for _, cmd := range copyCommands {
		err := runCommand(cmd)
		if err != nil {
			fmt.Printf("Error running copy command '%s': %v\n", cmd, err)
			return
		}
	}

	// Build app docker file
	fmt.Println("Build docker file")
	err := runCommand("docker build -t mtnmunuklu/bavul:v1.0.0 ../k8s/docker/")
	if err != nil {
		fmt.Println("Error building docker file:", err)
		return
	}

	// Push app image to local registry
	fmt.Println("Push app image to local registry")
	err = runCommand("docker push mtnmunuklu/bavul:v1.0.0")
	if err != nil {
		fmt.Println("Error pushing app image to local registry:", err)
		return
	}

	// Show docker images
	fmt.Println("Show docker images")
	err = runCommand("docker images")
	if err != nil {
		fmt.Println("Error showing docker images:", err)
		return
	}

	fmt.Println("Script execution completed successfully.")
}
