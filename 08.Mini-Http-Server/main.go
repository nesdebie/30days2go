package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "time"

    // Modern HTTP router for Go
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

// Handler to return a greeting message
func helloHandler(w http.ResponseWriter, r *http.Request) {
    name := chi.URLParam(r, "name")
    response := fmt.Sprintf("Hello, %s!", name)
    w.Write([]byte(response))
}

// Handler to return the current server time in JSON format
func timeHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "server_time": time.Now().Format(time.RFC3339),
    })
}

func main() {
    router := chi.NewRouter()

    // Create a request ID header
    router.Use(middleware.RequestID)

    // Log the start and end of each request
    router.Use(middleware.Logger)

    // Recover from panics and log errors
    router.Use(middleware.Recoverer)

    // Available routes
    router.Get("/hello/{name}", helloHandler)
    router.Get("/time", timeHandler)

    const port = ":8090"
    fmt.Println("Starting server at http://localhost" + port)
    fmt.Println("Visit: curl http://localhost" + port + "/hello/<your_name>")
    fmt.Println("Visit: curl http://localhost" + port + "/time")

    // Run the server
    if err := http.ListenAndServe(port, router); err != nil {
        log.Fatal(err)
    }
}
