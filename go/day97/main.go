
func handler(w http.ResponseWriter, r *http.Request){
	fmt.FPrintln(w, "Hello from Server")
}

http.handleFunc("/", handler)
http.ListenAndServe(":8080", nil)