// cmd/main.go

package main

import (
	"crypto/tls"
	"log"
	"mitm-proxy/certs"
	"mitm-proxy/pkg/auth"
	"mitm-proxy/pkg/config"
	"net/http"
)

func handleRequest(response http.ResponseWriter, request *http.Request) {
	// ... 代理的核心代码
}

func main() {
	cfg, err := config.LoadConfig("path/to/config.json")
	if err != nil {
		log.Fatal(err)
	}

	certificate, err := certs.LoadCertificate(cfg.CACert, cfg.CAKey)
	if err != nil {
		log.Fatal(err)
	}

	caCertPool, err := certs.LoadCA(cfg.CACert)
	if err != nil {
		log.Fatal(err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{certificate},
		RootCAs:      caCertPool,
	}

	server := &http.Server{
		Addr:      ":8080",
		TLSConfig: tlsConfig,
	}

	http.HandleFunc("/", auth.BasicAuthMiddleware(handleRequest, "path/to/your/database.sqlite", "Please enter your username and password"))
	log.Fatal(server.ListenAndServeTLS("", ""))
}
