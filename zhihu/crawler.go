package zhihu

import (
	"bytes"
	"context"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/qrcode"
	gq "github.com/skip2/go-qrcode"
	"image"
	"io/ioutil"
	"log"
	"time"
)

type Question struct {
	title   string
	answers []Answer
}

type Answer struct {
	author   string
	time     string
	content  string
	comments []Comment
}

type Comment struct {
	author   string
	time     string
	content  string
	comments []Comment
}

func initCtx() (ctx context.Context, allocCancel, ctxCancel context.CancelFunc) {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", false),
	}
	opts = append(chromedp.DefaultExecAllocatorOptions[:], opts...)
	allocCtx, allocCancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, ctxCancel = chromedp.NewContext(allocCtx)
	return
}

func printLoginQRCode(code []byte) (err error) {
	img, _, err := image.Decode(bytes.NewReader(code))
	if err != nil {
		return
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return
	}

	res, err := qrcode.NewQRCodeReader().Decode(bmp, nil)
	if err != nil {
		return
	}

	qr, err := gq.New(res.String(), gq.Medium)
	if err != nil {
		return
	}

	fmt.Println(qr.ToSmallString(true))
	return
}

func getLoginQRCode() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		var code []byte

		time.Sleep(time.Second)
		if err = chromedp.Screenshot(loginQRCodeImgSelector(), &code, chromedp.ByQuery).Do(ctx); err != nil {
			return
		}
		if err = ioutil.WriteFile("code.png", code, 0755); err != nil {
			return
		}

		//if err = printLoginQRCode(code); err != nil {
		//	return
		//}
		return
	}
}

func loginTasks() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(Zhihu.homepage),
		chromedp.Click(loginQRCodeSelector()), // 点击二维码登录
		chromedp.WaitVisible(loginQRCodeImgSelector(), chromedp.ByQuery),
		getLoginQRCode(),
		chromedp.Evaluate("window.scrollTo(0, document.body.scrollHeight);", nil),
	}
}

func login(ctx context.Context) {
	err := chromedp.Run(ctx, loginTasks()) // 登录
	if err != nil {
		log.Fatal(err)
	}
}

func saveCookiesTasks() chromedp.ActionFunc {
	return func(ctx context.Context) (err error) {
		if err = chromedp.WaitVisible(appHeaderSelector, chromedp.ByQuery).Do(ctx); err != nil {
			return
		}

		cookies, err := network.GetAllCookies().Do(ctx)
		if err != nil {
			return
		}

		cookiesData, err := network.GetAllCookiesReturns{
			Cookies: cookies,
		}.MarshalJSON()
		if err != nil {
			return
		}

		if err = ioutil.WriteFile("cookies.json", cookiesData, 0755); err != nil {
			return
		}
		return
	}
}

func saveCookies(ctx context.Context) {
	err := chromedp.Run(ctx, saveCookiesTasks()) // 保存登录后的cookies
	if err != nil {
		log.Fatal(err)
	}
}

func zhihu() {
	ctx, allocCancel, ctxCancel := initCtx()
	defer allocCancel()
	defer ctxCancel()

	login(ctx)
	saveCookies(ctx)
}
