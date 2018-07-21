// 爬虫任务的调度
package engine

import (
	"github.com/pkg/errors"
	"github.com/PuerkitoBio/goquery"
	"githubTrending/download"
	"fmt"
)

var (
	ErrorDocWrong = errors.New("document wrong")
)

type Trending struct {

}
type Yzt struct {

}

func (t Trending) Run(request RequestForGithub)  {
	var doc *goquery.Document
	doc, err := download.Download(request.URL)
	if err != nil {
		fmt.Println(ErrorDocWrong)
		return
	}
	if doc != nil {
		fmt.Println("Game start!")
		request.ParseFunc(doc)
	} else {
		fmt.Println("Game over!")
	}
}

func (t Yzt) Run(request RequestForGithub) {
	var doc *goquery.Document
	doc, err := download.Download(request.URL)
	if err != nil {
		fmt.Println(ErrorDocWrong)
		return
	}
	if doc != nil {
		fmt.Println("Game start yzt!")
		request.ParseFunc(doc)
	} else {
		fmt.Println("Game over yzt!")
	}
}