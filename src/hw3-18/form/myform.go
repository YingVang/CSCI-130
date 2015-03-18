package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", root)
	http.HandleFunc("/html", display)
	fmt.Println("listening...")
	err := http.ListenAndServe(GetPort(), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
		return
	}
}

func root(w http.ResponseWriter, r *http.Request) {
	rootForm, err := ioutil.ReadFile("html/rootF.html");
	if err != nil {
		fmt.Fprint(w, "404 Not Found")
	}
	fmt.Fprint(w, string(rootForm))
}

var results, _ = ioutil.ReadFile("html/display.html")
var resultsTemplate = template.Must(template.New("display").Parse(string(results)))

func display(w http.ResponseWriter, r *http.Request) {
	userStr := r.FormValue("str")
	if userStr == "Jerry Vang" {
		err := resultsTemplate.Execute(w, "Nice Job!!!")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
	} else {
		err := resultsTemplate.Execute(w, "That's not the correct name!")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		
		
	}
}

// Get the Port from the environment so we can run on Heroku
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}

