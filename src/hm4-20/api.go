package main

import (
	//"fmt"
	"html/template"
	"net/http"
)

var mytempl *template.Template

type Userstuff struct {
	Name, Email, Message string
}

func init() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/formhandler", formhandler)
  http.HandleFunc("/template", myTemplate)
  mytempl = template.Must(template.ParseFiles("mainForm.html", "link.html", "template.html"))
}

func exeTempl (w http.ResponseWriter, template string, data Userstuff) {
  err := mytempl.ExecuteTemplate(w, template+".html", data)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

func handler(w http.ResponseWriter, r *http.Request) {
  exeTempl(w, "mainForm", Userstuff{});
}

func formhandler(w http.ResponseWriter, r *http.Request) {
  theUser := Userstuff {r.FormValue("name"), r.FormValue("email"), r.FormValue("message")};
  exeTempl(w, "link", theUser);
}

func myTemplate(w http.ResponseWriter, r *http.Request) {
  theUser := Userstuff {r.FormValue("name"), r.FormValue("email"), r.FormValue("message")};
  exeTempl(w, "template", theUser);
}