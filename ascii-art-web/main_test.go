package main

import (
	"ascii-art-web/ascii"
	"ascii-art-web/handlers"
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  bool
	}{
		{name: "valid ASCII", input: "Hello World!", want: true},
		{name: "valid with newline", input: "Hello\nWorld", want: true},
		{name: "valid empty string", input: "", want: true},
		{name: "invalid byte", input: "Hello\x80World", want: false},
		{name: "unicode character", input: "Hello World!", want: true},
		{name: "emoji", input: "Hello👋World", want: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ascii.ValidateInput(tt.input); got != tt.want {
				t.Fatalf("ValidateInput(%q) = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}

func TestBannerAPIs(t *testing.T) {
	for _, banner := range []string{"standard", "shadow", "thinkertoy"} {
		if !ascii.IsValidBanner(banner) {
			t.Fatalf("IsValidBanner(%q) = false, want true", banner)
		}
	}

	if ascii.IsValidBanner("invalid") {
		t.Fatal("IsValidBanner(\"invalid\") = true, want false")
	}

	banners := ascii.GetAvailableBanners()
	if len(banners) != 3 {
		t.Fatalf("GetAvailableBanners() returned %d banners, want 3", len(banners))
	}

	found := map[string]bool{"standard": false, "shadow": false, "thinkertoy": false}
	for _, banner := range banners {
		found[banner] = true
	}
	for banner, ok := range found {
		if !ok {
			t.Fatalf("GetAvailableBanners() did not include %q", banner)
		}
	}
}

func TestGenerateAsciiArt(t *testing.T) {
	result, err := ascii.GenerateAsciiArt("A", "standard")
	if err != nil {
		t.Fatalf("GenerateAsciiArt() returned error: %v", err)
	}
	if !looksLikeAsciiArt(result) {
		t.Fatalf("GenerateAsciiArt() result does not look like ASCII art:\n%s", result)
	}

	result, err = ascii.GenerateAsciiArt("", "standard")
	if err != nil {
		t.Fatalf("GenerateAsciiArt(empty) returned error: %v", err)
	}
	if result != "" {
		t.Fatalf("GenerateAsciiArt(empty) = %q, want empty string", result)
	}

	result, err = ascii.GenerateAsciiArt("\n", "standard")
	if err != nil {
		t.Fatalf("GenerateAsciiArt(newline) returned error: %v", err)
	}
	if result != "\n" {
		t.Fatalf("GenerateAsciiArt(newline) = %q, want newline", result)
	}

	if _, err := ascii.GenerateAsciiArt("A", "invalid"); err == nil {
		t.Fatal("GenerateAsciiArt() with invalid banner returned nil error")
	}
}

func TestAsciiGenerator(t *testing.T) {
	generator := ascii.NewAsciiGenerator()

	for _, banner := range []string{"standard", "shadow", "thinkertoy"} {
		result, err := generator.GenerateFromReader(strings.NewReader("Hello"), banner)
		if err != nil {
			t.Fatalf("GenerateFromReader(%q) returned error: %v", banner, err)
		}
		if !looksLikeAsciiArt(result) {
			t.Fatalf("GenerateFromReader(%q) result does not look like ASCII art:\n%s", banner, result)
		}
	}

	if _, err := generator.GenerateFromReader(strings.NewReader("Hello"), "invalid"); err == nil {
		t.Fatal("GenerateFromReader() with invalid banner returned nil error")
	}

	var buf bytes.Buffer
	if err := generator.GenerateStreaming("Hello", "standard", &buf); err != nil {
		t.Fatalf("GenerateStreaming() returned error: %v", err)
	}
	if !looksLikeAsciiArt(buf.String()) {
		t.Fatalf("GenerateStreaming() result does not look like ASCII art:\n%s", buf.String())
	}

	size, err := generator.GetBannerSize("standard")
	if err != nil {
		t.Fatalf("GetBannerSize() returned error: %v", err)
	}
	if size <= 0 {
		t.Fatalf("GetBannerSize() = %d, want positive size", size)
	}

	if _, err := generator.GetBannerSize("invalid"); err == nil {
		t.Fatal("GetBannerSize() with invalid banner returned nil error")
	}
}

func TestReaderHelpers(t *testing.T) {
	lines, err := ascii.ReadBannerLines(strings.NewReader("line1\nline2\nline3"))
	if err != nil {
		t.Fatalf("ReadBannerLines() returned error: %v", err)
	}
	if got, want := strings.Join(lines, ","), "line1,line2,line3"; got != want {
		t.Fatalf("ReadBannerLines() = %q, want %q", got, want)
	}

	processed, err := ascii.ProcessInputReader(strings.NewReader("Hello\r\nWorld\nTest"))
	if err != nil {
		t.Fatalf("ProcessInputReader() returned error: %v", err)
	}
	if got, want := strings.Join(processed, ","), "Hello,World,Test"; got != want {
		t.Fatalf("ProcessInputReader() = %q, want %q", got, want)
	}
}

func TestHomeHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	handlers.HomeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("HomeHandler() status = %d, want %d", rr.Code, http.StatusOK)
	}
	if !strings.Contains(rr.Body.String(), "ASCII Art Web") {
		t.Fatal("HomeHandler() response did not include page title")
	}
}

func TestHomeHandlerBannerLinks(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/?banner=shadow", nil)
	rr := httptest.NewRecorder()

	handlers.HomeHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("HomeHandler() status = %d, want %d", rr.Code, http.StatusOK)
	}
	if !strings.Contains(rr.Body.String(), `href="/?banner=shadow"`) {
		t.Fatal("HomeHandler() response did not include shadow banner link")
	}
	if !strings.Contains(rr.Body.String(), `value="shadow" checked`) {
		t.Fatal("HomeHandler() did not preselect shadow banner")
	}
}

func TestHomeHandlerErrors(t *testing.T) {
	tests := []struct {
		name   string
		method string
		path   string
		want   int
	}{
		{name: "not found", method: http.MethodGet, path: "/missing", want: http.StatusNotFound},
		{name: "method not allowed", method: http.MethodPost, path: "/", want: http.StatusBadRequest},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.path, nil)
			rr := httptest.NewRecorder()

			handlers.HomeHandler(rr, req)

			if rr.Code != tt.want {
				t.Fatalf("HomeHandler() status = %d, want %d", rr.Code, tt.want)
			}
		})
	}
}

