package url_parser

import (
	"net/url"
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
