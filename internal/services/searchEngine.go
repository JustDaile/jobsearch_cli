package services

type JobSearchEngine interface {
	Search(string, string, ...int) (SearchResult, error)
}

type SearchResult interface {
	ResultsOnPage() int
	GetResult(int) JobPosting
	GetTotalPages() int
	GetCurrentPage() int
	NextPage() (SearchResult, error)
}

type JobDetail struct {
	// Full job description
	Description string
	// Salary
	Salary string
	// Working hours
	Hours string
	// Perm Part-time Full-time, etc
	Type string
	// Posted date
	DatePosted string
	// Ending date
	DateEnding string
}

type JobPosting interface {
	// Get the job title
	GetTitle() string
	// Get the company or agency name
	GetCompany() string
	// Get the job apply link
	GetLink() string
	// Is a new job posting or is recent
	IsNew() bool
	// Can apply to this posting on site
	NativeApply() bool
	// Marshal to JSON
	ToJSON() []byte
}
