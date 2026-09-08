package functions

import (
	"fmt"
	"html/template"
	"net/http"
)



func ErrorPage(w http.ResponseWriter, msg string, status int) {
	templ, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(w, msg, status)
		return
	}
	
	w.WriteHeader(status)

	err = templ.Execute(w, msg)
	if err != nil {
		fmt.Println(err)
	}
}
