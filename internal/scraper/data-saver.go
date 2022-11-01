package scraper

import (
	"encoding/json"
	"fmt"
	"os"
)

// "./outputs/testing.json"

func SaveScrapedData(name string, data interface{}) error {
	byteData, err := json.MarshalIndent(data, "", "	")
	if err != nil {
		return fmt.Errorf("ERROR\t%v\n", err)
	}

	err = os.WriteFile(fmt.Sprintf("%s/%s", "static/scraped-data/", name), byteData, 0644)
	if err != nil {
		return fmt.Errorf("ERROR\t%v\n", err)
	}

	return nil
}
