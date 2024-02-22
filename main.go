package main

import (
	"os"
	"webhit/webhit"

	"github.com/joho/godotenv"
)

// The main function loads environment variables, makes a web request using credentials, and logs the
// result to a file.
func main() {
	err := godotenv.Load()
	if err != nil {
		println("Please set the environment variables WEB_USER and WEB_PASSWORD in .env file.")
		panic(err)
	}

	url := "https://url.com"
	user := os.Getenv("WEB_USER")
	password := os.Getenv("WEB_PASSWORD")

	accessLog, err := os.OpenFile("access.logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		println("Error opening the access.logs file")
		panic(err)
	}
	defer accessLog.Close()

	result := webhit.PunchTheClock(url, user, password)
	_, err = accessLog.WriteString(result + "\n")
	if err != nil {
		println("Error writing to the access.logs file")
		panic(err)
	}
}
