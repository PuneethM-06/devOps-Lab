data := "Cloud Platform kit"

err := os.WriteFile("output.txt", data, 0644)
if err != nil{
	fmt.Println("Failed to write file")
	return
}
fmt.Println("file written successfully ")