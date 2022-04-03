package fish4

import (
	"fmt"
	"log"
	"net/url"

	"github.com/justdaile/jobsearch-cli/internal/services"
	"github.com/justdaile/jobsearch-cli/internal/utils"
)

type JobSearchEngineImpl struct {
	Alias     string
	Host      string
	SearchURI string
}

var Engine = JobSearchEngineImpl{
	Alias:     "Fish4",
	Host:      "https://www.fish4.co.uk",
	SearchURI: "/searchjobs/?Keywords=%s&radialtown=%s&RadialLocation=%v&Page=%v",
}

func (js JobSearchEngineImpl) Search(title string, location string, radius int, fromage int, page int) (services.SearchResult, error) {
	title = url.QueryEscape(title)
	location = url.QueryEscape(location)

	url := fmt.Sprintf("%s%s", js.Host, fmt.Sprintf(js.SearchURI, title, location, radius, page))

	doc, err := utils.QuickDoc(url)
	if err != nil {
		log.Fatal(err)
	}

	return Fish4ResultImpl{
		JobSearchEngine: js,
		SearchTitle:     title,
		SearchLocation:  location,
		SearchRadius:    radius,
		SearchFromage:   fromage,
		Page:            page,
		Document:        doc,
	}, nil
}
