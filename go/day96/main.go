data := []byte(`{
  "repository": "github.com/mycompany/payment-service",
  "branch": "main"
}`)
type ScanRequest struct {
	Repository string 
	Branch string
}
var scanRequest ScanRequest 
err := json.Unmarshal(data, &scanRequest)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(scanRequest.Repository, scanRequest.Branch)