package route

import (
	"net/http"
)

type HoardHandler struct {
	http.Handler
}

func (_ *HoardHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	default:
		http.Error(w, "Method not allowed", 405)

	case "GET", "HEAD":
		http.FileServer(http.Dir("./data")).ServeHTTP(w, r)

	}
}
