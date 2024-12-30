package htmlparser

import (
	"fmt"
	"net/http"
	"time"

	"github.com/PuerkitoBio/goquery"
)

const (
	responses_url = "https://www.banki.ru/services/responses/bank/response/"
	header_tag    = "body > div.page-container > main > div:nth-child(2) > section:nth-child(1) > header > div > div > div:nth-child(2) > div > h1"
)

func Parse(id int, amount int) ([]string, error) {
	res := make([]string, 0, amount)
	for i := 0; i < amount; i++ {
		time.Sleep(time.Second) // боремся с rate limiter
		resp, err := http.Get(fmt.Sprintf("%s%d/", responses_url, id+i))
		if err != nil {
			continue
		}
		defer resp.Body.Close()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			continue
		}
		text := doc.Find(header_tag).Text()

		res = append(res, text)
	}
	return res, nil
}
