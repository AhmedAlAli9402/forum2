package main

import (
	"fmt"
	forum "forum/webApp/handlers"
	"log"
	"net/http"
)

// Whatever needs to load before the server starts (Files/APIs)
func init() {
	forum.StaticFileLoader()
}

func main() {
	const port = ":8080"
	http.HandleFunc("/", forum.MainHandler)
	fmt.Println("http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
