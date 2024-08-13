package ascii

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var Data = struct {
	Title, Description, Content,Warning string
}{
	Title:       "Ascii Art Web",
	Description: "Convert text to ASCII art.",
	Content:     "",
	Warning:     "",
}

var FilePath = "download/file.txt"

var text, font = "", ""

func PageHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404\nThis page could not be found.", http.StatusNotFound)
		return
	}
	template, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		log.Fatal(w, "Internal Server Error 500", http.StatusInternalServerError)
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
		return

	}
	err = template.Execute(w, Data)
	if err != nil {
		log.Fatal(err)
	}
	Data.Content, Data.Warning = "", ""
}

func Downloader(w http.ResponseWriter, r *http.Request) {
	// if r.URL.Path != "/download" {
	// 	http.NotFound(w, r)
	// 	return
	// }
	fmt.Println("Downloading file...")
	// Open the file for reading
	file, err := os.Open(FilePath)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
	defer file.Close()
	// Set the content type and the filename
	w.Header().Set("Content-Disposition", "attachment; filename=file.txt")
	w.Header().Set("Content-Type", "text/plain")
	// Copy the file contents to the response writer
	_, err = io.Copy(w, file)
	if err != nil {
		http.Error(w, "Internal Server Error 500", http.StatusInternalServerError)
		return
	}
}

func HandleAscii(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed 405", http.StatusMethodNotAllowed)
		return

	}

	font = r.FormValue("select")
	if !IsBanner(font) {
		http.Error(w, "Bad Request 400", http.StatusBadRequest)
		return
	}

	text = r.FormValue("text")
	if !IsPrintable(text) {
		Data.Warning = "The text contains non-printable characters."
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Get the value of a form field and assign it to the Data struct
	Data.Content = AsciiArt(text, font)
	// Then redirect the user to the root page
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
