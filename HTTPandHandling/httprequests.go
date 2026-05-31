package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

const URL = "https://github.com:5000/rahulkumarparida?tab=repositories"
// const URL = "https://air-bnb-client.vercel.app/"


func main()  {
	// webrequest()
	// exploreResponse()
	postFomData()


}

func webrequest(){
	fmt.Println("Http GET request")

	response , err := http.Get(URL)

	if err != nil {
		panic(err)
	
	}

	fmt.Printf("Respone tyep %T\n", response)

	defer response.Body.Close()

	body , err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Body: ", string(body))
}

func exploreResponse(){
	result , _ := url.Parse(URL)

	fmt.Println("Scheme: ", result.Scheme)
	fmt.Println("Host: ", result.Host)
	fmt.Println("Path: ", result.Path)
	fmt.Println("Rawquery: ", result.RawQuery)
	fmt.Println("Port: ", result.Port())

	questyP := result.Query()
	// fmt.Println("Query: ", questyP)
	fmt.Println("source: ", questyP["source"])

	for _ , val := range questyP{
		fmt.Println("Params: ",val)
	}

	 buildUrl := &url.URL{
		Scheme: "https",
		Host: "github.com",
		Path: "/rahulkumarparida",
		// RawQuery: "tab=repositories",
	 }

	 fmt.Println("Built: ", buildUrl)

	

	var newurl = buildUrl.String()
	response , err := http.Get(newurl)

	if err != nil {
		panic(err)
	
	}

	fmt.Printf("Respone tyep %T\n", response)

	defer response.Body.Close()

	body , err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println("Body: ", string(body))
	 

}



func postFomData(){
	Url :="http://localhost:8000/formData"
	data := url.Values{}
	data.Add("name","Rahul")
	data.Add("mental health","deterioting")
	data.Add("doing","Nothing")

	res ,_:= http.PostForm(Url,data)	
	defer res.Body.Close()

	content , _ := io.ReadAll(res.Body)
	fmt.Println("res:", string(content))

}

