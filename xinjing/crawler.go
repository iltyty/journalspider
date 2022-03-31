package xinjing

import (
	"github.com/gocolly/colly"
	"github.com/iltyty/journalspider/model"
	"github.com/iltyty/journalspider/util"
	"log"
	"strings"
	"sync"
)

var xj = model.WebSite{
	Domain: "epaper.bjnews.com.cn",
	Urls: []string{
		"http://epaper.bjnews.com.cn/html/2022-03/29/node_1.htm",
	},
	AllowedDomains: []string{
		"epaper.bjnews.com.cn",
	},
}

var newsList = &model.NewsList{}

func parsePageLinks(nc *colly.Collector, elem *colly.HTMLElement) {
	elem.ForEach(pageLinkSelector, func(i int, e *colly.HTMLElement) {
		baseUrl := util.GetBaseUrl(e.Request.URL.String())
		pageUrl := baseUrl + e.Attr("href")
		err := nc.Visit(pageUrl)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func parseNewsLinks(dc *colly.Collector, elem *colly.HTMLElement) {
	elem.ForEach(newsLinkSelector, func(i int, e *colly.HTMLElement) {
		baseUrl := util.GetBaseUrl(e.Request.URL.String())
		newsUrl := baseUrl + e.Attr("href")
		err := dc.Visit(newsUrl)
		if err != nil {
			log.Fatal(err)
		}
	})
}

func parseNewsDetail(elem *colly.HTMLElement) {
	time := elem.ChildText(newsDateSelector)
	time = time[:strings.LastIndex(time, " ")+1]
	time = strings.TrimSpace(time)

	content := elem.ChildText(newsContentSelector)
	content = strings.ReplaceAll(content, "\u00a0", " ")
	content = strings.ReplaceAll(content, "\u3000", " ")

	news := &model.News{
		Time:    time,
		Press:   "新京报",
		Content: content,
		URL:     elem.Request.URL.String(),
		Page:    elem.ChildText(newsPageSelector),
		Title:   elem.ChildText(newsTitleSelector),
		Author:  elem.ChildText(newsAuthorSelector),
	}
	newsList.Data = append(newsList.Data, news)
	//util.PrintNews(news)
}

func registerCollectors() (pc *colly.Collector) {
	// Page collector, 爬取版次链接
	pc = colly.NewCollector(
		colly.AllowedDomains(xj.AllowedDomains...))
	nc := pc.Clone() // model collector, 爬取新闻链接
	dc := pc.Clone() // detailed collector, 爬取文章详情

	pc.OnHTML("html", func(elem *colly.HTMLElement) {
		parsePageLinks(nc, elem)
	})

	nc.OnHTML("html", func(elem *colly.HTMLElement) {
		parseNewsLinks(dc, elem)
	})

	dc.OnHTML("html", func(elem *colly.HTMLElement) {
		parseNewsDetail(elem)
	})

	return
}

func crawl(pc *colly.Collector) {
	for _, url := range xj.Urls {
		err := pc.Visit(url)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Entry(res *model.NewsList, wg *sync.WaitGroup) {
	pc := registerCollectors()
	crawl(pc)
	//storage.StoreNewsList(newsList, storage.BeijingNewsFile)
	util.AppendRes(res, newsList)
	wg.Done()
}
