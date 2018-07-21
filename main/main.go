package main

import (
	"githubTrending/engine"
	"githubTrending/parse/yzt"
)

func main() {
	//var simplerTest engine.Trending
	//
	//simplerTest.Run(
	//	engine.RequestForGithub{
	//		URL: "https://github.com/trending",
	//		ParseFunc:github.ParseForGithub,
	//	},
	//)

	var yztTest engine.Yzt
	yztTest.Run(
		engine.RequestForGithub{
			URL:"https://www.yingzt.com/invest/normalAppListB?deviceId=140034125209534465",
			ParseFunc:yzt.ParseForYzt,
		},
	)
}