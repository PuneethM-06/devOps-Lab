
func handler(w http.ResponseWriter, r *http.Request){
	fmt.FPrintln(w, "Hello from Server")
}

http.HandleFunc("/", handler)
http.HandleFunc("/health", healthHandler)
http.ListenAndServe(":8080", nil)

func healthHandler(w http.ResponseWriter, r *http.Request){
	fmt.FPrintln(w, "OK")
	fmt.Println(r.Method)
}
