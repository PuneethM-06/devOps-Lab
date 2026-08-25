# DAY 96 - GOLANG - FILE HANDLING, JSON, ENVIRONMENT VARIABLES PACKAGES

1. ### File handling 
- **File contains bytes, so Go reads the file as bytes**
- Example: `data, err := os.ReadFile("config.txt")`

2. ### File handling - Writing 
- This operation allows you to write the file 
- Example: `os.WriteFile("output.txt", data, 0644)`
```
"output.txt" → filename
data         → content to write
0644         → file permissions
```
- Example:
```
data := "Cloud Platform kit"

err := os.WriteFile("output.txt", data, 0644)
if err != nil{
	fmt.Println("Failed to write file")
	return
}
fmt.Println("file written successfully ")
```