func TestHandlerSwitch(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want int
	}{
		{
			name: "home",
			req:  httptest.NewRequest(http.MethodGet, "/", nil),
			want: http.StatusOK,
		},
		{
			name: "ascii art",
			req: func() *http.Request {
				req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader("text=Go&banner=standard"))
				req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
				return req
			}(),
			want: http.StatusOK,
		},
		{
			name: "missing",
			req:  httptest.NewRequest(http.MethodGet, "/missing", nil),
			want: http.StatusNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()

			handlers.HandlerSwitch(rr, tt.req)

			if rr.Code != tt.want {
				t.Fatalf("HandlerSwitch() status = %d, want %d", rr.Code, tt.want)
			}
		})
	}
}

func TestAsciiArtHandler(t *testing.T) {
	tests := []struct {
		name       string
		form       string
		wantStatus int
		wantBody   string
	}{
		{name: "valid post", form: "text=Hello&banner=standard", wantStatus: http.StatusOK, wantBody: "ascii-result"},
		{name: "empty input", form: "text=&banner=standard", wantStatus: http.StatusBadRequest, wantBody: "Please enter text to convert"},
		{name: "invalid input", form: "text=Hello\x80World&banner=standard", wantStatus: http.StatusBadRequest, wantBody: "Unsupported character"},
		{name: "invalid banner", form: "text=Hello&banner=invalid", wantStatus: http.StatusBadRequest, wantBody: "Invalid banner"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader(tt.form))
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			rr := httptest.NewRecorder()

			handlers.AsciiArtHandler(rr, req)

			if rr.Code != tt.wantStatus {
				t.Fatalf("AsciiArtHandler() status = %d, want %d", rr.Code, tt.wantStatus)
			}
			if !strings.Contains(rr.Body.String(), tt.wantBody) {
				t.Fatalf("AsciiArtHandler() response did not include %q", tt.wantBody)
			}
		})
	}

	req := httptest.NewRequest(http.MethodGet, "/ascii-art", nil)
	rr := httptest.NewRecorder()
	handlers.AsciiArtHandler(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("AsciiArtHandler(GET) status = %d, want %d", rr.Code, http.StatusBadRequest)
	}
}

func TestAsciiArtFileHandler(t *testing.T) {
	body, contentType := multipartBody(t, "file", "hello.txt", "Hello", map[string]string{"banner": "standard"})
	req := httptest.NewRequest(http.MethodPost, "/ascii-art-file", body)
	req.Header.Set("Content-Type", contentType)
	rr := httptest.NewRecorder()

	handlers.AsciiArtFileHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("AsciiArtFileHandler() status = %d, want %d", rr.Code, http.StatusOK)
	}
	if !strings.Contains(rr.Body.String(), "hello.txt") || !strings.Contains(rr.Body.String(), "ascii-result") {
		t.Fatal("AsciiArtFileHandler() response did not include file result")
	}

	req = httptest.NewRequest(http.MethodGet, "/ascii-art-file", nil)
	rr = httptest.NewRecorder()
	handlers.AsciiArtFileHandler(rr, req)
	if rr.Code != http.StatusBadRequest {
		t.Fatalf("AsciiArtFileHandler(GET) status = %d, want %d", rr.Code, http.StatusBadRequest)
	}
}

func TestFullWebFlow(t *testing.T) {
	homeReq := httptest.NewRequest(http.MethodGet, "/", nil)
	homeRR := httptest.NewRecorder()
	handlers.HomeHandler(homeRR, homeReq)
	if homeRR.Code != http.StatusOK {
		t.Fatalf("home page status = %d, want %d", homeRR.Code, http.StatusOK)
	}

	formReq := httptest.NewRequest(http.MethodPost, "/ascii-art", strings.NewReader("text=Test&banner=standard"))
	formReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	formRR := httptest.NewRecorder()
	handlers.AsciiArtHandler(formRR, formReq)

	if formRR.Code != http.StatusOK {
		t.Fatalf("form submission status = %d, want %d", formRR.Code, http.StatusOK)
	}
	if !strings.Contains(formRR.Body.String(), "Test") || !strings.Contains(formRR.Body.String(), "ascii-result") {
		t.Fatal("form submission response did not include submitted text and ASCII result")
	}
}

func looksLikeAsciiArt(result string) bool {
	return strings.Contains(result, "|") || strings.Contains(result, "_")
}

func multipartBody(t *testing.T, fieldName, fileName, content string, fields map[string]string) (io.Reader, string) {
	t.Helper()

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	for name, value := range fields {
		if err := writer.WriteField(name, value); err != nil {
			t.Fatalf("WriteField(%q) failed: %v", name, err)
		}
	}

	part, err := writer.CreateFormFile(fieldName, fileName)
	if err != nil {
		t.Fatalf("CreateFormFile() failed: %v", err)
	}
	if _, err := part.Write([]byte(content)); err != nil {
		t.Fatalf("writing multipart file failed: %v", err)
	}
	if err := writer.Close(); err != nil {
		t.Fatalf("closing multipart writer failed: %v", err)
	}

	return &body, writer.FormDataContentType()
}
