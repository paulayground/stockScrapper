package request

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// get Request
func GetStockList(url string) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	// parsing
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	doc.Find("tr>td:nth-child(2)").Each(func(i int, s *goquery.Selection) {
		stockList := strings.TrimSpace(s.Text())
		fmt.Println(stockList)
	})
}
