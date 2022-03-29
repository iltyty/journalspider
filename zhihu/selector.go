package zhihu

import "fmt"

func loginQRCodeSelector() string {
	return ".SignFlow-qrcodeTab"
}

func loginQRCodeImgSelector() string {
	return ".Qrcode-img"
}

const appHeaderSelector = ".AppHeader-userInfo"

func answerNthChildPrefix(answerIdx int) string {
	return fmt.Sprint(`div[class="List-item"]:nth-child(`, answerIdx, `)`)
}

func questionTitleSelector() string {
	return ".QuestionHeader-title"
}

func answerAuthorSelector(answerIdx int) string {
	return answerNthChildPrefix(answerIdx) + ` [itemprop="name"]`
}

func answerTimeSelector(answerIndex int) string {
	return answerNthChildPrefix(answerIndex) + ` .ContentItem-time span`
}
func answerContentSelector(answerIndex int) string {
	return answerNthChildPrefix(answerIndex) + ` .RichContent-inner span`
}
func commentAuthorSelector() string {

	return ".QuestionHeader-title"
}

func commentTimeSelector() string {

	return ".QuestionHeader-title"
}
func commentContentSelector() string {

	return ".QuestionHeader-title"
}
