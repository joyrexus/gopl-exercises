// See page 110.
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	// "gopl.io/ch4/github"
	"github.com/joyrexus/github"
)

const (
	DAY   = 24
	WEEK  = DAY * 7
	MONTH = DAY * 30
	YEAR  = MONTH * 12
)

var mock = flag.Bool("mock", false, "do a mock search")

// Issues holds our search function.
type Issues struct{
	search func([]string) (*github.IssuesSearchResult, error)
}

// Search returns a mock response, which we use for local testing.
func Search(terms []string) (*github.IssuesSearchResult, error) {
	var result github.IssuesSearchResult
	data, err := ioutil.ReadFile("issues.json")
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err := json.Unmarshal(data, &result); err != nil {
		log.Fatalf(err.Error())
	}
	return &result, nil
}

func print(issues []*github.Issue) {
	for _, issue := range issues {
		fmt.Printf(
			"  #%-3d %9.9s %.55s\n",
			issue.Number, issue.User.Login, issue.Title,
		)
	}
}

func categorize(issues *github.IssuesSearchResult) *categories {
	cat := new(categories)
	for _, issue := range issues.Items {
		t := time.Since(issue.CreatedAt).Hours()
		switch {
		case t < MONTH:
			cat.new = append(cat.new, issue)
		case t < YEAR:
			cat.med = append(cat.med, issue)
		default:
			cat.old = append(cat.old, issue)
		}
	}
	return cat
}

type categories struct {
	new []*github.Issue
	med []*github.Issue
	old []*github.Issue
}

func main() {
	flag.Parse()
	var issues *Issues
	if *mock == true {
		issues = &Issues{Search}
	} else {
		issues = &Issues{github.SearchIssues}
	}
	result, err := issues.search(flag.Args())
	if err != nil {
		log.Fatalf(err.Error())
	}
	cat := categorize(result) // categorize returned issues into new/med/old

	fmt.Println("\n< a month old:")
	print(cat.new)

	fmt.Println("\n< a year old:")
	print(cat.med)
}
