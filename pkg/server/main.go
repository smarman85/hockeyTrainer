package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	// "hockeyTrainer/pkg/fastHands"

	"github.com/gorilla/mux"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

type task struct {
	workout []int
}

func (h spaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// get the absolute path to prevent directory traversal
	path, err := filepath.Abs(r.URL.Path)
	fmt.Printf("Path1: %s\n", path)
	if err != nil {
		// if we failed to get the absolute path respond with a 400 bad request
		// and stop
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// prepend the path with the path to the static directory
	path = filepath.Join(h.staticPath, path)
	fmt.Printf("Path2: %s\n", path)

	// check whether a file exists at the given path
	_, err = os.Stat(path)
	fmt.Printf("Path3: %s\n", path)
	if os.IsNotExist(err) {
		// file does not exist, serve index.html
		// http.ServeFile(w, r, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		// if we got an error (that wasn't that the file doesn't exist) stating the
		// file, return a 500 internal server error and stop
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// otherwise, use http.FileServer to serve the static dir
	// fmt.Printf("staticPath: %s\n", h.staticPath)
	// http.FileServer(http.Dir(h.staticPath)).ServeHTTP(w, r)
	data := task{
		workout: []int{1, 3, 4},
	}
	tmpl := template.Must(template.ParseFiles("html/index.gohtml"))
	tmpl.Execute(w, data)
}

func StartServer() {
	router := mux.NewRouter()

	router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		// an example API handler
		json.NewEncoder(w).Encode(map[string]bool{"ok": true})
	})

	//spa := spaHandler{staticPath: "html", indexPath: "index.gohtml"}
	spa := spaHandler{staticPath: "html", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8999",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
