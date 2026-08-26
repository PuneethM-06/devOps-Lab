package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request){
	if r.Method != "POST"{
		http.Error(w, "Method Not allowed", http.StatusMethodNotAllowed)
		return
	}
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	var scanRequest ScanRequest
	var scanResponse ScanResponse

	err = json.Unmarshal(data, &scanRequest)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	fmt.Fprintln(w, scanRequest.Repository)
	scanResponse.Repository = scanRequest.Repository
	scanResponse.Status = "Scan started"

	scan, err := json.Marshall(ScanResponse)
	if err != nil {
		http.Error(w, "Failed", http.StatusInternalServerError)
		return 
	}

}
type ScanRequest struct {
	Repository string `json:"repository"`
}

type ScanResponse struct {
	Repository string `json:"repository"`
	Status string `json:"status"`
}

func main(){
	http.HandleFunc("/scan", handler)
	http.ListenAndServe(":8080", nil)
	
}
