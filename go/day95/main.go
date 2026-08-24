func getPort()(int, error) {
	return 0, errors.New("port not configured")
}

func startService()error {
	result, err := getPort()
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil
}
func main(){
	err := startService()
	if err != nil {
		fmt.Println(err)
	}
}

type Service struct {
	Name string
	Port int
	Status string
}

service := Service{
	Name: "Puneeth",
	Port: 8080,
	Status: "running"
}