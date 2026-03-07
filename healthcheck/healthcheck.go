package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = "8080"
	}

	client := http.Client{
		Timeout: 2 * time.Second,
	}

	resp, err := client.Get("http://localhost:" + port + "/apis/health")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Health Check Response Status:", resp.Status)

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Health check failed with status: %s", resp.Status)
	} else {
		log.Println("Health check successful!")
		os.Exit(1)
	}
}
