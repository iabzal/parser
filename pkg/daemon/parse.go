package daemon

import (
	"bufio"
	"fmt"
	"github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/gocolly/colly"
	"io/ioutil"
	"os"
	"strings"
)

const (
	home      = 1
	twoRoom   = 2
	threeRoom = 3
)

func SearchTwoRoom(url string) {
	m := make(map[int]string)
	index := 0
	collector := colly.NewCollector()
	collector.OnHTML("section.a-list.a-search-list.a-list-with-favs div:first-child", func(e *colly.HTMLElement) {
		m[index] = e.ChildAttr("a.a-card__image", "href")
		index++
	})
	_ = collector.Visit(url)
	oldUrl := readFile(twoRoom)
	fmt.Println(strings.Trim(oldUrl, "\n"))
	fmt.Println(m[0])
	if strings.Trim(oldUrl, "\n") != m[0] {
		sendToTelegram("https://krisha.kz" + m[0])
		writeFile(m[0], twoRoom)
	}
}

func SearchThreeRoom(url string) {
	m := make(map[int]string)
	index := 0
	collector := colly.NewCollector()
	collector.OnHTML("section.a-list.a-search-list.a-list-with-favs div:first-child", func(e *colly.HTMLElement) {
		m[index] = e.ChildAttr("a.a-card__image", "href")
		index++
	})
	_ = collector.Visit(url)
	oldUrl := readFile(threeRoom)
	fmt.Println(strings.Trim(oldUrl, "\n"))
	fmt.Println(m[0])
	if strings.Trim(oldUrl, "\n") != m[0] {
		sendToTelegram("https://krisha.kz" + m[0])
		writeFile(m[0], threeRoom)
	}
}

func SearchHome(url string) {
	m := make(map[int]string)
	index := 0
	collector := colly.NewCollector()
	collector.OnHTML("section.a-list.a-search-list.a-list-with-favs div:first-child", func(e *colly.HTMLElement) {
		m[index] = e.ChildAttr("a.a-card__image", "href")
		index++
	})
	_ = collector.Visit(url)
	oldUrl := readFile(home)
	fmt.Println(strings.Trim(oldUrl, "\n"))
	fmt.Println(m[0])
	if strings.Trim(oldUrl, "\n") != m[0] {
		sendToTelegram("https://krisha.kz" + m[0])
		writeFile(m[0], home)
	}
}

func readFile(urlType int) string {
	urlFile := "url_home"
	if urlType == twoRoom {
		urlFile = "url_two"
	}
	if urlType == threeRoom {
		urlFile = "url_three"
	}

	//b, err := ioutil.ReadFile("/var/www/microservices/parser/"+ urlFile +".txt")
	b, err := ioutil.ReadFile("./" + urlFile + ".txt")
	check(err)
	return string(b)
}

func writeFile(url string, urlType int) {
	urlFile := "url_home"
	if urlType == twoRoom {
		urlFile = "url_two"
	}
	if urlType == threeRoom {
		urlFile = "url_three"
	}
	//f, err := os.Create("/var/www/microservices/parser/"+ urlFile +".txt")
	f, err := os.Create("./" + urlFile + ".txt")
	check(err)
	defer f.Close()

	w := bufio.NewWriter(f)
	_, err = fmt.Fprintf(w, "%v\n", url)
	check(err)
	_ = w.Flush()
}

func sendToTelegram(link string) {
	bot, err := tgbotapi.NewBotAPI("1922189938:AAEvFwEqVc7Mh6AcX6SO07nf2Bfc_TA0SEY")
	check(err)
	bot.Debug = true
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	msg := tgbotapi.NewMessage(314115700, link)
	_, _ = bot.Send(msg)
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
