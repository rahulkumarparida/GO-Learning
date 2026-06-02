package main

import (
	"fmt"
	"log"
	"mongosetup/router"
	"net/http"
)


func main(){

	fmt.Println("MongoDB API")

	log.Fatal(http.ListenAndServe(":8000",router.Router()))
	fmt.Println("listening at http://localhost:8000")
}


