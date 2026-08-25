
APP_ENV=production
environment := os.Getenv("APP_ENV")

if environment == "" {
	fmt.Println("APP_ENV does not exist")
}