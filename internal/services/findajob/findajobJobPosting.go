package findajob

import (
	"encoding/json"
)

type FindAJobJobPostingImpl struct {
	Title    string `json:"title"`
	Company  string `json:"company"`
	Link     string `json:"link"`
	New      bool   `json:"is-new"`
	IsNative bool   `json:"is-native"`
}

func (i FindAJobJobPostingImpl) GetTitle() string {
	return i.Title
}

func (i FindAJobJobPostingImpl) GetCompany() string {
	return i.Company
}

func (i FindAJobJobPostingImpl) GetLink() string {
	return i.Link
}

func (i FindAJobJobPostingImpl) IsNew() bool {
	return i.New
}

func (i FindAJobJobPostingImpl) NativeApply() bool {
	return i.IsNative
}

func (i FindAJobJobPostingImpl) ToJSON() []byte {
	j, _ := json.Marshal(i)
	return j
}
