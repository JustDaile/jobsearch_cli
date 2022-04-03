package findajob

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
	Alias:     "FindAJob",
	Host:      "https://findajob.dwp.gov.uk",
	SearchURI: "/search?q=%s&w=%s&p=%v",
}

func (js JobSearchEngineImpl) Search(title string, location string, radius int, fromage int, page int) (services.SearchResult, error) {
	title = url.QueryEscape(title)
	location = url.QueryEscape(location)

	url := fmt.Sprintf("%s%s", js.Host, fmt.Sprintf(js.SearchURI, title, location, page))

	doc, err := utils.QuickDoc(url)
	if err != nil {
		log.Fatal(err)
	}

	return FindAJobResultImpl{
		JobSearchEngine: js,
		SearchTitle:     title,
		SearchLocation:  location,
		SearchRadius:    radius,
		SearchFromage:   fromage,
		Page:            page,
		Document:        doc,
	}, nil
}
