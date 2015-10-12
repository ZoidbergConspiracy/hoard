package auth

import (
	"fmt"
	"io"
	"net/http"
)

type AuthHandler struct {
	handler http.Handler
	out     io.Writer
}

func (h *AuthHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	if cs := r.TLS; len(cs.PeerCertificates) != 0 {
		for _, crt := range cs.PeerCertificates {
			fmt.Fprintf(h.out, "Common name: %v\n", crt.Subject.CommonName)
			fmt.Fprintf(h.out, "Subject Names: %v\n", crt.Subject.Names)
			fmt.Fprintf(h.out, "Extra Names: %v\n", crt.Subject.ExtraNames)
			fmt.Fprintf(h.out, "Signature: %X\n", crt.Signature)
			fmt.Fprintf(h.out, "Signature: %X\n", crt.SignatureAlgorithm)
			fmt.Fprintf(h.out, "KeyID: %X\n", crt.SubjectKeyId)
		}
	}

	h.handler.ServeHTTP(w, r)

}

func NewAuthHandler(handler http.Handler, out io.Writer) http.Handler {
	return &AuthHandler{
		handler: handler,
		out:     out,
	}
}
