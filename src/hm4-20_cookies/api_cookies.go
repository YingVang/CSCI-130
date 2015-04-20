package main

import (
	//"fmt"
	"html/template"
	"net/http"
	"time"
)

var mytempl *template.Template

type Userstuff struct {
	Name, Email, Message string
}

func init() {
  http.HandleFunc("/", handler)
  http.HandleFunc("/formhandler", formhandler)
  http.HandleFunc("/msgCookie", msgTempl)
  http.HandleFunc("/template", msgCookieTempl)
  mytempl = template.Must(template.ParseFiles("mainForm.html", "link.html", "msgTempl.html", "msgCookieTempl.html"))
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
  	  expire := time.Now().AddDate(0, 0, 1)

      theUser := Userstuff {r.FormValue("name"), r.FormValue("email"), r.FormValue("message")}
      cookie := http.Cookie{
          Name:       "myTest1",
          Value:      theUser.Message,
          Expires:    expire,
          RawExpires: expire.Format(time.UnixDate),
          MaxAge:   86400,
          Secure:   false,
          HttpOnly: false,
      };
      http.SetCookie(w, &cookie)
      exeTempl(w, "link", theUser);
}

func msgTempl(w http.ResponseWriter, r *http.Request) {
  theUser := Userstuff {r.FormValue("name"), r.FormValue("email"), r.FormValue("message")};
  exeTempl(w, "msgTempl", theUser);  
}

func msgCookieTempl(w http.ResponseWriter, r *http.Request) {
    cookie, _ := r.Cookie("myTest1")
    theUser := Userstuff {"", "", cookie.Value}
    exeTempl(w, "msgCookieTempl", theUser);
}

