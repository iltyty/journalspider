package storage

import (
	"bytes"
	"encoding/xml"
	"github.com/iltyty/journalspider/model"
	"io"
	"log"
	"os"
)

const (
	resDir           = "data/"
	PeopleDailyFile  = resDir + "people_daily.xml"
	HunanDailyFile   = resDir + "hunan_daily.xml"
	JieFangDailyFile = resDir + "jiefang_daily.xml"
	BeijingNewsFile  = resDir + "xinjing.xml"
	BeijingYouthFile = resDir + "qingnian.xml"
)

func exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func checkDirExist() {
	if exists(resDir) {
		return
	}
	// 先创建目录
	err := os.Mkdir(resDir, 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func StoreNewsList(newsList *model.NewsList, path string) {
	checkDirExist()

	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	_, err = io.WriteString(file, xml.Header)
	if err != nil {
		log.Fatal(err)
	}

	output, err := xml.MarshalIndent(newsList, "", "    ")
	if err != nil {
		log.Fatal(err)
	}
	output = bytes.Replace(output, []byte("&#xA;"), []byte(""), -1)
	_, err = file.Write(output)
	if err != nil {
		log.Fatal(err)
	}
}
