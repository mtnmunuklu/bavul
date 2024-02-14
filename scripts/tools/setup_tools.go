package main

import (
	"fmt"
	"os"
	"os/exec"
)

const (
	RED      = "\033[31m"
	GREEN    = "\033[32m"
	ENDCOLOR = "\033[0m"
)

func runCommand(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Install Tools
	fmt.Println(GREEN + "Install Tools" + ENDCOLOR)

	err := runCommand("sudo yum update -y")
	if err != nil {
		fmt.Println("Error updating yum:", err)
		return
	}

	err = runCommand("sudo yum upgrade -y")
	if err != nil {
		fmt.Println("Error upgrading yum:", err)
		return
	}

	err = runCommand("sudo yum install -y wget tar")
	if err != nil {
		fmt.Println("Error installing wget and tar:", err)
		return
	}

	// Install Go
	fmt.Println(GREEN + "Install Go" + ENDCOLOR)

	// Download the Go binary
	fmt.Println(GREEN + "Download the Go binary" + ENDCOLOR)
	err = runCommand("wget https://dl.google.com/go/go1.19.2.linux-amd64.tar.gz --no-check-certificate")
	if err != nil {
		fmt.Println("Error downloading Go binary:", err)
		return
	}

	// Extract the archive into /usr/local
	fmt.Println(GREEN + "Extract the archive into /usr/local" + ENDCOLOR)
	err = runCommand("rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.2.linux-amd64.tar.gz")
	if err != nil {
		fmt.Println("Error extracting Go archive:", err)
		return
	}

	// Add /usr/local/go/bin to the PATH environment variable
	fmt.Println(GREEN + "Add /usr/local/go/bin to the PATH environment variable" + ENDCOLOR)
	err = runCommand("export PATH=$PATH:/usr/local/go/bin >> ~/.bashrc")
	if err != nil {
		fmt.Println("Error adding Go to PATH:", err)
		return
	}

	err = runCommand("source ~/.bashrc")
	if err != nil {
		fmt.Println("Error sourcing bashrc:", err)
		return
	}

	// Verify Go installation
	fmt.Println(GREEN + "Verify Go installation" + ENDCOLOR)
	err = runCommand("go version")
	if err != nil {
		fmt.Println("Error verifying Go installation:", err)
		return
	}

	// Install Docker
	fmt.Println(GREEN + "Install Docker" + ENDCOLOR)

	// Uninstall old Docker versions
	fmt.Println(GREEN + "Uninstall old Docker versions" + ENDCOLOR)
	err = runCommand("sudo yum remove docker docker-client docker-client-latest docker-common docker-latest docker-latest-logrotate docker-logrotate docker-engine")
	if err != nil {
		fmt.Println("Error uninstalling old Docker versions:", err)
		return
	}

	// Install yum-utils and set up the stable repository
	fmt.Println(GREEN + "Install yum-utils and set up the stable repository" + ENDCOLOR)
	err = runCommand("sudo yum install -y yum-utils")
	if err != nil {
		fmt.Println("Error installing yum-utils:", err)
		return
	}

	err = runCommand("sudo yum-config-manager --add-repo https://download.docker.com/linux/centos/docker-ce.repo")
	if err != nil {
		fmt.Println("Error adding Docker repository:", err)
		return
	}

	// Install Docker engine
	fmt.Println(GREEN + "Install Docker engine" + ENDCOLOR)
	err = runCommand("sudo yum install -y docker-ce docker-ce-cli")
	if err != nil {
		fmt.Println("Error installing Docker engine:", err)
		return
	}

	// Enable Docker
	fmt.Println(GREEN + "Enable Docker" + ENDCOLOR)
	err = runCommand("sudo systemctl enable docker")
	if err != nil {
		fmt.Println("Error enabling Docker:", err)
		return
	}

	// Start Docker
	fmt.Println(GREEN + "Start Docker" + ENDCOLOR)
	err = runCommand("sudo systemctl start docker")
	if err != nil {
		fmt.Println("Error starting Docker:", err)
		return
	}

	fmt.Println("Installation process completed successfully.")
}
