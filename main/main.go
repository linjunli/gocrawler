package main

import (
	"githubTrending/engine"
	"githubTrending/parse/yzt"
)

func main() {
	var simplerTest engine.Trending
	
	simplerTest.Run(
		engine.RequestForGithub{
			URL: "https://github.com/trending",
			ParseFunc:github.ParseForGithub,
		},
	)
}
