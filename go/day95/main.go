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
	err = startService()
	if err != nil {
		fmt.Println(err)
	}
}