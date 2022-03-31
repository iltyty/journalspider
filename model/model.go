package model

import (
	"encoding/xml"
	"sync"
)

type WebSite struct {
	Domain         string
	Urls           []string
	AllowedDomains []string
}

type News struct {
	URL     string `xml:"url"`                   // 网页地址
	Title   string `xml:"title,attr"`            // 标题
	Press   string `xml:"press"`                 // 报刊名称
	Time    string `xml:"time"`                  // 时间
	Page    string `xml:"page"`                  // 版次
	Author  string `xml:"author,attr,omitempty"` // 作者
	Content string `xml:"content"`               // 正文内容
}

type NewsList struct {
	Mutex   sync.Mutex `xml:"-"`
	XMLName xml.Name   `xml:"data"`
	Data    []*News    `xml:"news"`
}
