package functions

import (
	"fmt"
	"html/template"
	"net/http"
)

type ErrorData struct {
	Error   string
	Message string
}

func ErrorPage(w http.ResponseWriter, errorText string, msg string, status int) {
	templ, err := template.ParseFiles("templates/errors.html")
	if err != nil {
		http.Error(w, msg, status)
		return
	}

	w.WriteHeader(status)
	data := ErrorData{
		Error:   errorText,
		Message: msg,
	}
	err = templ.Execute(w, data)
	if err != nil {
		fmt.Println(err)
	}
}
