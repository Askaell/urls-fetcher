package url_parser

import (
	"errors"
	"log"
	"net/url"
	"strings"
)

type urlParser interface {
	ParseToURL() ([]*url.URL, error)
}

type Parser struct {
	urlParser
}

func NewArgsParser() *Parser {
	return &Parser{
		urlParser: newArgsParser(),
	}
}

func NewJSONParser(urlsJSON []byte) *Parser {
	return &Parser{
		urlParser: newJSONParser(urlsJSON),
	}
}

func hasHTTPSPrefix(url string) bool {
	return strings.HasPrefix(url, "https://")
}

func strToURL(str ...string) ([]*url.URL, error) {
	urls := make([]*url.URL, len(str))
	var parseFails int

	for i, strURL := range str {
		if len(strURL) == 0 {
			log.Printf("unable to parse as url string %s\n", strURL)
			parseFails++
			continue
		}

		if !hasHTTPSPrefix(strURL) {
			strURL = "https://" + strURL
		}

		parsedURL, err := url.ParseRequestURI(strURL)
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
