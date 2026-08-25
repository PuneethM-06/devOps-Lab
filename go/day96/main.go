data, err := os.ReadFile("message.txt")

if err != nil {
	fmt.Println(err)
	return 
}
fmt.Println(string(data)