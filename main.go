package main

import (
	"os"
	"webhit/webhit"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		println("Please set the environment variables WEB_USER and WEB_PASSWORD in .env file.")
	}

	url := "https://cliente.apdata.com.br/everisparceiro/"
	user := os.Getenv("WEB_USER")
	password := os.Getenv("WEB_PASSWORD")

	accessLog, err := os.OpenFile("access.logs", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer accessLog.Close()

	result := webhit.PunchTheClock(url, user, password)
	_, err = accessLog.WriteString(result + "\n")
	if err != nil {
		panic(err)
	}
}
