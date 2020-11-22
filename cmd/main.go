package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"

	"github.com/Askaell/urls-fetcher/pkg/url_parser"
	"github.com/Askaell/urls-fetcher/pkg/web_scanner"
)

func main() {
	var urls []*url.URL
	var parser *url_parser.Parser

	switch {
	case len(os.Args[1:]) == 0:
		parser = url_parser.NewJSONParser(readFile("/Users/admin/go/src/web-pinger/urls.json"))
	default:
		parser = url_parser.NewArgsParser()
	}

	urls = parseURL(parser)

	webScanner := web_scanner.NewWebScanner(5)

	result := webScanner.MakeAllRequests(urls)

	fmt.Println(result)
}

func parseURL(parser *url_parser.Parser) []*url.URL {
	urls, err := parser.ParseToURL()
	if err != nil {
		log.Fatalf("fail url parse: %s\n", err)
	}

	return urls
}

func readFile(fileName string) []byte {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalf("fail file read: %s\n", err)
	}

	return bs
}
