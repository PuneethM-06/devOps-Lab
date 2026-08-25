type Service struct {
	Name string 
	Port int 
	Status string 
}

service := Service {
	Name: "api",
	Port: 8080,
	Status: "running"
}
data, err := json.Marshal(service)
if err != nil {
	fmt.Println(err)
	return
}
fmt.Println(string(data))