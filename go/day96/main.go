data, err := os.ReadFile("message.txt")

if err != nil {
	fmt.Println("Failed to read message.txt")
	return 
}
fmt.Println(string(data)