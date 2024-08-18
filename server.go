package main

import (
	ascii "ascii/source"
	"log"
	"net/http"
	"os"
)

func main() {
	PORT := "3100"
	if len(os.Args) == 2 {
		PORT = os.Args[1]
	}
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./web/css/"))))
	http.HandleFunc("/", ascii.PageHandler)
	http.HandleFunc("/ascii-art", ascii.HandleAscii)
	http.HandleFunc("/download", ascii.DownloadHandler)
	log.Println("\033[32mServer is running on port " + PORT + "...🚀\033[0m")
	log.Println("\033[32mhttp://localhost:" + PORT + "\033[0m")
	http.ListenAndServe(":"+PORT, nil)
}
