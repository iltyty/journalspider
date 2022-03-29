package jiefang

import (
	"encoding/json"
	"fmt"
	"github.com/iltyty/journalspider/model"
	"github.com/iltyty/journalspider/storage"
	"github.com/iltyty/journalspider/util"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// ArticleLink 代表文章详情链接
// link 为文章详情链接
// page 为文章所在版次，方便后面生成NewsList
// pageName 为文章所在版次名称，方便后面生成NewsList
type articleLink struct {
	link     string
	page     string
	pageName string
}

var dates = []string{
	"2022-03-29",
}

var newsList = &model.NewsList{}

// 根据日期构造出请求当日所有新闻信息的链接（包括所有版次的所有新闻信息如标题id等）
func toArticleListReqURL(date string) string {
	return fmt.Sprintf("https://www.jfdaily.com/staticsg/data/journal/%s/navi.json", date)
}

// 根据文章信息和版次信息构造出请求新闻详情的链接
func toArticleDetailReqURL(page PageJSON, article ArticleJSON) string {
	return fmt.Sprintf(
		"https://www.jfdaily.com/staticsg/data/journal/%s/%s/article/%d.json",
		article.Jdate, page.Pnumber, article.ID)
}

// 根据文章信息和版次信息构造出新闻网址
func toArticleURL(page string, article ArticleDetailJSON) string {
	return fmt.Sprintf(
		"https://www.jfdaily.com/staticsg/res/html/journal/detail.html?date=%s&id=%d&page=%s",
		article.Jdate, article.ID, page)
}

func parseRespJSON(response *ArticleListResponseJSON) []articleLink {
	var links []articleLink
	for _, page := range response.Pages {
		for _, article := range page.ArticleList {
			link := toArticleDetailReqURL(page, article)
			links = append(links, articleLink{link, page.Pnumber, page.Pname})
		}
	}
	return links
}

func getAllNewsLinks() []articleLink {
	var links []articleLink
	for _, date := range dates {
		url := toArticleListReqURL(date)
		res, err := http.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		response := &ArticleListResponseJSON{}
		err = json.Unmarshal(body, response)
		if err != nil {
			log.Fatal(err)
		}
		links = parseRespJSON(response)
	}
	return links
}

func getAllNewsDetail(links []articleLink) {
	for _, l := range links {
		res, err := http.Get(l.link)
		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal(err)
		}

		resp := &ArticleDetailReqResponse{}
		err = json.Unmarshal(body, resp)
		if err != nil {
			log.Fatal(err)
		}
		article := resp.Article

		page := fmt.Sprintf("第%s版次：%s", l.page, l.pageName)
		content := strings.ReplaceAll(article.Content, "<br/>", "")
		news := &model.News{
			URL:     toArticleURL(l.page, article),
			Title:   article.Title,
			Press:   "解放日报",
			Time:    util.TimestampToDate(article.Addtime),
			Page:    page,
			Author:  article.Author,
			Content: content,
		}
		newsList.Data = append(newsList.Data, news)
	}
}

func Entry() {
	links := getAllNewsLinks()
	getAllNewsDetail(links)
	storage.StoreNewsList(newsList, storage.JieFangDailyFile)
}
