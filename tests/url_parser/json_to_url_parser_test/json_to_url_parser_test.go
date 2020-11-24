package json_to_url_parser_test

import (
	"encoding/json"
	"log"
	"net/url"
	"testing"

	"github.com/Askaell/urls-fetcher/pkg/url_parser"
)

type urlsList struct {
	Urls []string
}

type testCase struct {
	name           string
	inputValue     *urlsList
	expectedResult []*url.URL
}

func Test(t *testing.T) {
	var stringUrl1 = "google.com"
	var stringUrl2 = "https://yandex.ru"
	var url1, _ = url.ParseRequestURI("https://" + stringUrl1)
	var url2, _ = url.ParseRequestURI(stringUrl2)

	tests := []*testCase{
		{
			name: "JSON positive test",
			inputValue: &urlsList{
				Urls: []string{stringUrl1, stringUrl2},
			},
			expectedResult: []*url.URL{url1, url2},
		},
	}

	for _, testCase := range tests {
		jsonInput, err := json.Marshal(testCase.inputValue)
		if err != nil {
			log.Fatal(err)
		}

		parser := url_parser.NewJSONParser(jsonInput)
		jsonOutput, err := parser.ParseToURL()
		for i, result := range testCase.expectedResult {
			if *result != *jsonOutput[i] {
				t.Error(
					"\nTest case: ", testCase.name,
					"For: ", testCase.inputValue,
					"Expected: ", testCase.expectedResult,
					"Got: ", jsonOutput, "\n\n",
				)
			}
		}
	}

}
