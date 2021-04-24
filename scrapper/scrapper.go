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

var baseURL string = "https://kr.indeed.com/jobs?q=react&l=%EC%84%9C%EC%9A%B8"

type jobsInfo struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

func Scrapper() {

	arrJobsInfo := []jobsInfo{}

	totalPages := getPages(baseURL)

	for i := 0; i < totalPages; i++ {
		arrJobsInfo = append(arrJobsInfo, getPage(i)...)
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

func getPage(page int) []jobsInfo {

	arrInfo := []jobsInfo{}

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
		result := exportJobs(card)
		arrInfo = append(arrInfo, result)
	})

	return arrInfo
}

func exportJobs(card *goquery.Selection) jobsInfo {
	id, _ := card.Attr("data-jk")
	title := cleanString(card.Find(".title>a").Text())
	location := cleanString(card.Find(".sjcl").Text())
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())

	return jobsInfo{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func cleanString(str string) string {
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
