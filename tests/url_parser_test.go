package url_parser_test

import (
	"net/url"
	"os"
	"testing"

	"github.com/Askaell/urls-fetcher/pkg/url_parser"
)

var (
	argsParser = url_parser.NewArgsParser()
)

type testCasePositive struct {
	name           string
	inputValue     []string
	expectedResult []*url.URL
}

//here we will get an error and wait a bool value:
//1) wait true - if error is
//2) wait false - if error isn't
type testCaseNegative struct {
	name           string
	inputValue     []string
	expectedResult bool
}

func TestPositive(t *testing.T) {
	var stringUrl1 = "google.com"
	var stringUrl2 = "https://yandex.ru"
	var url1, _ = url.ParseRequestURI("https://" + stringUrl1)
	var url2, _ = url.ParseRequestURI(stringUrl2)

	var positiveTests = []*testCasePositive{
		{
			name:           "One arg, without url's scheme",
			inputValue:     []string{"", stringUrl1},
			expectedResult: []*url.URL{url1},
		},
		{
			name:           "One arg, with url's scheme",
			inputValue:     []string{"", stringUrl2},
			expectedResult: []*url.URL{url2},
		},
		{
			name:           "Several args, with url's scheme and without url's scheme",
			inputValue:     []string{"", stringUrl1, stringUrl2},
			expectedResult: []*url.URL{url1, url2},
		},
	}

	argsParser := url_parser.NewArgsParser()
	runPositiveTests(t, positiveTests, argsParser)
}

func TestNegative(t *testing.T) {
	var negativeTests = []*testCaseNegative{
		{
			name:           "Missing args. Must get an error",
			inputValue:     []string{},
			expectedResult: true,
		},
		{
			name:           "Parsing fail all urls. Must get an error",
			inputValue:     []string{"", "itIsNotAnUrl"},
			expectedResult: true,
		},
	}

	runNegativeTests(t, negativeTests, argsParser)
}

func runPositiveTests(t *testing.T, positiveTests []*testCasePositive, parser *url_parser.Parser) {
	for _, testCase := range positiveTests {
		os.Args = testCase.inputValue

		result, _ := parser.ParseToURL()

		for i, r := range result {
			if *testCase.expectedResult[i] != *r {
				t.Error(
					testCase.name, "failed!",
					"For: ", testCase.inputValue,
					"Expected: ", testCase.expectedResult,
					"Got: ", result, "\n\n",
				)
			}
		}
	}
}

func runNegativeTests(t *testing.T, negativeTests []*testCaseNegative, parser *url_parser.Parser) {
	for _, testCase := range negativeTests {
		var result bool

		os.Args = testCase.inputValue

		_, err := parser.ParseToURL()

		if err != nil {
			result = true
		}

		if !result {
			t.Error(
				"\nTest case: ", testCase.name,
				"For: ", testCase.inputValue,
				"Expected: ", testCase.expectedResult,
				"Got: ", result, "\n\n",
			)
		}
	}
}
