package main

import (
	"bufio"
	"io/fs"
	"os"
	"smlbeltran/crypto-excel-tracker/boot"
	"smlbeltran/crypto-excel-tracker/internal/spreadsheet"
	webscrapper "smlbeltran/crypto-excel-tracker/internal/web-scrapper"
)

func main() {
	cfg := boot.NewConfig()

	wsc := webscrapper.Client{
		URL: cfg.URL,
	}

	coinList := readCoinsFromFile(os.DirFS(cfg.DirPath))

	coins := wsc.GetPrice(coinList)

	spreadsheet.Update(coins, cfg)
}

func readCoinsFromFile(fileSystem fs.FS) []string {
	var coins []string

	file, _ := fs.ReadDir(fileSystem, ".")

	fop, err := fileSystem.Open(file[0].Name())
	if err != nil {
		panic(err)
	}

	defer fop.Close()

	scanner := bufio.NewScanner(fop)

	for scanner.Scan() {
		if len(scanner.Text()) == 0 {
			continue
		}
		coins = append(coins, scanner.Text())
	}

	return coins
}
