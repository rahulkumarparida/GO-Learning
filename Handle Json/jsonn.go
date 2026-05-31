package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name string `json:"name"`
	Price int	`json:"price"`
	Platform string `json:"platform"`
	Password string	`json:"-"` // the - also helps ot remove the password to be sent even when anything is public
	Tags []string `json:"tags,omitempty"`
}

func main(){
	fmt.Println("Learning Json")
	// encodingJson()
	decodingJson()
}


func encodingJson(){

	currentStat := []course{
		{"React.js",500,"udemy","password",[]string{"Web-dev","js"}},
		{"Programming with python",799,"tutedude","password",[]string{"django","py", "full py library"}},
		{"NodeJs",90,"YouTube","password",[]string{"Backend","js"}},
		{"ChaiAurGo",90,"YouTube","password",[]string{"Cli","go"}},
	}

	courseJson , err := json.MarshalIndent(currentStat,"","  ")

	if err != nil {
		panic(err)
	}
	fmt.Println("Course Json: ", string(courseJson))

}

func decodingJson(){

	inp  := []byte(`
		{
    	"name": "ChaiAurGo",
    	"price": 90,
    	"platform": "YouTube",
		"password":"123jhagw",
    	"tags": ["Cli","go"]
	  }
	`)


	var jsonData course
	var jsonMapData map[string]interface{}
	json.Unmarshal(inp , &jsonData)
	json.Unmarshal(inp , &jsonMapData)

	// The jsonData did not show the password as it is taking refrence of course where password is ommited while jsonMapData shows all the field and key and valuse that come form the data.
	fmt.Println("jsonData: ", jsonMapData)
	for k,v := range jsonMapData{
		fmt.Println(k,v)
	}

	fmt.Printf("jsonData: %#v", jsonData)



}