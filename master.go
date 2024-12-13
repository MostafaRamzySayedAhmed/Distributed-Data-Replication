package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func replicateData(w http.ResponseWriter, r *http.Request) {
	// Simulate the process of replicating data
	data := "Sample Data from Node 1"

	// Replicate the data to Node 2
	resp, err := http.Post("http://localhost:8081/replicate", "text/plain", ioutil.NopCloser(strings.NewReader(data)))
	if err != nil {
		log.Fatalf("Failed to replicate data: %v", err)
	}
	defer resp.Body.Close()

	// Send the data back to the requester
	fmt.Fprintf(w, "Data replicated to Node 2: %s", data)
}

func main() {
	http.HandleFunc("/replicate", replicateData)
	log.Println("Node 1 (Leader) is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
