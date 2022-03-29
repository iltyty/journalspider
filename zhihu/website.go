package zhihu

import "net/http"

type Website struct {
	domain         string
	homepage       string
	urls           []string
	allowedDomains []string
	cookies        []*http.Cookie
}

var Zhihu = &Website{
	domain:   "www.zhihu.com",
	homepage: "https://www.zhihu.com",
	urls: []string{
		"https://www.zhihu.com/question/318657727",
	},
	allowedDomains: []string{
		"www.zhihu.com",
	},
}
