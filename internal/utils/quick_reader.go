package utils

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func QuickDoc(url string) (*goquery.Document, error) {
	log.Printf("reading job list: [GET] %s\n", url)

	req, reqErr := http.NewRequest("GET", url, nil)
	if reqErr != nil {
		return nil, reqErr
	}
	res, resErr := http.DefaultClient.Do(req)
	if resErr != nil {
		return nil, resErr
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("got status %v - %v", res.StatusCode, res.Status)
	}

	log.Printf("loading goquery document")
	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func QuickRead(url string, a func(*goquery.Document)) error {
	d, e := QuickDoc(url)
	if e != nil {
		return e
	}
	a(d)
	return nil
}
