package main

import (
	"github.com/iltyty/journalspider/hunan"
	"github.com/iltyty/journalspider/jiefang"
	"github.com/iltyty/journalspider/model"
	"github.com/iltyty/journalspider/qingnian"
	"github.com/iltyty/journalspider/renmin"
	"github.com/iltyty/journalspider/storage"
	"github.com/iltyty/journalspider/xinjing"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	res := &model.NewsList{}

	wg.Add(5)
	go renmin.Entry(res, &wg)
	go hunan.Entry(res, &wg)
	go jiefang.Entry(res, &wg)
	go xinjing.Entry(res, &wg)
	go qingnian.Entry(res, &wg)
	wg.Wait()

	storage.StoreNewsList(res)
}
