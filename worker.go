package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
)

func handleReplication(w http.ResponseWriter, r *http.Request) {
	// Simulate handling the replicated data
	data, _ := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	// Store or process the data (for simplicity, just printing it here)
	fmt.Printf("Node 2 (Replica) received data: %s\n", string(data))

	// Respond back to Node 1
	fmt.Fprintf(w, "Data successfully replicated to Node 2")
}

func main() {
	http.HandleFunc("/replicate", handleReplication)
	log.Println("Node 2 (Replica) is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))
}
