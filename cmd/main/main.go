package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"

    "github.com/rs/cors"
)

// LogMessage logs a given message to the terminal with the specified color
func LogMessage(message string, color string) {
    coloredMessage := fmt.Sprintf("\033[%sm%s\033[0m", color, message)
    log.Println(coloredMessage)
}

// LogHandler handles logging messages received via HTTP POST
func LogHandler(w http.ResponseWriter, r *http.Request) {
    var requestBody struct {
        Message string `json:"message"`
        Color   string `json:"color"`
    }

    if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    LogMessage(requestBody.Message, requestBody.Color)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintln(w, "Message logged")
}

func main() {
    fmt.Println("Starting server...")

    http.HandleFunc("/log", LogHandler)

    c := cors.New(cors.Options{
        AllowedOrigins:   []string{"*"},
        AllowedMethods:   []string{http.MethodGet, http.MethodPost, http.MethodOptions},
        AllowCredentials: true,
    })

    handler := c.Handler(http.DefaultServeMux)
    fmt.Println("Server started at :8080")
    log.Fatal(http.ListenAndServe(":8080", handler))
}