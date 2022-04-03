package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/justdaile/jobsearch-cli/internal"
	"github.com/justdaile/jobsearch-cli/internal/services"
	"github.com/justdaile/jobsearch-cli/internal/services/findajob"
	"github.com/justdaile/jobsearch-cli/internal/services/fish4"
	"github.com/justdaile/jobsearch-cli/internal/services/indeed"
)

var (
	withEngine  string
	withEngines string
	title       string
	titles      string
	location    string
	radius      int
	within      int
)

type nullWriter struct{}

func (nullWriter) Write(p []byte) (n int, err error) {
	return
}

func main() {
	showVersion := flag.Bool("version", false, "show the current version and build timestamp")
	newOnly := flag.Bool("new", false, "get only new job postings")
	csv := flag.Bool("csv", false, "output in comma seperated format")
	csvHeader := flag.Bool("csv-header", false, "output the csv header")
	debug := flag.Bool("debug", false, "show debug logging")

	flag.StringVar(&withEngine, "with-engine", "indeed", "The search engine to use [indeed, findajob, fish4]")
	flag.StringVar(&withEngines, "with-engines", "", "Comma seperated search engine to use [indeed, findajob, fish4]")
	flag.StringVar(&title, "title", "", "Job title")
	flag.StringVar(&titles, "titles", "", "comma seperated titles")
	flag.StringVar(&location, "location", "", "Job location")
	flag.IntVar(&radius, "radius", 5, "Radius from given location in miles. 0 is exact location")
	flag.IntVar(&within, "posted-within", 7, "The max age in days of the job posting")

	flag.Parse()

	if *showVersion {
		fmt.Printf("v%s - %s\n", internal.Version, internal.Timestamp)
		os.Exit(0)
	}

	if !*debug {
		log.SetOutput(nullWriter{})
	}

	searchTitles := []string{}
	searchEngines := []string{}
	jobPostings := []services.JobPosting{}

	if len(title) < 1 && len(titles) < 1 {
		exitIf(fmt.Errorf("requires flag -title or -titles"))
	}
	if len(titles) > 0 {
		searchTitles = append(searchTitles, strings.Split(titles, ",")...)
	} else {
		searchTitles = append(searchTitles, title)
	}
	if len(location) < 1 {
		exitIf(fmt.Errorf("requires flag -location"))
	}
	if len(withEngines) > 0 {
		searchEngines = append(searchEngines, strings.Split(withEngines, ",")...)
	} else {
		searchEngines = append(searchEngines, withEngine)
	}

	for _, searchEngine := range searchEngines {
		for _, searchTitle := range searchTitles {
			search, searchErr := createSearch(strings.TrimSpace(searchEngine), strings.TrimSpace(searchTitle), location, radius, within, 1)
			if logIf(searchErr) {
				continue
			}
			postings, postErr := doSearch(search, *newOnly)
			if logIf(postErr) {
				continue
			}

			jobPostings = append(jobPostings, postings...)
		}
	}

	log.Printf("%v total job postings", len(jobPostings))
	if *csv {
		if *csvHeader {
			fmt.Println("title,company,link,is-new,is-native")
		}
		for _, post := range jobPostings {
			fmt.Printf("\"%s\",\"%s\",\"%s\",%v,%v\n", post.GetTitle(), post.GetCompany(), post.GetLink(), post.IsNew(), post.NativeApply())
		}
	} else {
		jsonBytes, marshalErr := json.Marshal(jobPostings)
		exitIf(marshalErr)
		fmt.Print(string(jsonBytes))
	}

	os.Exit(0)
}

func createSearch(engine, title, location string, radius, within, fromage int) (services.SearchResult, error) {
	switch engine {
	case "indeed":
		return indeed.Engine.Search(title, location, radius, within, fromage)
	case "findajob":
		return findajob.Engine.Search(title, location, radius, within, fromage)
	case "fish4":
		return fish4.Engine.Search(title, location, radius, within, fromage)
	default:
		exitIf(fmt.Errorf("unknown engine %v", withEngine))
	}
	return nil, nil
}

func doSearch(searchResult services.SearchResult, newOnly bool) (p []services.JobPosting, searchErr error) {
	totalPages := searchResult.GetTotalPages()
	log.Printf("%v pages\n", totalPages)
	for i := 0; i < totalPages; i++ {
		totalJobs := searchResult.ResultsOnPage()
		log.Printf("%v jobs on page %v", totalJobs, (i + 1))
		for j := 0; j < totalJobs; j++ {
			posting := searchResult.GetResult(j)
			if newOnly && !posting.IsNew() {
				log.Printf("skipped job that was not new")
				continue
			}
			p = append(p, posting)
		}
		searchResult, searchErr = searchResult.NextPage()
	}
	return
}

func exitIf(e error) {
	if e != nil {
		fmt.Printf("error: %v\n", e)
		os.Exit(1)
	}
}

func logIf(e error) bool {
	if e != nil {
		log.Printf("error: %v", e)
		return true
	}
	return false
}
