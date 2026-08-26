package main 
import (
	"fmt"
	"net/http"
)
http.HandleFunc("/scan", handler)
http.ListenAndServe(":8080", nil )

func main() {
	func handler(w http.ResponseWriter, r *httpRequest) {
		fmt.FPrintln("Server")
	}
}