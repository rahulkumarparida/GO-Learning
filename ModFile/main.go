package main

import (
	"fmt"
	"log"
	"net/http"
	"modulesFile/greetpage"
	"github.com/gorilla/mux"
)

func main()  {
	greeter()

	router := mux.NewRouter()
	
	router.HandleFunc("/",serveHome).Methods("GET")
	router.HandleFunc("/home",greetpage.HomeFile).Methods("GET")

	fmt.Println("Listening at :6969")	
	log.Fatal(http.ListenAndServe(":6969",router))



}


func greeter(){
	
	fmt.Println("Hello User")

}

func serveHome(w http.ResponseWriter,r  *http.Request){
	w.Write([]byte("<h1>Hello world this go with chai</h1>"))
}
