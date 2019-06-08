package main

import(
	// "db"
	// "fmt"
	"log"
	"net/http"
	"controller/funcs"
)

func main() {
	// var mux = http.NewServeMux()
	Port := ":8080"
	http.HandleFunc("/", SayHello)
	http.HandleFunc("/registration", funcs.Registration)
	err := http.ListenAndServe(Port, nil)
	if err != nil {
		log.Println("Server connection error", err)
	}

}

func SayHello(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}