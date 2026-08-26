package main 
import (
	"fmt"
	"net/http"
)
func handler(w http.ResponseWriter, r *http.Request){
	fmt.FPrintln(w, "Server")
}

type ScanRequest struct {
	Repository string `json:"repository"`
}

func main(){
	http.HandleFunc("/scan", handler)
	http.ListenAndServe(":8080", nil)
}
