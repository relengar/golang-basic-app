package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// ServeStatic serves static files for frontend app
func ServeStatic(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("public/" + r.URL.Path)
	if err != nil {
		fmt.Printf("could not load file %s", err)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No file"))
	}

	defer f.Close()
	w.Header().Add("Content-Type", "text/html")
	io.Copy(w, f)
}
