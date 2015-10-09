package main

import (
	"fmt"
	"os"
		"crypto/tls"
	"net/http"

	"github.com/zoidbergconspiracy/hoard/route"
	"github.com/zoidbergconspiracy/hoard/log"
	"github.com/zoidbergconspiracy/hoard/auth"
	
)

func main() {

	hoard := route.HoardHandler{}

	srv := &http.Server{
		Addr:    "0.0.0.0:9443",
		Handler: log.NewLoggingHandler( auth.NewAuthHandler(&hoard, os.Stderr), os.Stdout),
	}

	tlsConfig := &tls.Config{}
	tlsConfig.ClientAuth = tls.RequestClientCert
	srv.TLSConfig = tlsConfig

	fmt.Println(srv.ListenAndServeTLS("server.crt", "server.key"))

}


