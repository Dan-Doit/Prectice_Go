package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Dan-Doit/prectice-go/scrapper"
	"github.com/labstack/echo"
)

const FILENAME string = "jobs.csv"

type user struct {
	name string
	age  int
}

var User user

func main() {
	User = user{}
	User.age = 29
	User.name = "dan"
	fmt.Println(User)
	// e := echo.New()
	// e.GET("/", handleHome)
	// e.POST("/scrapper", handleScrape)
	// e.Logger.Fatal(e.Start(":1323"))
}

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer os.Remove(FILENAME)
	turm := strings.ToLower(scrapper.CleanString(c.FormValue("turm")))
	scrapper.Scrapper(turm)
	return c.Attachment(FILENAME, FILENAME)
}
