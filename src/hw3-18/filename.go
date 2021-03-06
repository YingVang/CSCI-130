package static

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"io/ioutil"
)

func init() {
	http.HandleFunc("/", static)
}

func static(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/"+r.URL.Path)
}



/*
code reference:
http://blog.extramaster.net/2014/05/creating-static-server-in-go-for-google.html
*/
