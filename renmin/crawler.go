package renmin

import (
	"github.com/gocolly/colly"
	"github.com/iltyty/journalspider/model"
	"github.com/iltyty/journalspider/util"
	"log"
	"strings"
	"sync"
)

var pd = model.WebSite{
	Domain: "paper.people.com.cn",
	Urls: []string{
		"http://paper.people.com.cn/rmrb/html/2022-03/28/nbs.D110000renmrb_01.htm",
	},
	AllowedDomains: []string{
		"paper.people.com.cn",
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

func convertTimeFormat(time string) (res string) {
	res = time
	if strings.Contains(time, "\n") {
		res = time[strings.Index(time, "\n")+1:]
		res = strings.TrimSpace(res)
	}
	return
}

func convertAuthorFormat(author string) (res string) {
	res = author
	if strings.Contains(author, "\n") {
		res = author[:strings.Index(author, "\n")+1]
		res = strings.Replace(res, "《", "", -1)
		res = strings.Replace(res, "本报记者", "", -1)
		res = strings.Replace(res, "本报评论员", "", -1)
		res = strings.Replace(res, "本报评论部", "", -1)
		res = strings.TrimSpace(res)
	}
	return
}

func parseNewsDetail(elem *colly.HTMLElement) {
	time := convertTimeFormat(elem.ChildText(newsDateSelector))
	author := convertAuthorFormat(elem.ChildText(newsAuthorSelector))
	news := &model.News{
		Time:    time,
		Author:  author,
		Press:   "人民日报",
		URL:     elem.Request.URL.String(),
		Title:   elem.ChildText(newsTitleSelector),
		Page:    elem.ChildText(newsPageSelector),
		Content: elem.ChildText(newsContentSelector),
	}
	newsList.Data = append(newsList.Data, news)
	//util.PrintNews(news)
}

func registerCollectors() (pc *colly.Collector) {
	// Page collector, 爬取版次链接
	pc = colly.NewCollector(
		colly.AllowedDomains(pd.AllowedDomains...))
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
	for _, url := range pd.Urls {
		err := pc.Visit(url)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func Entry(res *model.NewsList, wg *sync.WaitGroup) {
	pc := registerCollectors()
	crawl(pc)
	//storage.StoreNewsList(newsList, storage.PeopleDailyFile)
	util.AppendRes(res, newsList)
	wg.Done()
}
