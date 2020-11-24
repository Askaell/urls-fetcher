package url_parser

import (
	"encoding/json"
	"net/url"
)

type jsonParser struct {
	fileJSON []byte
}

type urlsList struct {
	Urls []string
}

func newJSONParser(urlsJSON []byte) *jsonParser {
	return &jsonParser{
		fileJSON: urlsJSON,
	}
}

func (j *jsonParser) ParseToURL() ([]*url.URL, error) {
	urlsList, err := j.parseJSONToStrings()
	if err != nil {
		return nil, err
	}

	urls, err := strToURL(urlsList.Urls...)

	return urls, nil
}

func (j *jsonParser) parseJSONToStrings() (*urlsList, error) {
	urls := &urlsList{}

	if err := json.Unmarshal(j.fileJSON, urls); err != nil {
		return nil, err
	}

	return urls, nil
}
