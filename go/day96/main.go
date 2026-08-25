type ScanRequest struct {
	Repository string 
	Branch string
}
var scanRequest ScanRequest 
res, err := json.Unmarshal(data, &scanRequest)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(ScanRequest.Repository, scanRequest.Branch)