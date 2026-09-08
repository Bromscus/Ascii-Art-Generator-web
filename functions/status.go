package functions

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"web/ascii"
)

type PageData struct {
	Result string
	Font   string
}

//
//templ.Execute basically means we write the final HTML into the HTTP response
//
func Handler(w http.ResponseWriter, r *http.Request) {
	// if the path is neither a / or ascii-art we show this message
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		ErrorPage(w, "404: page not found", "Please check that the Website address is spelled correctly or go to our home page.", http.StatusNotFound)

		return
	}
	// template allows Go to generate a HTML output
	if !ascii.CheckFile("templates/index.html", htmlHash) {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}
	// Display the empty form when the home page is first opened.
	if r.URL.Path == "/" && r.Method == http.MethodGet {
		err = templ.Execute(w, PageData{})
		if err != nil {
			ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		}
		return
	}
	// If the user sends a non post request to /ascii-art
	if r.URL.Path == "/ascii-art" && r.Method != http.MethodPost {
		ErrorPage(w, "405: Method Not Allowed", "This request method is not allowed for this page.", http.StatusMethodNotAllowed)
		return
	}
	// text is the input in the text area
	// font is the banner
	text := r.FormValue("text")
	font := r.FormValue("styles")

	msg, valid := ascii.IsValid(text)
	if !valid {
		data := PageData{
			Font:   font,
			Result: msg,
		}
		// the server sends the https response header that has code (can be viewed in the network section in the inspect of the browser)
		// a http response header is a metadata, in simple terms like a package of data
		w.WriteHeader(http.StatusBadRequest)
		// just to show the issue along with the status error code
		err = templ.Execute(w, data)
		if err != nil {
			fmt.Println(err)
		}
		return
	}

	fontFile := ""

	switch font {
	case "standard":
		if !ascii.CheckFile("formats/standard.txt", StandardHash) {
			ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
			return
		}
		fontFile = "formats/standard.txt"

	case "shadow":
		if !ascii.CheckFile("formats/shadow.txt", ShadowHash) {
			ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
			return
		}
		fontFile = "formats/shadow.txt"

	case "thinkertoy":
		if !ascii.CheckFile("formats/thinkertoy.txt", ThinkertoyHash) {
			ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
			return
		}
		fontFile = "formats/thinkertoy.txt"

	default:

		ErrorPage(w, "400: Bad Request", "The request could not be processed. Please check your input and try again.", http.StatusBadRequest)
		return
	}

	banner, err := os.ReadFile(fontFile)
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Font:   font,
		Result: ascii.Process(text, banner),
	}

	err = templ.Execute(w, data)
	if err != nil {
		ErrorPage(w, "500: Internal Server Error", "Something went wrong on the server.", http.StatusInternalServerError)
		return
	}
}
