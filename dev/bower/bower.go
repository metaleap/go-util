package udevbower

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

type BowerFile struct {
	Name        string `json:"name"`
	HomePage    string `json:"homepage,omitempty"`
	Description string `json:"description,omitempty"`
	License     string `json:"license,omitempty"`

	Repository struct {
		Type string `json:"type,omitempty"`
		URL  string `json:"url,omitempty"`
	} `json:"repository,omitempty"`
	Ignore          []string          `json:"ignore,omitempty"`
	Dependencies    map[string]string `json:"dependencies,omitempty"`
	DevDependencies map[string]string `json:"devDependencies,omitempty"`

	Version     string `json:"version,omitempty"`
	_Release    string `json:"_release,omitempty"`
	_Resolution struct {
		Type   string `json:"type,omitempty"`
		Tag    string `json:"tag,omitempty"`
		Commit string `json:"commit,omitempty"`
	} `json:"_resolution,omitempty"`
	_Source         string `json:"_source,omitempty"`
	_Target         string `json:"_target,omitempty"`
	_OriginalSource string `json:"_originalSource,omitempty"`
	_Direct         bool   `json:"_direct,omitempty"`

	repositoryUrl *url.URL
}

func (me *BowerFile) RepositoryURLParsed() (repoUrl *url.URL) {
	repoUrl, _ = url.ParseRequestURI(me.Repository.URL)
	return
}

func LoadFromFile(jsonFilePath string, intoStructWithBowerFile interface{}) (err error) {
	var jsonbytes []byte
	if jsonbytes, err = ioutil.ReadFile(jsonFilePath); err == nil {
		err = json.Unmarshal(jsonbytes, intoStructWithBowerFile)
	}
	return
}
