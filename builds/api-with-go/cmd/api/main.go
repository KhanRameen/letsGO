package main

import (
	"net/http"
	"fmt"

	"api-with-go/internal/handler/handler"

	"github.com/go-chi/chi"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true) //a logger function that logs the file name and line number
	var r *chi.Mux = chi.NewRouter() //struct that holds the routes

	handler.Handler(r) //call the handler function from the internal/handler we created

	fmt.Println("Starting Go Api Service")

	fmt.Print(`
  ██████╗  ██████╗      █████╗ ██████╗ ██╗
 ██╔════╝ ██╔═══██╗    ██╔══██╗██╔══██╗██║
 ██║  ███╗██║   ██║    ███████║██████╔╝██║
 ██║   ██║██║   ██║    ██╔══██║██╔═══╝ ██║
 ╚██████╔╝╚██████╔╝    ██║  ██║██║     ██║
  ╚═════╝  ╚═════╝     ╚═╝  ╚═╝╚═╝     ╚═╝`)

  err := http.ListenAndServe(":8080", r) //start the server on port 8080 and pass the router as the handler
  if err != nil {
	  log.Error(err) //log the error if the server fails to start		
  }
}