package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/iltyty/journalspider/model"
	"strings"
	"time"
)

func GetBaseUrl(url string) string {
	if !strings.Contains(url, "/") {
		return url
	}

	return url[:strings.LastIndex(url, "/")+1]
}

func PrintNews(news *model.News) {
	bs, err := json.Marshal(news)
	if err != nil {
		fmt.Println(err)
	}
	var out bytes.Buffer
	_ = json.Indent(&out, bs, "", "\t")
	fmt.Printf("%v\n", out.String())
}

func TimestampToDate(timestamp int64) string {
	ts := time.Unix(timestamp, 0)
	return ts.Format("2006-01-02")
}

func AppendRes(res *model.NewsList, newsList *model.NewsList) {
	res.Mutex.Lock()
	res.Data = append(res.Data, newsList.Data...)
	res.Mutex.Unlock()
}
