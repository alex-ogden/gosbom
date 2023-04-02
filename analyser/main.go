package main

import (
	"log"
	"os"
)

// Max upload size 3MB
const MAX_UPLOAD_SIZE int64 = 3 * 1024 * 1024
const STATIC_DIR string = "../static"

func main() {
	// Define host and port
	log.Print("Setting host and port values\n")
	var srv_host string = "0.0.0.0"
	var srv_port string = os.Getenv("PORT")

	if srv_port == "" {
		log.Print("Couldn't find environment variable: PORT\n")
		log.Print("Defaulting to: 4433\n")
		srv_port = "4433"
	}
	log.Printf("Host: %s\n", srv_host)
	log.Printf("Port: %s\n", srv_port)

	// Start server
	srv_host = srv_host + ":" + srv_port
	startServer(srv_host)
}
