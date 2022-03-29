package jiefang

type ArticleListResponseJSON struct {
	Pages []PageJSON `json:"pages"`
}

type ArticleJSON struct {
	Pincurls   string `json:"pincurls"`
	Author     string `json:"author"`
	Subtitle   string `json:"subtitle"`
	Articleid  int    `json:"articleid"`
	Pid        int    `json:"pid"`
	ID         int    `json:"id"`
	Jdate      string `json:"jdate"`
	Title      string `json:"title"`
	Points     string `json:"points"`
	Introtitle string `json:"introtitle"`
}

type PageJSON struct {
	Imgurl      string        `json:"imgurl"`
	ArticleList []ArticleJSON `json:"articleList"`
	Pname       string        `json:"pname"`
	Width       int           `json:"width"`
	ID          int           `json:"id"`
	Pnumber     string        `json:"pnumber"`
	Jdate       string        `json:"jdate"`
	Issue2      string        `json:"issue2"`
	Height      int           `json:"height"`
}

type ArticleDetailReqResponse struct {
	Article ArticleDetailJSON `json:"article"`
}

type ArticleDetailJSON struct {
	Addtime    int64   `json:"addtime"`
	Articleid  int     `json:"articleid"`
	Author     string  `json:"author"`
	Columns    string  `json:"columns"`
	Content    string  `json:"content"`
	Deleted    bool    `json:"deleted"`
	Edittime   int64   `json:"edittime"`
	ExtArea    float64 `json:"ext_area"`
	ID         int     `json:"id"`
	Introtitle string  `json:"introtitle"`
	Jdate      string  `json:"jdate"`
	Parts      string  `json:"parts"`
	Pid        int     `json:"pid"`
	Pincurls   string  `json:"pincurls"`
	Points     string  `json:"points"`
	Slock      string  `json:"slock"`
	Source     string  `json:"source"`
	Subtitle   string  `json:"subtitle"`
	Title      string  `json:"title"`
}
