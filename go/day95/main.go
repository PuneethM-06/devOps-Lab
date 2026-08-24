func getPort()(int, error) {
	return 8080, nil
}

func startService()error {
	result, err := getPort()
	if err != nil {
		return err
	}
	fmt.Println(result)
	return nil

