package route

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"

	"github.com/codahale/blake2"
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

	case "POST", "PUT":
		UploadFile(w, r)

	case "DELETE":
		DeleteFile(w, r)
	}

}

func UploadFile(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Write file to ./data%v\n", r.URL.Path)

	dir := path.Dir("./data" + r.URL.Path)

	if os.MkdirAll(dir, os.ModeDir|os.ModePerm) != nil {
		http.Error(w, "Failed to create directory.", 500)
		return
	}

	var err error

	r.ParseMultipartForm(32 << 20)
	in, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Failed to read file from multipart form.", 500)
		return
	}
	defer in.Close()

	var out *os.File

	out, err = os.OpenFile("./data"+r.URL.Path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		http.Error(w, "Failed to write file.", 500)
		return
	}
	defer out.Close()

	hash := blake2.NewBlake2B()

	writer := io.MultiWriter(out, hash)

	io.Copy(writer, in)

	d := hash.Sum(nil)
	fmt.Fprintf(w, "Hash is %X\n", d)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {

	err := os.Remove("./data" + r.URL.Path)

	if err != nil {
		http.Error(w, "Failed to delete file.", 500)
		return
	}

	fmt.Fprintf(w, "Deleted file ./data%v\n", r.URL.Path)

}
