package greetpage

import (
	"net/http"
	"os"

)



func HomeFile(w http.ResponseWriter , r *http.Request){

	// files ,_ := os.ReadDir(".")

	text , err := os.ReadFile("/home/rahulroxx/MyHDD/NewVolume/Study/Basic_GO/ModFile/greetpage/index.html")
	if err != nil {
		panic(err)
	}


	w.Write(text) // required to be in byte

}