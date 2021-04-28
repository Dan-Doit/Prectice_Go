package main

import (
	"os"
	"strings"

	"github.com/Dan-Doit/prectice-go/scrapper"
	"github.com/labstack/echo"
)

const FILENAME string = "jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrapper", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
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
