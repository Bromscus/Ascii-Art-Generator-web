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
	//if the path is neither a / or ascii-art we show this message 
	if r.URL.Path != "/" && r.URL.Path != "/ascii-art" {
		http.Error(w, "404: Status page not found Error", http.StatusNotFound)
		return
	}
//template allows Go to generate a HTML output
	templ, err := template.ParseFiles("templates/index.html")
	if err != nil {
		if os.IsNotExist(err) {
			http.Error(w, "404: Template Not Found", http.StatusNotFound)
			return
		}

		http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
		return
	}
	//Display the empty form when the home page is first opened.
	if r.URL.Path == "/" && r.Method == http.MethodGet {
		err = templ.Execute(w, PageData{})
		if err != nil {
			http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
		}
		return
	}
	//If the user sends a non post request to /ascii-art
	if r.URL.Path == "/ascii-art" && r.Method != http.MethodPost {
		http.Error(w, "405: Status Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	//text is the input in the text area
	//font is the banner
	text := r.FormValue("text")
	font := r.FormValue("styles")

	msg, valid := ascii.IsValid(text)
	if !valid {
		data := PageData{
			Font:   font,
			Result: msg,
		}
		//the server sends the https response header that has status code (can be viewed in the network section in the inspect of the browser)
		//a http response header is a metadata, in simple terms like a package of data
		w.WriteHeader(http.StatusBadRequest)
		//just to show the issue along with the status error code
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
			http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
			return
		}
		fontFile = "formats/standard.txt"

	case "shadow":
		if !ascii.CheckFile("formats/shadow.txt", ShadowHash) {
			http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
			return
		}
		fontFile = "formats/shadow.txt"

	case "thinkertoy":
		if !ascii.CheckFile("formats/thinkertoy.txt", ThinkertoyHash) {
			http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
			return
		}
		fontFile = "formats/thinkertoy.txt"

	default:

		http.Error(w, "400: Status Bad Request", http.StatusBadRequest)
		return
	}

	banner, err := os.ReadFile(fontFile)
	if err != nil {
		http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Font:   font,
		Result: ascii.Process(text, banner),
	}

	err = templ.Execute(w, data)
	if err != nil {
		http.Error(w, "500: Status Internal Server Error", http.StatusInternalServerError)
		return
	}
}
