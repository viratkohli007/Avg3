package main

import(
	// "db"
	// "fmt"
	// "log"
	// "net/http"
    "github.com/zenazn/goji"
    
	"flag"

	"controller/funcs"

)

const Port string =":8080"

func main() {
	// var mux = http.NewServeMux()
	startserver := flag.Bool("startserver", false, "starts the server")
	flag.Parse()

	if *startserver{
		
		LoadRoutes()
		StartServer()
	}
}

func StartServer(){

	flag.Set("bind", Port)
	goji.Serve()
}

func LoadRoutes(){
	// mux := goji.NewMux()
	// mux := web.New()
	
	goji.Post("/registration", funcs.Registration)
	goji.Post("/login", funcs.Login)
	
}