//go:generate statik -src=./client/build
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/rakyll/statik/fs"
	log "github.com/sirupsen/logrus"
	_ "github.com/tarkalabs/embedded_resources/statik"

	"github.com/gorilla/mux"
)

var useStatik bool

func init() {
	flag.BoolVar(&useStatik, "use-statik", false, "serve embedded static resources")
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

type greeting struct {
	Message string `json:"message"`
}

func sayHello(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Add("Content-Type", "application/json")
	encoder := json.NewEncoder(rw)
	_ = encoder.Encode(greeting{Message: "Hello Gophercon"})
}

func main() {
	flag.Parse()
	router := mux.NewRouter()
	router.HandleFunc("/api/hello", sayHello).Methods("GET")
	port := os.Getenv("PORT")
	if port == "" {
		port = "4545"
	}
	statikFS, err := fs.New()
	if err != nil {
		log.Error("Unable to initialize statik file system")
	}
	router.PathPrefix("/").Handler(http.FileServer(statikFS))
	log.Infof("Starting server on port %s", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), router)
}
