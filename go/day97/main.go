package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request){
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var scanRequest ScanRequest
	err = json.Unmarshal(data, &scanRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, scanRequest.Repository)
}
type ScanRequest struct {
	Repository string `json:"repository"`
}

func main(){
	http.HandleFunc("/scan", handler)
	http.ListenAndServe(":8080", nil)
	
}
