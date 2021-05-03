package main

import (
	"fmt"
	"net/http"
	"time"
)

/*var URLS = [...]string{
	"https://www.google.com/",
	"https://www.bing.com/",
	"https://www.youtube.com/",
}*/

var URLS = [...]string{
	"https://www.google.com/",
	"https://www.youtube.com/",
	"https://www.facebook.com/",
	"https://www.twitter.com/",
	"https://www.instagram.com/",
	"https://www.baidu.com/",
	"https://www.xvideos.com/",
	"https://www.wikipedia.org/",
	"https://www.xnxx.com/",
	"https://www.yahoo.com/",
	"https://www.yandex.ru/",
	"https://www.amazon.com/",
	"https://www.pornhub.com/",
	"https://www.live.com/",
	"https://www.vk.com/",
	"https://www.yahoo.co.jp/",
	"https://www.netflix.com/",
	"https://www.xhamster.com/",
	"https://www.zoom.us/",
	"https://www.reddit.com/",
	"https://www.naver.com/",
	"https://www.discord.com/",
	"https://www.office.com/",
	"https://www.linkedin.com/",
	"https://www.whatsapp.com/",
	"https://www.pinterest.com/",
	"https://www.google.com.br/",
	"https://www.ebay.com/",
	"https://www.tiktok.com/",
	"https://www.ok.ru/",
	"https://www.mail.ru/",
	"https://www.bing.com/",
	"https://www.microsoft.com/",
	"https://www.duckduckgo.com/",
	"https://www.qq.com/",
	"https://www.roblox.com/",
	"https://www.msn.com/",
	"https://www.twitch.tv/",
	"https://www.amazon.co.jp/",
	"https://www.bilibili.com/",
	"https://www.amazon.de/",
	"https://www.news.yahoo.co.jp/",
	"https://www.rakuten.co.jp/",
	"https://www.paypal.com/",
	"https://www.google.de/",
	"https://www.fandom.com/",
	"https://www.aliexpress.com/",
	"https://www.globo.com/",
	"https://www.indeed.com/",
	"https://www.samsung.com/",
}

func lookup_seq(url string) {
	start := time.Now()        // starting time
	http.Get(url)              // request
	end := time.Now()          // ending time
	duration := end.Sub(start) // duration = end - start
	fmt.Println(url, "time: ", duration)
}

func main() {

	// SEQUENTIAL LOOKUP EXECUTION
	fmt.Println("Sequential time tracking...")
	start := time.Now() // global starting time

	for _, url := range URLS { //iterating through urls and calling the lookup_seq method
		lookup_seq(url)
	}

	end := time.Now()          // global ending time
	duration := end.Sub(start) // global duration = end - start
	fmt.Println("Sequential execution: ", duration)
}
