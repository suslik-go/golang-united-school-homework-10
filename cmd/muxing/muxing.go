package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/name/{name}", NameHandler).Methods(http.MethodGet)
	router.HandleFunc("/bad", BadHandler).Methods(http.MethodGet)
	router.HandleFunc("/data/{body}/{param}", DataHandler).Methods(http.MethodPost)
	router.HandleFunc("/headers/{headers}", HeadersHandler).Methods(http.MethodPost)
	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}

}

func NameHandler(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	fmt.Fprint(w, "Hello, "+params["name"]+"!")
}

func BadHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func DataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	w.Write([]byte(params["body"] + ":\n" + params["param"]))
}

func HeadersHandler(w http.ResponseWriter, r *http.Request) {
	var postHeaders map[string]string
	params := mux.Vars(r)
	json.Unmarshal([]byte(params["headers"]), &postHeaders)
	for key, value := range postHeaders {
		w.Header().Set(key, value)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
