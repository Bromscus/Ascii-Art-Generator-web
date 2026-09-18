package functions

import (
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
		ErrorPage(w, err.Error(), "The request could not be processed. Please check your input and try again.", status)
		return
	}

	file := []byte(result)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(file)))
	w.Header().Set("Content-Disposition", `attachment; filename="ascii-art.txt"`)
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(file); err != nil {
		fmt.Println(err)
	}
}

func generateResult(text, font string) (string, int, error) {
	if msg, valid := ascii.IsValid(text); !valid {
		return "", http.StatusBadRequest, fmt.Errorf("400: Bad Request: %s", msg)
	}

	fontFiles := map[string]string{
		"standard":   "formats/standard.txt",
		"shadow":     "formats/shadow.txt",
		"thinkertoy": "formats/thinkertoy.txt",
	}
	fontFile, ok := fontFiles[font]
	if !ok {
		return "", http.StatusBadRequest, fmt.Errorf("400: Bad Request")
	}

	banner, err := os.ReadFile(fontFile)
	if err != nil || !ascii.CheckFile(fontFile, expectedHash(font)) {
		return "", http.StatusInternalServerError, fmt.Errorf("500: Internal Server Error")
	}

	return ascii.Process(text, banner), http.StatusOK, nil
}

func expectedHash(font string) [32]byte {
	switch font {
	case "standard":
		return StandardHash
	case "shadow":
		return ShadowHash
	default:
		return ThinkertoyHash
	}
}
