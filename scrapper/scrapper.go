package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type jobsInfo struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

func Scrapper(turm string) {
	var baseURL string = "https://kr.indeed.com/jobs?q=" + turm + "&l=%EC%84%9C%EC%9A%B8"

	arrJobsInfo := []jobsInfo{}

	ch := make(chan []jobsInfo)

	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		go getPage(i, baseURL, ch)
	}

	for i := 0; i < totalPages; i++ {
		arrJobsInfo = append(arrJobsInfo, <-ch...)
	}

	writeCSV(arrJobsInfo)

	fmt.Println("Done! check your cvs file!")
}

func writeCSV(jobs []jobsInfo) {
	file, err := os.Create("jobs.csv")
	errChecker(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	header := []string{"Id", "Title", "Location", "Salary", "Summary"}
	wErr := w.Write(header)
	errChecker(wErr)

	for _, v := range jobs {
		job := []string{v.id, v.title, v.location, v.salary, v.summary}
		jErr := w.Write(job)
		errChecker(jErr)
	}
}

func getPage(page int, baseURL string, ch chan<- []jobsInfo) {

	arrInfo := []jobsInfo{}
	c := make(chan jobsInfo)
	pageURL := baseURL + "&start=" + strconv.Itoa(page*10)
	fmt.Println("Requesting now : ", pageURL)

	res, err := http.Get(pageURL)
	// check
	resChecker(res)
	errChecker(err)
	// clean up
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	errChecker(err)

	doc.Find(".jobsearch-SerpJobCard").Each(func(i int, card *goquery.Selection) {
		go exportJobs(card, c)
	})

	for i := 0; i < doc.Find(".jobsearch-SerpJobCard").Length(); i++ {
		arrInfo = append(arrInfo, <-c)
	}

	ch <- arrInfo
}

func exportJobs(card *goquery.Selection, c chan<- jobsInfo) {
	id, _ := card.Attr("data-jk")
	title := CleanString(card.Find(".title>a").Text())
	location := CleanString(card.Find(".sjcl").Text())
	salary := CleanString(card.Find(".salaryText").Text())
	summary := CleanString(card.Find(".summary").Text())

	c <- jobsInfo{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func CleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}

func getPages(url string) int {
	pages := 0
	res, err := http.Get(url)
	// check
	resChecker(res)
	errChecker(err)
	// clean up
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	errChecker(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func resChecker(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Can't Join in with : ", res)
	}
}

func errChecker(err error) {
	if err != nil {
		log.Fatalln("error : ", err)
	}
}
