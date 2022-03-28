package boot

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	SheetID string `json:"sheet_id"`
	URL     string
	Range   map[string][]string
	DirPath string `json:"crypto_list_path"`
}

func NewConfig() *Config {
	var config Config

	file, err := os.Open(filepath.Clean("conf.json"))
	if err != nil {
		panic(err)
	}

	err = json.NewDecoder(file).Decode(&config)
	if err != nil {
		panic(err)
	}

	return &config

}
