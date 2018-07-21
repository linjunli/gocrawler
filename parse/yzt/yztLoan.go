package yzt

import (
	"github.com/PuerkitoBio/goquery"
	"githubTrending/infra"
	"fmt"
)

func ParseForYzt(document *goquery.Document)  {
	document.Find("div.mod-box-sort ul.bbd li").Each(func(i int, selection *goquery.Selection) {
		Rate, _ := infra.HandleCommon(selection.Find("div.inner a.main div.hd span.value").Text())
		href, _ := selection.Find("div.inner a.main").Attr("href")
		skuid, _ := selection.Attr("data-skuid")
		Period, _ := infra.HandleCommon(selection.Find("div.inner a.main div.bd p.features strong").Text())
		Limit, _ := infra.HandleCommon(selection.Find("div.inner a.main div.bd p.features span").Text())
		Insurance, _ := infra.HandleCommon(selection.Find("div.inner a.main div.bd p.tags em").Text())
		fmt.Println(Rate,href,skuid,Period,Limit,Insurance)
	})
}