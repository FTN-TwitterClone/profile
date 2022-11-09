package tls

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"os"
)

func GetgRPCServerTLSConfig() *tls.Config {
	serverCertPath := os.Getenv("CERT")
	serverKeyPath := os.Getenv("KEY")
	caCertPath := os.Getenv("CA_CERT")

	serverCert, err := tls.LoadX509KeyPair(serverCertPath, serverKeyPath)
	if err != nil {
		log.Fatalf("Failed to load server certificate and key. %s.", err)
	}

	trustedCert, err := ioutil.ReadFile(caCertPath)
	if err != nil {
		log.Fatalf("Failed to load trusted certificate. %s.", err)
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(trustedCert) {
		log.Fatalf("Failed to append trusted certificate to certificate pool. %s.", err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		RootCAs:      certPool,
		ClientCAs:    certPool,
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
	}
}
