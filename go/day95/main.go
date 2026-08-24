func main() {
	serviceName := "cloud-platform-kit"
	port := 8080
	production := false
	maxRequests := 3
	fmt.Println("Starting service:", serviceName, "port is:", port, "production:", production, "maxRequests:", maxRequests)

	environemnt := "production"
	replicas := 3
	cpuUsage := 72.5
	healthy := true

	service := []string{"api", "worker", "frontend"}
	service = append(service, "database")
	fmt.Println(service[0]) 

	servicePorts := map[string]int{
		"api": 8080,
		"worker": 8081,
		"frontend": 8082
	}
	servicePorts[worker] = 9090
	serviceNew := map[string]int{
		"api": 8080,
		"worker": 8081,
		"frontend": 8082,
		"database": 5432
	}
	fmt.Println(service)
	
}
