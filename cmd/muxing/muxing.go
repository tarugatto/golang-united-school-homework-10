package main

import (
	"fmt"
	"io"
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

	router.HandleFunc("/name/{PARAM}", handleGetParam).Methods(http.MethodGet)
	router.HandleFunc("/bad", handleGetBad).Methods(http.MethodGet)
	router.HandleFunc("/data", handlePostData).Methods(http.MethodPost)
	router.HandleFunc("/headers", handlePostHeaders).Methods(http.MethodPost)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func handleGetParam(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)["PARAM"]
	fmt.Fprintf(w, "Hello, %s!", param)
}

func handleGetBad(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func handlePostData(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("I got message:\n"))
	w.Write(data)
}

func handlePostHeaders(w http.ResponseWriter, r *http.Request) {
	a := r.Header.Get("a")
	b := r.Header.Get("b")

	c, _ := strconv.Atoi(a)
	d, _ := strconv.Atoi(b)

	w.Header().Add("a+b", strconv.Itoa(c+d))
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
