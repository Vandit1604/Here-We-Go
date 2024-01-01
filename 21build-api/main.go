package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ---------------------- MODEL ----------------------
// model of the course - file
type Course struct {
	CourseId    string `json:"courseid"`
	CourseName  string `json:"coursename"`
	CoursePrice int    `json:"courseprice"`
	// figure out why it should be a pointer to Author not a copy
	Author *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db will be slice. NOTE: OLD DAYS WHEN I WAS LEARNING JAVA FOR JENKINS HEHE
var courses []Course

// middleware or helper functions. I knew helper functions - file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

// -------------------- CONTROLLERS -----------------------

func serveHome(w http.ResponseWriter, r *http.Request) {
	// I REPEAT EVERYTHING ON INTERNET IS IN FORM OF BYTEs THAT'S WHY SLICE OF BYTE
	w.Write([]byte(`<h1>WELCOME TO VANDIT'S FIRST API IN GOLANG. HUM COURSE BECHTEY HAI<h1/>`))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get All Courses here")
	// we can set the header of the response
	w.Header().Set("Content-Type", "application/json")

	// i was confused about the NewEncoder. Look here to understand clearly - https://stackoverflow.com/a/63480055
	// basically encoder does two things 1. marshalling 2. writing to stream or can be said writing as response(?)
	json.NewEncoder(w).Encode(courses)
}

// NOTE: In big production enviorment, we don't create API functions like these get all or get one functions. In prod, we create `courses` and they will be respond accordingly to give one or all or any n number of products. I'LL FIGURE OUT THIS ON MY OWN
func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get One Course")
	w.Header().Set("Content-Type", "application/json")

	// Vars returns all variables(if any) that we have after (?) like "https://example.com?search=hello" here hello is the param
	params := mux.Vars(r)
	// fmt.Println(params)

	// checks the param for the courseid and if it's present respond with the course
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("We couldn't find your course")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create One Course")
	w.Header().Set("Content-Type", "application/json")

	// what if request is empty with no body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	// what if the data is like this - {}
	var courseData Course
	json.NewDecoder(r.Body).Decode(&courseData)
	if courseData.IsEmpty() {
		json.NewEncoder(w).Encode("Please send some JSON data in body")
	}

	// TODO: if course name is same don't let it populate the slice
	// get the title from the request and loop in the courses.coursename if it's same encode msg to please send data with different name
	for _, course := range courses {
		if course.CourseName == courseData.CourseName {
			json.NewEncoder(w).Encode("Change course name")
		}
	}

	// we will assign the course a random ID
	// random, should be in string, get a random number convert it to a string
	courseData.CourseId = strconv.Itoa(rand.Intn(1000))
	courses = append(courses, courseData)
	fmt.Println(courses)
	return
}

// update a course - take the ID and use that to find the course that needs to be updated, remove the value and add it back with the same ID and updated values
func updateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update Course")
	w.Header().Set("Content-Type", "application/json")

	// what if request is empty with no body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	params := mux.Vars(r)
	// what if the data is like this - {}
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)
			var course Course
			json.NewDecoder(r.Body).Decode(&course)
			course.CourseId = params["id"]
			courses = append(courses, course)
			json.NewEncoder(w).Encode(course)
			return
		}
	}
	json.NewEncoder(w).Encode("The Course was not found")
}

// delete a course
func deleteACourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete a Course")
	w.Header().Set("Content-Type", "application/json")

	// what if request is empty with no body
	if r.Body == nil {
		json.NewEncoder(w).Encode("Please send some data")
	}

	params := mux.Vars(r)
	// what if the data is like this - {}
	var course Course
	json.NewDecoder(r.Body).Decode(&course)
	for index, course := range courses {
		if params["id"] == course.CourseId {
			courses = append(courses[:index], courses[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode("The Course was not found")
}

func main() {
	fmt.Println("API CHECK HELLO HELLO")

	// seeding
	courses = append(courses, Course{CourseId: "2", CourseName: "ReactJS", CoursePrice: 299, Author: &Author{Fullname: "Hitesh Choudhary", Website: "lco.dev"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "MERN Stack", CoursePrice: 199, Author: &Author{Fullname: "Hitesh Choudhary", Website: "go.dev"}})

	r := mux.NewRouter()
	// handling routes
	r.HandleFunc("/", serveHome).Methods("GET")
	r.HandleFunc("/courses", getAllCourses).Methods("GET")
	r.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	r.HandleFunc("/course", createOneCourse).Methods("POST")
	r.HandleFunc("/course/{id}", updateCourse).Methods("PUT")
	r.HandleFunc("/course/{id}", deleteACourse).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":4000", r))
}
