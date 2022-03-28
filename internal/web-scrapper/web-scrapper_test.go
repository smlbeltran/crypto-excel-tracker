package webscrapper_test

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	webscrapper "smlbeltran/crypto-excel-tracker/internal/web-scrapper"
	"testing"
)

func TestScrapCryptoCurrencyPrice(t *testing.T) {
	t.Run("get crypto prices", func(t *testing.T) {
		srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			html := `
			<html>
				<body>
					<div class="priceValue">
						<span>$1</span>
					</div>	
				</body>
			</html>
			`
			w.Write([]byte(html))
		}))

		client := webscrapper.Client{URL: srv.URL}

		got := client.GetPrice([]string{"bitcoin", "eos", "xrp"})

		want := map[string]string{
			"bitcoin": "1",
			"eos":     "1",
			"xrp":     "1",
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want:%q, got: %q", want, got)
		}
	})
}
