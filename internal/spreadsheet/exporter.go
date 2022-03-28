package spreadsheet

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"smlbeltran/crypto-excel-tracker/boot"

	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
)

func Update(coins map[string]string, c *boot.Config) {
	srv, err := sheets.NewService(context.Background(),
		option.WithCredentialsFile(filepath.Clean("service_account.json")))

	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	var readRange []string

	for coin, price := range coins {
		switch coin {
		case "bitcoin-cash":
			readRange = c.Range[coin]
		case "eos":
			readRange = c.Range[coin]
		case "xrp":
			readRange = c.Range[coin]
		case "stellar":
			readRange = c.Range[coin]
		}

		rb := &sheets.ValueRange{
			Values: [][]interface{}{
				{price},
			},
		}

		for _, cell := range readRange {
			fmt.Println("Updating Cell:", cell)

			resp, err := srv.Spreadsheets.Values.Update(c.SheetID, cell, rb).
				ValueInputOption("USER_ENTERED").
				Context(context.Background()).
				Do()

			if err != nil {
				log.Fatalf("Unable to update sheet: %v", err)
			}
			fmt.Fprintf(os.Stdout, "coin %s response status: %d \n ---- \n", coin, resp.HTTPStatusCode)
		}

	}

	fmt.Fprintf(os.Stdout, "view url: https://docs.google.com/spreadsheets/d/%s\n", c.SheetID)
}
