package indeed

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/justdaile/jobsearch-cli/internal/services"
	"github.com/justdaile/jobsearch-cli/internal/utils"
)

type IndeedSearchResultImpl struct {
	JobSearchEngine JobSearchEngineImpl
	SearchTitle     string
	SearchLocation  string
	SearchRadius    int
	SearchFromage   int
	Page            int
	Document        *goquery.Document
}

func (sr IndeedSearchResultImpl) ResultsOnPage() int {
	return sr.Document.Find(".jobTitle").Length()
}

func (sr IndeedSearchResultImpl) GetResult(i int) services.JobPosting {
	result := sr.Document.Find("a.result").Eq(i)
	jobTitleContainer := result.Find(".jobTitle")
	jobTitleChildren := jobTitleContainer.Children().Length()
	return IndeedJobPostingImpl{
		Title:    utils.FixString(jobTitleContainer.Children().Eq(jobTitleChildren-1).AttrOr("title", "unknown")),
		Company:  utils.FixString(result.Find(".companyName").Text()),
		Link:     fmt.Sprintf("%s/viewjob?jk=%s", sr.JobSearchEngine.Host, result.AttrOr("data-jk", "no-link")),
		New:      jobTitleContainer.Children().Eq(0).Is("div.new"),
		IsNative: result.Find(".ialbl").Length() > 0,
	}
}

func (sr IndeedSearchResultImpl) GetTotalPages() int {
	return sr.Document.Find(".pagination-list").Children().Length() - 1
}

func (sr IndeedSearchResultImpl) GetCurrentPage() int {
	return sr.Page
}

func (sr IndeedSearchResultImpl) NextPage() (services.SearchResult, error) {
	return sr.JobSearchEngine.Search(sr.SearchTitle, sr.SearchLocation, sr.SearchRadius, sr.SearchFromage, sr.GetCurrentPage()+1)
}
