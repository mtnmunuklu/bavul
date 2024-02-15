package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

const (
	caKeyFile      = "../certs/services/ca-key.pem"
	caCertFile     = "../certs/services/ca-cert.pem"
	serverKeyFile  = "../certs/services/server-key.pem"
	serverReqFile  = "../certs/services/server-req.pem"
	serverCertFile = "../certs/services/server-cert.pem"
	apiCertFile    = "../certs/api/bavul-cert.pem"
	apiKeyFile     = "../certs/api/bavul-key.pem"
)

func main() {
	// Remove old certificates
	fmt.Println("Remove old certificates")
	os.Remove(serverCertFile)
	os.Remove(serverReqFile)
	os.Remove(serverKeyFile)
	os.Remove(caCertFile)
	os.Remove(caKeyFile)
	os.Remove(apiCertFile)
	os.Remove(apiKeyFile)

	// 1. Generate CA's private key and self-signed certificate
	fmt.Println("1. Generate CA's private key and self-signed certificate")
	caKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Printf("Error generating CA private key: %v\n", err)
		return
	}
	caCert, err := generateCACertificate(caKey)
	if err != nil {
		fmt.Printf("Error generating CA certificate: %v\n", err)
		return
	}
	savePrivateKey(caKeyFile, caKey)
	saveCertificate(caCertFile, caCert)

	// CA's self-signed certificate
	fmt.Println("CA's self-signed certificate")
	displayCertificate(caCert)

	// 2. Generate web server's private key and certificate signing request (CSR)
	fmt.Println("2. Generate web server's private key and certificate signing request (CSR)")
	serverKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Printf("Error generating server private key: %v\n", err)
		return
	}
	savePrivateKey(serverKeyFile, serverKey)

	serverReq, err := generateCertificateRequest(serverKey)
	if err != nil {
		fmt.Printf("Error generating server certificate request: %v\n", err)
		return
	}
	saveCertificateRequest(serverReqFile, serverReq)

	// 3. Use CA's private key to sign web server's CSR and get back the signed certificate
	fmt.Println("3. Use CA's private key to sign web server's CSR and get back the signed certificate")
	serverCert, err := signCertificateRequest(serverReq, caCert, caKey)
	if err != nil {
		fmt.Printf("Error signing server certificate request: %v\n", err)
		return
	}
	saveCertificate(serverCertFile, serverCert)

	// Server's signed certificates
	fmt.Println("Server's signed certificates")
	displayCertificate(serverCert)

	// Creation of certificates for ingress
	fmt.Println("Creation of certificates for ingress")
	apiKey, err := ecdsa.GenerateKey(elliptic.P384(), rand.Reader)
	if err != nil {
		fmt.Printf("Error generating API private key: %v\n", err)
		return
	}
	savePrivateKey(apiKeyFile, apiKey)

	apiCert, err := generateCertificate(apiKey, caCert, caKey, "DNS:bavul.com, DNS:api.bavul.com, DNS:api-service, DNS:web-service")
	if err != nil {
		fmt.Printf("Error generating API certificate: %v\n", err)
		return
	}
	saveCertificate(apiCertFile, apiCert)
}

func generateCACertificate(caKey *ecdsa.PrivateKey) ([]byte, error) {
	caTemplate := x509.Certificate{
		Subject:               pkix.Name{CommonName: "CA"},
		SerialNumber:          big.NewInt(1),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	caCertDER, err := x509.CreateCertificate(rand.Reader, &caTemplate, &caTemplate, &caKey.PublicKey, caKey)
	if err != nil {
		return nil, err
	}

	return caCertDER, nil
}

func generateCertificateRequest(privateKey *ecdsa.PrivateKey) ([]byte, error) {
	reqTemplate := x509.CertificateRequest{
		Subject: pkix.Name{CommonName: "*.bavul.com"},
	}

	reqDER, err := x509.CreateCertificateRequest(rand.Reader, &reqTemplate, privateKey)
	if err != nil {
		return nil, err
	}

	return reqDER, nil
}

func signCertificateRequest(req []byte, caCertBytes []byte, caKey *ecdsa.PrivateKey) ([]byte, error) {
	reqCert, err := x509.ParseCertificateRequest(req)
	if err != nil {
		return nil, err
	}

	caCert, err := x509.ParseCertificate(caCertBytes)
	if err != nil {
		return nil, err
	}

	certTemplate := x509.Certificate{
		Subject:               reqCert.Subject,
		SerialNumber:          big.NewInt(2),
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(1, 0, 0),
		KeyUsage:              x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth, x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &certTemplate, caCert, &reqCert.PublicKey, caKey)
	if err != nil {
		return nil, err
	}

	return certDER, nil
}

func generateCertificate(privateKey *ecdsa.PrivateKey, caCertBytes []byte, caKey *ecdsa.PrivateKey, dnsNames string) ([]byte, error) {
	caCert, err := x509.ParseCertificate(caCertBytes)
	if err != nil {
		return nil, err
	}

	certTemplate := x509.Certificate{
		Subject:  pkix.Name{CommonName: "*.bavul.com"},
		DNSNames: []string{"bavul.com", "api.bavul.com", "api-service", "web-service"},
	}

	certDER, err := x509.CreateCertificate(rand.Reader, &certTemplate, caCert, &privateKey.PublicKey, caKey)
	if err != nil {
		return nil, err
	}

	return certDER, nil
}

func savePrivateKey(filename string, key *ecdsa.PrivateKey) {
	keyBytes, err := x509.MarshalECPrivateKey(key)
	if err != nil {
		fmt.Printf("Error marshaling private key: %v\n", err)
		return
	}

	pemBlock := &pem.Block{Type: "EC PRIVATE KEY", Bytes: keyBytes}
	err = os.WriteFile(filename, pem.EncodeToMemory(pemBlock), 0644)
	if err != nil {
		fmt.Printf("Error writing private key to file: %v\n", err)
	}
}

func saveCertificateRequest(filename string, req []byte) {
	pemBlock := &pem.Block{Type: "CERTIFICATE REQUEST", Bytes: req}
	err := os.WriteFile(filename, pem.EncodeToMemory(pemBlock), 0644)
	if err != nil {
		fmt.Printf("Error writing certificate request to file: %v\n", err)
	}
}

func saveCertificate(filename string, cert []byte) {
	pemBlock := &pem.Block{Type: "CERTIFICATE", Bytes: cert}
	err := os.WriteFile(filename, pem.EncodeToMemory(pemBlock), 0644)
	if err != nil {
		fmt.Printf("Error writing certificate to file: %v\n", err)
	}
}

func displayCertificate(cert []byte) {
	parsedCert, err := x509.ParseCertificate(cert)
	if err != nil {
		fmt.Printf("Error parsing certificate: %v\n", err)
		return
	}

	fmt.Printf("Subject: %s\n", parsedCert.Subject.CommonName)
	fmt.Printf("Issuer: %s\n", parsedCert.Issuer.CommonName)
	fmt.Printf("Serial Number: %s\n", parsedCert.SerialNumber)
	fmt.Printf("Not Before: %s\n", parsedCert.NotBefore)
	fmt.Printf("Not After: %s\n", parsedCert.NotAfter)
}
