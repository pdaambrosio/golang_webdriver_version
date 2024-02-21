// webhit.go
package webhit

import (
	"fmt"
	"log"
	"time"

	"github.com/sclevine/agouti"
)

func PunchTheClock(url, user, password string) string {
	pageResult := ""

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{"--no-sandbox", "--disable-dev-shm-usage"}),
	)

	err := driver.Start()
	if err != nil {
		log.Fatal("Failed to start driver: %s", err)
	}

	page, err := driver.NewPage()
	if err != nil {
		panic(err)
	}

	err = page.Navigate(url)
	if err != nil {
		panic(err)
	}

	time.Sleep(3 * time.Second)

	err = page.FindByID("button-1020-btnInnerEl").Click()
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	err = page.FindByID("ext-136").Fill(user)
	if err != nil {
		panic(err)
	}

	err = page.FindByID("ext-138").SendKeys(password)
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	err = page.FindByID("ext-140").Click()
	if err != nil {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	resultText, err := page.FindByID("ext-142").Text()
	if err != nil {
		panic(err)
	}

	pageResult = resultText

	date := time.Now().Format("02-01-2006 15:04:05")

	time.Sleep(2 * time.Second)

	err = driver.Stop()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s - %s", pageResult, date)
}
