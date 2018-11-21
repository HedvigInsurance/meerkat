package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func FetchXmlFromUrl(url string) ([]byte, error) {

	start_xml := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Read body: %v", err)
	}

	log.Println("XML took ", time.Since(start_xml))

	return data, nil
}
