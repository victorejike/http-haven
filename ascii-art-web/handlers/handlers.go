package handlers

import (
	"ascii-art-web/ascii"
	"bytes"
	"errors"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type AsciiArtData struct {
	Input     string
	Banner    string
	Result    string
	Error     string
	HasResult bool
	FileName  string
}

var generator = ascii.NewAsciiGenerator()

func HandlerSwitch(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		HomeHandler(w, r)
	case "/ascii-art":
		AsciiArtHandler(w, r)
	case "/ascii-art-file":
		AsciiArtFileHandler(w, r)
	default:
		ErrorHandler(w, http.StatusNotFound, "Page not found")
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		ErrorHandler(w, http.StatusNotFound, "Page not found")
		return
	}

	if r.Method != http.MethodGet {
		ErrorHandler(w, http.StatusBadRequest, "Method not allowed")
		return
	}

	banner := bannerOrDefault(r.URL.Query().Get("banner"))
	if !generator.IsValidBanner(banner) {
		ErrorHandler(w, http.StatusBadRequest, "Invalid banner")
		return
	}

	if err := renderIndex(w, AsciiArtData{Banner: banner, HasResult: false}); err != nil {
		handleRenderError(w, err)
	}
}

func AsciiArtHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusBadRequest, "Method not allowed")
		return
	}

	err := r.ParseForm()
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, "Invalid form data")
		return
	}

	input := r.FormValue("text")
	banner := r.FormValue("banner")

	if input == "" {
		data := AsciiArtData{
			Banner: bannerOrDefault(banner),
			Error:  "Please enter text to convert",
		}
		if err := renderIndexWithStatus(w, http.StatusBadRequest, data); err != nil {
			handleRenderError(w, err)
		}
		return
	}

	if !ascii.ValidateInput(input) {
		data := AsciiArtData{
			Input:  input,
			Banner: bannerOrDefault(banner),
			Error:  "Unsupported character detected. Please use only ASCII printable characters.",
		}
		if err := renderIndexWithStatus(w, http.StatusBadRequest, data); err != nil {
			handleRenderError(w, err)
		}
		return
	}

	banner = bannerOrDefault(banner)
	if !generator.IsValidBanner(banner) {
		ErrorHandler(w, http.StatusBadRequest, "Invalid banner")
		return
	}

	result, err := generator.Generate(input, banner)
	if err != nil {
		if errors.Is(err, ascii.ErrBannerNotFound) {
			ErrorHandler(w, http.StatusNotFound, "Banner not found")
			return
		}
		data := AsciiArtData{
			Input:  input,
			Banner: banner,
			Error:  "Error generating ASCII art: " + err.Error(),
		}
		if err := renderIndexWithStatus(w, http.StatusInternalServerError, data); err != nil {
			handleRenderError(w, err)
		}
		return
	}

	data := AsciiArtData{
		Input:     input,
		Banner:    banner,
		Result:    result,
		HasResult: true,
	}
	if err := renderIndex(w, data); err != nil {
		handleRenderError(w, err)
	}
}

func AsciiArtFileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ErrorHandler(w, http.StatusBadRequest, "Method not allowed")
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		ErrorHandler(w, http.StatusBadRequest, "Error parsing form: "+err.Error())
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		data := AsciiArtData{
			Error: "Please select a file to upload",
		}
		if err := renderIndexWithStatus(w, http.StatusBadRequest, data); err != nil {
			handleRenderError(w, err)
		}
		return
	}
	defer file.Close()

	banner := r.FormValue("banner")
	if banner == "" {
		banner = "standard"
	}
	if !generator.IsValidBanner(banner) {
		ErrorHandler(w, http.StatusBadRequest, "Invalid banner")
		return
	}

	var content strings.Builder
	_, err = io.Copy(&content, file)
	if err != nil {
		data := AsciiArtData{
			Error: "Error reading file: " + err.Error(),
		}
		if err := renderIndexWithStatus(w, http.StatusInternalServerError, data); err != nil {
			handleRenderError(w, err)
		}
		return
	}

	input := content.String()

	if !ascii.ValidateInput(input) {
		data := AsciiArtData{
			Input:    input[:min(100, len(input))] + "...",
			FileName: header.Filename,
			Error:    "File contains unsupported characters. Please use only ASCII printable characters.",
		}
		if err := renderIndexWithStatus(w, http.StatusBadRequest, data); err != nil {
			handleRenderError(w, err)
		}
		return
	}

	file.Seek(0, 0)
	result, err := generator.GenerateFromReader(file, banner)
	if err != nil {
		if errors.Is(err, ascii.ErrBannerNotFound) {
			ErrorHandler(w, http.StatusNotFound, "Banner not found")
			return
		}
		data := AsciiArtData{
			FileName: header.Filename,
			Banner:   banner,
			Error:    "Error generating ASCII art: " + err.Error(),
		}
		if err := renderIndexWithStatus(w, http.StatusInternalServerError, data); err != nil {
			handleRenderError(w, err)
		}
		return
	}

	data := AsciiArtData{
		Input:     "File: " + header.Filename,
		Banner:    banner,
		Result:    result,
		HasResult: true,
		FileName:  header.Filename,
	}
	if err := renderIndex(w, data); err != nil {
		handleRenderError(w, err)
	}
}

func ErrorHandler(w http.ResponseWriter, status int, message string) {
	http.Error(w, message, status)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func renderIndex(w http.ResponseWriter, data AsciiArtData) error {
	return renderIndexWithStatus(w, http.StatusOK, data)
}

func renderIndexWithStatus(w http.ResponseWriter, status int, data AsciiArtData) error {
	tmpl, err := template.ParseFiles(indexTemplatePath())
	if err != nil {
		return err
	}

	var page bytes.Buffer
	if err := tmpl.Execute(&page, data); err != nil {
		return err
	}

	w.WriteHeader(status)
	_, err = w.Write(page.Bytes())
	return err
}

func indexTemplatePath() string {
	paths := []string{
		filepath.Join("templates", "index.html"),
		filepath.Join("..", "templates", "index.html"),
	}

	for _, path := range paths {
		if _, err := os.Stat(path); err == nil {
			return path
		}
	}

	return paths[0]
}

func bannerOrDefault(banner string) string {
	if banner == "" {
		return "standard"
	}
	return banner
}

func handleRenderError(w http.ResponseWriter, err error) {
	if os.IsNotExist(err) {
		ErrorHandler(w, http.StatusNotFound, "Template not found")
		return
	}
	ErrorHandler(w, http.StatusInternalServerError, "Internal server error")
}
