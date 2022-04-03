package fish4

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/justdaile/jobsearch-cli/internal/services"
	"github.com/justdaile/jobsearch-cli/internal/utils"
)

type Fish4ResultImpl struct {
	JobSearchEngine JobSearchEngineImpl
	SearchTitle     string
	SearchLocation  string
	SearchRadius    int
	SearchFromage   int
	Page            int
	Document        *goquery.Document
}

func (sr Fish4ResultImpl) ResultsOnPage() int {
	return sr.Document.Find(".lister__details").Length()
}

func (sr Fish4ResultImpl) GetResult(i int) services.JobPosting {
	result := sr.Document.Find(".lister__details").Eq(i)
	posted := utils.FixString(sr.Document.Find(".lister__footer").Eq(i).Find("ul.job-actions").Children().Eq(1).Text())
	jobTitleContainer := result.Find(".lister__header")
	isNew := strings.Contains(posted, "ago") && strings.Contains(posted, "1")
	link := fmt.Sprintf("%s%s", sr.JobSearchEngine.Host, utils.FixLink(jobTitleContainer.Find("a").AttrOr("href", "no-link")))

	isNative := false
	utils.QuickRead(link, func(doc *goquery.Document) {
		applyBtn := doc.Find(".button--apply").Eq(0)
		if applyBtn.AttrOr("href", "none") == "#apply-form" {
			isNative = true
		}
	})

	return Fish4JobPostingImpl{
		Title:    utils.FixString(jobTitleContainer.Find("span").Text()),
		Company:  utils.FixString(result.Find(".lister__meta").Children().Eq(2).Text()),
		Link:     link,
		New:      isNew,
		IsNative: isNative,
	}
}

func (sr Fish4ResultImpl) GetTotalPages() int {
	paginatorItems := sr.Document.Find(".paginator__items").Children().Length()
	if paginatorItems < 1 {
		return 1
	}
	text := sr.Document.Find(".paginator__items").Children().Eq(paginatorItems - 2).Find("a").Text()
	i, e := strconv.ParseInt(text, 10, 64)
	if e != nil {
		log.Printf("unable to parse %s as integer - %v", text, e)
	}
	return int(i)
}

func (sr Fish4ResultImpl) GetCurrentPage() int {
	return sr.Page
}

func (sr Fish4ResultImpl) NextPage() (services.SearchResult, error) {
	return sr.JobSearchEngine.Search(sr.SearchTitle, sr.SearchLocation, sr.SearchRadius, sr.SearchFromage, sr.GetCurrentPage()+1)
}
