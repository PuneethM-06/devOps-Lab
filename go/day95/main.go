func main() {
	status := "running"

	switch status {
	case "running":
		fmt.Println("Service is running")
	case "stopped":
		fmt.Println("Service is stopped")
	case "failed":
		fmt.Println("service has failed")	
	default:
		fmt.Println("unknown status")
	}	
}