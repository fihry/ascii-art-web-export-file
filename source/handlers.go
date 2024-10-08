package ascii

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
)

var Data = struct {
	Title, Description, Content, Warning string
}{
	Title:       "Ascii Art Web",
	Description: "Convert text to ASCII art.",
	Content:     "",
	Warning:     "",
}

var StatusError = struct {
	Status string
	Code   int
}{
	Status: "OK",
	Code:   200,
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	tmplPath := filepath.Join("web/template", tmpl)
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("Error parsing template:", err)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		log.Println("Error executing template:", err)
	}
}

func PageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		StatusError.Status = "Not Found"
		StatusError.Code = http.StatusNotFound
		w.WriteHeader(http.StatusNotFound)
		renderTemplate(w, "Error.html", StatusError)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	renderTemplate(w, "index.html", Data)
	Data.Content, Data.Warning = "", ""
}

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		StatusError.Status = "Method Not Allowed"
		StatusError.Code = http.StatusMethodNotAllowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		renderTemplate(w, "Error.html", StatusError)
		return
	}

	r.ParseForm()
	font := r.FormValue("select")
	if !IsBanner(font) {
		StatusError.Status = "Bad Request"
		StatusError.Code = http.StatusBadRequest
		w.WriteHeader(http.StatusBadRequest)
		renderTemplate(w, "Error.html", StatusError)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	text := r.FormValue("text")
	if !IsPrintable(text) {
		Data.Warning = "The text contains non-printable characters."
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	} else if len(text) > 200 {
		Data.Warning = "The text is too long."
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	Data.Content = AsciiArt(text, font)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/download" {
		StatusError.Status = "Not Found"
		StatusError.Code = http.StatusNotFound
		renderTemplate(w, "Error.html", StatusError)
		return
	}
	Content := r.URL.Query().Get("content")
	if len(Content) == 0 {
		Data.Warning = "Type some text and click on the 'generate' button."
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
	w.Header().Set("Content-length", strconv.Itoa(len(Content)))
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(Content))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
