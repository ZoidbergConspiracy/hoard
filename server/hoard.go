package main

import (
	"crypto/tls"
	"fmt"

	"io"
	"net/http"
	"os"
)

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)

	file, handler, er := r.FormFile("data")
	if er != nil {
		fmt.Printf("Early %v", er)
		return
	}
	defer file.Close()

	fmt.Fprintf(w, "%v", handler.Header)

	var f *os.File
	var err error

	if handler.Filename == "" {
		fmt.Println("No filename")
		f, err = os.OpenFile("./test/foo", os.O_WRONLY|os.O_CREATE, 0666)
	} else {

		f, err = os.OpenFile("./test/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	defer f.Close()
	io.Copy(f, file)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {

	if cs := r.TLS; len(cs.PeerCertificates) != 0 {
		for _, crt := range cs.PeerCertificates {
			fmt.Println("")
			fmt.Printf("Common name: %v\n", crt.Subject.CommonName)
			fmt.Printf("Subject Names: %v\n", crt.Subject.Names)
			fmt.Printf("Extra Names: %v\n", crt.Subject.ExtraNames)
			fmt.Printf("Signature: %X\n", crt.Signature)
		}
	}

}

func Default(w http.ResponseWriter, r *http.Request) {

  fmt.Println("URL    : " + r.URL.Path)
  fmt.Println("Method : " + r.Method)



  io.WriteString(w, "hello, world!\n")

}


func main() {

	http.HandleFunc("/", Default)

	srv := &http.Server{
		Addr:    "0.0.0.0:9443",
	}

	tlsConfig := &tls.Config{}
	tlsConfig.ClientAuth = tls.RequestClientCert
	srv.TLSConfig = tlsConfig


	fmt.Println(srv.ListenAndServeTLS("server.crt", "server.key"))

}
