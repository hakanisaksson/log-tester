package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "time"
    "strconv"
)

var total_messages_sent int = 0

// Simple helper function to read an environment or return a default value
func getEnv(key string, defaultVal string) string {
    if value, exists := os.LookupEnv(key); exists {
            return value
    }

    return defaultVal
}

// Simple helper function to read an environment variable into integer or return a default value
func getEnvAsInt(name string, defaultVal int) int {
    valueStr := getEnv(name, "")
    if value, err := strconv.Atoi(valueStr); err == nil {
            return value
    }

    return defaultVal
}

func healthGET(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "{\"status\": \"UP\"}\n")
}

func readyGET(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "{\"status\": \"OK\"}\n")
}

func metricsGET(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "# HELP log_tester_messages_sent_counter_total\n")
    fmt.Fprintf(w, "# TYPE log_tester_messages_sent_counter_total counter\n")
    fmt.Fprintf(w, "log_tester_messages_sent_counter_total %d\n", total_messages_sent)
}

func periodic_update(update_interval int) {
    if update_interval != 0 {
            total_messages_sent = total_messages_sent + 1
            log.Printf("Hello #%d", total_messages_sent )
    }
}


func hello(w http.ResponseWriter, r *http.Request) {
    log.Printf("Serving request: %s", r.URL.Path)
    host, _ := os.Hostname()
    fmt.Fprintf(w, "Hello, world!\n")
    fmt.Fprintf(w, "Version: 1.0.0\n")
    fmt.Fprintf(w, "Hostname: %s\n", host)
}


func startPolling1( update_interval int ) {
  for {
    time.Sleep(time.Duration(update_interval) * time.Second)
    go periodic_update(update_interval)
  }
}

func main() {
    // use PORT environment variable, or default to 8080
    port := "8080"
    if fromEnv := os.Getenv("PORT"); fromEnv != "" {
            port = fromEnv
    }

    update_interval := getEnvAsInt("INTERVAL", 3)

    // register hello function to handle all requests
    server := http.NewServeMux()
    server.HandleFunc("/", hello)
    server.HandleFunc("/health", healthGET)
    server.HandleFunc("/ready", readyGET)
    server.HandleFunc("/metrics", metricsGET)
    go startPolling1(update_interval)

    // start the web server on port and accept requests
    log.Printf("Server listening on port %s", port)
    log.Printf("INTERVAL: %d\n", update_interval)
    err := http.ListenAndServe(":"+port, server)
    log.Fatal(err)
}
