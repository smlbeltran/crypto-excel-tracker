package webscrapper

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

type Client struct {
	URL string
}

func (c Client) GetPrice(coins []string) map[string]string {
	cryptoList := map[string]string{}

	ch := make(chan map[string]string)

	for i, coin := range coins {

		go func(coin string, i int) {
			resp, _ := http.Get(fmt.Sprintf("%s/%s", c.URL, coin))

			data, _ := io.ReadAll(resp.Body)

			docHTML, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
			if err != nil {
				panic(err)
			}

			price := strings.TrimSpace(docHTML.Find(".priceValue").Text())

			cryptoList[coin] = strings.ReplaceAll(price, "$", "")

			ch <- cryptoList

			if i == len(coins)-1 {
				time.Sleep(500 * time.Millisecond)
				close(ch)
			}

		}(coin, i)
	}

	for {
		_, ok := <-ch

		if !ok {
			break
		}

	}

	return cryptoList
}
