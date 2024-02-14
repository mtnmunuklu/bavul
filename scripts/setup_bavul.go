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

func applyK8sConfigs() {
	// Apply traefik configuration
	fmt.Println(GREEN + "Apply traefik configuration" + ENDCOLOR)
	runCommand("kubectl apply -f ../k8s/traefik/00-role.yml " +
		"-f ../k8s/traefik/00-account.yml " +
		"-f ../k8s/traefik/01-role-binding.yml " +
		"-f ../k8s/traefik/02-traefik.yml " +
		"-f ../k8s/traefik/02-traefik-services.yml")

	// Apply common configurations
	fmt.Println(GREEN + "Apply common configurations" + ENDCOLOR)
	runCommand("kubectl apply -f ../k8s/common/00-scripts.yml " +
		"-f ../k8s/common/01-ds.yml")

	// Apply mongodb configuration
	fmt.Println(GREEN + "Apply mongodb configuration" + ENDCOLOR)
	runCommand("kubectl apply -f ../k8s/mongodb/00-stc.yml " +
		"-f ../k8s/mongodb/01-pv.yml " +
		"-f ../k8s/mongodb/02-pvc.yml " +
		"-f ../k8s/mongodb/03-configs.yml " +
		"-f ../k8s/mongodb/04-scripts.yml " +
		"-f ../k8s/mongodb/05-secrets.yml " +
		"-f ../k8s/mongodb/06-sts.yml")
}

func generateCertificates() {
	// Generate certificates for docker registry, services, and ingress
	fmt.Println(GREEN + "Generate certificates for docker registry, services, and ingress" + ENDCOLOR)
	runCommand("go run generate_certificates.go")
}

func createSecrets() {
	// Create secret for secure communication between services
	fmt.Println(GREEN + "Create secret for secure communication between services" + ENDCOLOR)
	runCommand("kubectl create secret generic service-certs " +
		"--from-file=../certs/services/ca-cert.pem " +
		"--from-file=../certs/services/server-cert.pem " +
		"--from-file=../certs/services/server-key.pem")

	// Create secret for ingress
	fmt.Println(GREEN + "Create secret for ingress" + ENDCOLOR)
	runCommand("kubectl create secret tls ingress-certs " +
		"--key ../certs/api/bavul-key.pem " +
		"--cert ../certs/api/bavul-cert.pem")
}

func applyServiceConfigs() {
	// Apply service configuration
	fmt.Println(GREEN + "Apply service configuration" + ENDCOLOR)
	runCommand("kubectl apply -f ../k8s/services/00-configs.yml " +
		"-f ../k8s/services/01-secrets.yml " +
		"-f ../k8s/services/02-authentication.yml " +
		"-f ../k8s/services/03-vulnerability.yml " +
		"-f ../k8s/services/04-api.yml " +
		"-f ../k8s/services/05-web.yml " +
		"-f ../k8s/services/06-tls_ingress.yml")
}

func getAllInformation() {
	// Get all information
	fmt.Println(GREEN + "Get all information" + ENDCOLOR)
	runCommand("kubectl get all")
}

func main() {
	applyK8sConfigs()
	generateCertificates()
	createSecrets()
	applyServiceConfigs()
	getAllInformation()
}
