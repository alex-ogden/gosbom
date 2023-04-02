package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func startServer(srv_host string) {
	// Define endpoints
	fileServer := http.FileServer(http.Dir(STATIC_DIR))
	http.Handle("/", fileServer)
	http.HandleFunc("/health", handleHealthCheck)
	http.HandleFunc("/upload", handleUpload)

	log.Printf("Starting web server on %s", srv_host)
	if err := http.ListenAndServe(srv_host, nil); err != nil {
		log.Fatal("Failed to start web server: ", err)
	}
}

func handleHealthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		log.Printf("Request method recieved was not GET\n")
		log.Printf("Request method: %s\n", r.Method)
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	log.Printf("Got request for /health")
	fmt.Fprintf(w, "Healthy")
}

func handleUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		log.Printf("Request method recieved was not POST\n")
		log.Printf("Request method: %s\n", r.Method)
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
		return
	}

	// Process request
	log.Print("Checking size of upload request")
	r.Body = http.MaxBytesReader(w, r.Body, int64(MAX_UPLOAD_SIZE))
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		log.Print("Upload request error: ", err)
		http.Error(w, "Upload request was too large", http.StatusRequestEntityTooLarge)
		return
	}

	// Get file from request
	log.Print("Getting file from upload request")
	file, _, err := r.FormFile("upload_file")
	if err != nil {
		log.Print("Error getting file from upload request: ", err)
		fmt.Fprint(w, "Error getting file from upload request: ", err)
	}

	// Read file into memory
	log.Print("Reading file into memory")
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Print("Error reading file from upload request: ", err)
		fmt.Fprint(w, "Error reading file from upload request: ", err)
	}

	// Parse file contents
	log.Print("Parsing file contents")
	parsedJson, err := parseFile(fileBytes)
	if err != nil {
		log.Print("Error parsing file from upload request: ", err)
		fmt.Fprint(w, "Error parsing file from upload request: ", err)
	}

	fmt.Fprintf(w, string(parsedJson))
}
