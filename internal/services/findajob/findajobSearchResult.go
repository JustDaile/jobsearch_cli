package findajob

import (
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/justdaile/jobsearch-cli/internal/services"
	"github.com/justdaile/jobsearch-cli/internal/utils"
)

type FindAJobResultImpl struct {
	JobSearchEngine JobSearchEngineImpl
	SearchTitle     string
	SearchLocation  string
	SearchRadius    int
	SearchFromage   int
	Page            int
	Document        *goquery.Document
}

func (sr FindAJobResultImpl) ResultsOnPage() int {
	return sr.Document.Find(".search-result").Length()
}

func (sr FindAJobResultImpl) GetResult(i int) services.JobPosting {
	result := sr.Document.Find(".search-result").Eq(i)
	jobTitleContainer := result.Find("h3 > a")
	details := result.Find(".search-result-details")

	isNew, dateErr := utils.CheckDateWithinDays("02 January 2006", strings.TrimSpace(details.Children().Eq(0).Text()), 1)
	if dateErr != nil {
		log.Printf("error while checking if date recent - %v", dateErr)
	}
	return FindAJobJobPostingImpl{
		Title:    utils.FixString(jobTitleContainer.Text()),
		Company:  utils.FixString(details.Children().Eq(1).Children().Eq(0).Text()),
		Link:     jobTitleContainer.AttrOr("href", "no-link"),
		New:      isNew,
		IsNative: false,
	}
}

func (sr FindAJobResultImpl) GetTotalPages() int {
	text := sr.Document.Find(".pager-items").Children().Last().Children().Eq(0).Text()
	i, e := strconv.ParseInt(text, 10, 64)
	if e != nil {
		log.Printf("unable to parse %s as integer - %v", text, e)
	}
	return int(i)
}

func (sr FindAJobResultImpl) GetCurrentPage() int {
	return sr.Page
}

func (sr FindAJobResultImpl) NextPage() (services.SearchResult, error) {
	return sr.JobSearchEngine.Search(sr.SearchTitle, sr.SearchLocation, sr.SearchRadius, sr.SearchFromage, sr.GetCurrentPage()+1)
}
