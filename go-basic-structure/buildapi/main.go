package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Model for course - file
type Course struct {
	CourseID    string  `json: "courseid"`
	CourseName  string  `json: "coursename"`
	CoursePrice int     `json: "price"`
	Author      *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Course

// middleware, helper
func (c *Course) IsEmpty() bool {
	return c.CourseID == "" && c.CourseName == ""
}

func main() {

}

// controllers

// serve home route

func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API by golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type","application/json")
	json.NewEncoder(w).Encode(courses)
}
