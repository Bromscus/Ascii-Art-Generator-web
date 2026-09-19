package functions

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	recorder := httptest.NewRecorder()

	Handler(recorder, req)

	if recorder.Code != http.StatusNotFound {
		t.Fatalf("expected status %d, got %d", http.StatusNotFound, recorder.Code)
	}
}

func TestHandlerHomePage(t *testing.T) {
	t.Chdir("..")
	LoadHashes()
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	recorder := httptest.NewRecorder()

	Handler(recorder, req)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, recorder.Code)
	}
}

func TestExportFormUsesPostMethod(t *testing.T) {
	t.Chdir("..")
	page, err := os.ReadFile("templates/index.html")
	if err != nil {
		t.Fatalf("read index template: %v", err)
	}

	content := string(page)
	if !strings.Contains(content, "formaction=\"/export\"") || !strings.Contains(content, "formmethod=\"post\"") {
		t.Fatal("export form must submit to /export with method=post")
	}
}

func TestExportHandlerUsesErrorPageForBadInput(t *testing.T) {
	t.Chdir("..")
	LoadHashes()
	form := url.Values{}
	form.Set("text", "")
	form.Set("styles", "standard")

	req := httptest.NewRequest(http.MethodPost, "/export", strings.NewReader(form.Encode()))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	recorder := httptest.NewRecorder()

	ExportHandler(recorder, req)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected status %d, got %d", http.StatusBadRequest, recorder.Code)
	}
	body := recorder.Body.String()
	if !strings.Contains(body, "error-result") || !strings.Contains(body, "400: Bad Request") {
		t.Fatalf("expected shared error page for invalid export, got body: %s", body)
	}
}
