package functions

import (
	"crypto/sha256"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"web/ascii"
)

func ExportHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorPage(w, "405: Method Not Allowed", "This request method is not allowed for this page.", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		ErrorPage(w, "400: Bad Request", "The request could not be processed. Please check your input and try again.", http.StatusBadRequest)
		return
	}

	result, status, err := generateResult(r.FormValue("text"), r.FormValue("styles"))
	if err != nil {
		message := "The request could not be processed. Please check your input and try again."
		if status == http.StatusInternalServerError {
			message = "Something went wrong on the server."
		}
		ErrorPage(w, err.Error(), message, status)
		return
	}

	file := []byte(result)

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(file)))
	w.Header().Set("Content-Disposition", `attachment; filename="ascii-art.txt"`)

	if _, err := w.Write(file); err != nil {
		fmt.Println(err)
	}
}

func generateResult(text, font string) (string, int, error) {
	if msg, valid := ascii.IsValid(text); !valid {
		return "", http.StatusBadRequest,
			fmt.Errorf("400: Bad Request: %s", msg)
	}

	var fontFile string
	var hash [32]byte

	switch font {
	case "standard":
		fontFile = "formats/standard.txt"
		hash = StandardHash
	case "shadow":
		fontFile = "formats/shadow.txt"
		hash = ShadowHash
	case "thinkertoy":
		fontFile = "formats/thinkertoy.txt"
		hash = ThinkertoyHash
	default:
		return "", http.StatusBadRequest,
			fmt.Errorf("400: Bad Request")
	}

	banner, err := os.ReadFile(fontFile)
	if err != nil || sha256.Sum256(banner) != hash {
		return "", http.StatusInternalServerError,
			fmt.Errorf("500: Internal Server Error")
	}

	return ascii.Process(text, banner), http.StatusOK, nil
}
