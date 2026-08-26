package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request){
	fmt.FPrintln(w, "Server")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	var scanRequest ScanRequest
	err = json.Unmarshal(data, &scanRequest)
	if err != nil {
		fmt.Println(err)
	}
}

type ScanRequest struct {
	Repository string `json:"repository"`
}

func main(){
	http.HandleFunc("/scan", handler)
	http.ListenAndServe(":8080", nil)
}
