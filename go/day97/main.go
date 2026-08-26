
func handler(w http.ResponseWriter, r *http.Request){
	if r.Method != "GET"{
		fmt.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.FPrintln(w, "Hello from Server")
}

http.HandleFunc("/", handler)
http.HandleFunc("/health", healthHandler)
http.ListenAndServe(":8080", nil)

func healthHandler(w http.ResponseWriter, r *http.Request){
	fmt.FPrintln(w, "OK")
	fmt.Println(r.Method)
}

if err != nil {
	fmt.Println(err)
	return
}

res, err := io.ReadAll(resp.Body)
fmt.Println(string(res))