package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Models
type Courses struct {
	CourseId string `json:"courseid"`
	CourseName string `json:"coursename"`
	Price int `json:"price"`
	Author *Author `json:"author"`
}


type Author struct {
	Fullname string `json:"fullname"`
	Website string `json:"website"`
}



// Middleware or helper functions
func (c *Courses) IsEmpty() bool {

	return  c.Price == 0 && c.CourseName == ""
}

var courses []Courses


func main(){
	 author1 := &Author{Fullname: "John Doe", Website: "https://johndoe.dev"}
	 author2 := &Author{Fullname: "Jane Smith", Website: "https://janesmith.io"}
	 author3 := &Author{Fullname: "Rahul Roxx", Website: "https://rahulroxx.com"}



	router := mux.NewRouter()

	courses = append(courses, Courses{CourseId: "C001", CourseName: "Complete Go BootCamp", Price: 49, Author: author1})	
	courses = append(courses, Courses{CourseId: "C002", CourseName: "Advanced Microservices in Go", Price: 89, Author: author1})
	courses = append(courses, Courses{CourseId: "C003", CourseName: "Docker & Kubernetes for Devs", Price: 59, Author: author1})
	courses = append(courses , Courses{CourseId: "C004", CourseName: "Mastering React & TypeScript", Price: 39, Author: author2})
	courses = append(courses , Courses{CourseId: "C005", CourseName: "Next.js FullStack Applications", Price: 69, Author: author2})
	courses = append(courses , Courses{CourseId: "C006", CourseName: "Data Structures & Algorithms", Price: 29, Author: author2})
	courses = append(courses , Courses{CourseId: "C007", CourseName: "System Design for Scale", Price: 119, Author: author3})
	courses = append(courses , Courses{CourseId: "C008", CourseName: "Building High-Performance APIs", Price: 79, Author: author3})
	courses = append(courses , Courses{CourseId: "C009", CourseName: "Linux Terminal Mastery", Price: 19, Author: author3})
	courses = append(courses , Courses{CourseId: "C010", CourseName: "Git & GitHub Collaboration", Price: 0, Author: author3})

	router.HandleFunc("/", healthCheck).Methods("GET")
	router.HandleFunc("/courses", getAllCourse).Methods("GET")	
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", addOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":8000", router))
}

// Controllers 

func healthCheck(w http.ResponseWriter , r *http.Request){
	fmt.Println("Healthy and Working")
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Status","OK")
	json.NewEncoder(w).Encode("{'message':'api is running and is healthy'}")

	defer r.Body.Close()
}

func getAllCourse(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("Healthy and Working")
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Status","OK")
	json.NewEncoder(w).Encode(courses)
	defer r.Body.Close()
}

func getOneCourse(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("Healthy and Working")
	w.Header().Set("Content-Type","application/json")
	w.Header().Set("Status","OK")

	params := mux.Vars(r)
	fmt.Println("Params: ", params)

	courseId := params["id"]

	
	defer r.Body.Close()
	
	for _ , val := range courses{
		if val.CourseId == courseId{
			json.NewEncoder(w).Encode(val)	
			return
		}
	}

	json.NewEncoder(w).Encode("{'message':'No relative couse found.'}")

}


func addOneCourse(w http.ResponseWriter , r *http.Request){
	fmt.Println("Add course")
	w.Header().Set("Content-Type","application/json")

	if r.Body == nil {
		json.NewEncoder(w).Encode("{'msg':'please fill out the data'}")
	}

	var newCourse Courses;

	_ = json.NewDecoder(r.Body).Decode(&newCourse)

	for _ , val := range courses{
		if val.CourseName == newCourse.CourseName{
			json.NewEncoder(w).Encode("{'msg':'the Coursename already exists'}")	

			return
		}
	}

	if newCourse.IsEmpty(){
		fmt.Println("Empty")
		json.NewEncoder(w).Encode("{'msg':'please fill out the form correctly'}")	
	}
	id := time.Now().UnixNano()
	newCourse.CourseId = string(id)

	courses = append(courses, newCourse)
	

	defer r.Body.Close()
	json.NewEncoder(w).Encode(newCourse)
	return
}

// Works only for PUT cause we are rewriting the course completely
func updateOneCourse(w http.ResponseWriter , r *http.Request)  {
	fmt.Println("Update the valuses of a course")
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	fmt.Println("Params: ", params )
	id := params["id"]

	
	
	for idx , val := range courses{
		
		if id == val.CourseId {
			courses = append(courses[:idx], courses[idx+1:]...)
			var course Courses;
			_ = json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = id
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			
			return
		}
	
	}
	json.NewEncoder(w).Encode("{'msg':'The course not found'}")
	return

}




func deleteOneCourse(w http.ResponseWriter , r *http.Request)  {
fmt.Println("Update the valuses of a course")
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	fmt.Println("Params: ", params )
	id := params["id"]

	for idx , val := range courses{
		
		if id == val.CourseId {
			courses = append(courses[:idx], courses[idx+1:]...)
			json.NewEncoder(w).Encode("{'msg':'The course sucessfully deleted'}")
			return
		}
	
	}
	json.NewEncoder(w).Encode("{'msg':'The course not found'}")
	return

	

}
