package url_parser

import (
	"encoding/json"
	"errors"
	"log"
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

	urls := make([]*url.URL, len(urlsList.Urls))
	var parseFails int

	for i, strURL := range urlsList.Urls {
		parsedURL, err := url.Parse(strURL)
		if err != nil {
			log.Printf("unable to parse as url string %s\n", strURL)
			parseFails++
			continue
		}

		urls[i] = parsedURL
	}

	if parseFails == len(urls) {
		return nil, errors.New("fail parse for all args")
	}

	return urls, nil
}

func (j *jsonParser) parseJSONToStrings() (*urlsList, error) {
	urls := &urlsList{}

	if err := json.Unmarshal(j.fileJSON, urls); err != nil {
		return nil, err
	}

	return urls, nil
}
