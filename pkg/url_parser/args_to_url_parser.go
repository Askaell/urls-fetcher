package url_parser

import (
	"errors"
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

	urls, err := strToURL(os.Args[1:]...)

	return urls, err
}
