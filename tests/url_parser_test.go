package url_parser_test

import (
	"net/url"
	"os"
	"testing"

	"github.com/Askaell/urls-fetcher/pkg/url_parser"
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

type testsAll struct {
	testCasesPositive []*testCasePositive
	testCasesNegative []*testCaseNegative
}

func runAllTests(t *testing.T, tests *testsAll, parser *url_parser.Parser) {
	for _, testCasePositive := range tests.testCasesPositive {
		os.Args = testCasePositive.inputValue

		result, _ := parser.ParseToURL()

		for i, r := range result {
			if *testCasePositive.expectedResult[i] != *r {
				t.Error(
					testCasePositive.name, "failed!",
					"For: ", testCasePositive.inputValue,
					"Expected: ", testCasePositive.expectedResult,
					"Got: ", result, "\n\n",
				)
			}
		}
	}
}

func TestArgsToUrlParser(t *testing.T) {
	var negativeTests = []*testCaseNegative{
		{
			name:           "missing args",
			inputValue:     []string{},
			expectedResult: true,
		},
		{
			name:           "parsing fail all urls",
			inputValue:     []string{"", "itIsNotAnUrl"},
			expectedResult: true,
		},
	}

	var stringUrl1 = "google.com"
	var stringUrl2 = "https://yandex.ru"
	var url1, _ = url.ParseRequestURI("https://" + stringUrl1)
	var url2, _ = url.ParseRequestURI(stringUrl2)

	var positiveTests = []*testCasePositive{
		{
			name:           "one arg, without url's scheme",
			inputValue:     []string{"", stringUrl1},
			expectedResult: []*url.URL{url1},
		},
		{
			name:           "one arg, with url's scheme",
			inputValue:     []string{"", stringUrl2},
			expectedResult: []*url.URL{url2},
		},
		{
			name:           "several args, with url's scheme and without url's scheme",
			inputValue:     []string{"", stringUrl1, stringUrl2},
			expectedResult: []*url.URL{url1, url2},
		},
	}

	var tests = &testsAll{
		testCasesPositive: positiveTests,
		testCasesNegative: negativeTests,
	}
	var parser = url_parser.NewArgsParser()

	runAllTests(t, tests, parser)
}
