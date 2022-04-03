package fish4

import (
	"encoding/json"
)

type Fish4JobPostingImpl struct {
	Title    string `json:"title"`
	Company  string `json:"company"`
	Link     string `json:"link"`
	New      bool   `json:"is-new"`
	IsNative bool   `json:"is-native"`
}

func (i Fish4JobPostingImpl) GetTitle() string {
	return i.Title
}

func (i Fish4JobPostingImpl) GetCompany() string {
	return i.Company
}

func (i Fish4JobPostingImpl) GetLink() string {
	return i.Link
}

func (i Fish4JobPostingImpl) IsNew() bool {
	return i.New
}

func (i Fish4JobPostingImpl) NativeApply() bool {
	return i.IsNative
}

func (i Fish4JobPostingImpl) ToJSON() []byte {
	j, _ := json.Marshal(i)
	return j
}
