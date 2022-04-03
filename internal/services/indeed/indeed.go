package indeed

import (
	"fmt"
	"log"
	"math"
	"net/url"

	"github.com/justdaile/jobsearch-cli/internal/services"
	"github.com/justdaile/jobsearch-cli/internal/utils"
)

type JobSearchEngineImpl struct {
	Alias     string
	Host      string
	SearchURI string
}

// Example: https://uk.indeed.com/jobs?q=Software+Developer&l=Birmingham&radius=5&fromage=1&start=10
var Engine = JobSearchEngineImpl{
	Alias:     "Indeed",
	Host:      "https://uk.indeed.com",
	SearchURI: "/jobs?q=%s&l=%s&radius=%v&fromage=%v&start=%v",
}

func (js JobSearchEngineImpl) Search(title string, location string, radius int, fromage int, page int) (services.SearchResult, error) {
	title = url.QueryEscape(title)
	location = url.QueryEscape(location)

	url := fmt.Sprintf("%s%s", js.Host, fmt.Sprintf(js.SearchURI, title, location, radius, fromage, 10*int(math.Abs(float64(page-1)))))

	doc, err := utils.QuickDoc(url)
	if err != nil {
		log.Fatal(err)
	}

	return IndeedSearchResultImpl{
		JobSearchEngine: js,
		SearchTitle:     title,
		SearchLocation:  location,
		SearchRadius:    radius,
		SearchFromage:   fromage,
		Page:            page,
		Document:        doc,
	}, nil
}
