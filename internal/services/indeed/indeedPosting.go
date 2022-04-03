package indeed

import (
	"encoding/json"

	"github.com/justdaile/jobsearch-cli/internal/services"
)

type IndeedJobPostingImpl struct {
	Title    string              `json:"title"`
	Company  string              `json:"company"`
	Link     string              `json:"link"`
	New      bool                `json:"is-new"`
	IsNative bool                `json:"is-native"`
	Detail   *services.JobDetail `json:"detail,omitempty"`
}

func (i IndeedJobPostingImpl) GetTitle() string {
	return i.Title
}

func (i IndeedJobPostingImpl) GetCompany() string {
	return i.Company
}

func (i IndeedJobPostingImpl) GetLink() string {
	return i.Link
}

func (i IndeedJobPostingImpl) IsNew() bool {
	return i.New
}

func (i IndeedJobPostingImpl) NativeApply() bool {
	return i.IsNative
}

func (i IndeedJobPostingImpl) ToJSON() []byte {
	j, _ := json.Marshal(i)
	return j
}
