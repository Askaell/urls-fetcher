package url_parser

import (
	"errors"
	"log"
	"net/url"
	"os"
)

type argsParser struct{}

func newArgsParser() *argsParser {
	return &argsParser{}
}

func (a *argsParser) ParseToURL() ([]*url.URL, error) {
	if len(os.Args[1:]) == 0 {
		return nil, errors.New("missing args")
	}

	urls := make([]*url.URL, len(os.Args[1:]))
	var parseFails int

	for i, arg := range os.Args[1:] {
		if len(arg) == 0 {
			log.Printf("unable to parse as url string %s\n", arg)
			parseFails++
			continue
		}

		if !hasHTTPSPrefix(arg) {
			arg = "https://" + arg
		}

		parsedURL, err := url.ParseRequestURI(arg)
		if err != nil {
			log.Printf("unable to parse as url string %s\n", arg)
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
