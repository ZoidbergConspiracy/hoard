package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/zoidbergconspiracy/hoard/auth"
	"github.com/zoidbergconspiracy/hoard/log"
	"github.com/zoidbergconspiracy/hoard/route"
)

func main() {

	hoard := route.HoardHandler{}

	srv := &http.Server{
		Addr:    "0.0.0.0:9443",
		Handler: log.NewLoggingHandler(auth.NewAuthHandler(&hoard, os.Stderr), os.Stdout),
	}

	tlsConfig := &tls.Config{}
	tlsConfig.ClientAuth = tls.RequestClientCert
	srv.TLSConfig = tlsConfig

	fmt.Println(srv.ListenAndServeTLS("server.crt", "server.key"))

}
