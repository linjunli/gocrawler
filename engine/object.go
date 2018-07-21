// 定义抓取的字段
package engine

import "github.com/PuerkitoBio/goquery"

type RequestForGithub struct {
	URL string
	ParseFunc func(doc *goquery.Document)
}
type Repositories struct {
	RespName string
	URL string
	Stars int
	Fork int
	TodayStars string
	Description string
}
type Developers struct {
	DevName string
	Description string
	URL string
}
type YztLoan struct {
	Rate float32
	Id float64
	Period string
	Remain string
	Insurance string
}