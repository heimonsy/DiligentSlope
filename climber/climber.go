package climber

import (
	"regexp"

	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/franela/goreq"
)

type Climber struct {
	workChan  chan *Work
	workLimit chan int
}

func (self *Climber) Stop() {
	close(self.workChan)
}

func (self *Climber) Start(workChan chan *Work, workLimit int) {
	self.workChan = workChan
	self.workLimit = make(chan int, workLimit)
	for i := 0; i < workLimit; i++ {
		self.workLimit <- 1
	}

	go func() {
		for {
			_, ok := <-self.workLimit
			if !ok {
				return
			}

			work, ok := <-self.workChan
			if !ok {
				return
			}

			go func() {
				self.oneStep(work.Url)
				self.workLimit <- 1
			}()
		}
	}()
}

func (self *Climber) oneStep(url string) {
	resp, err := self.fetch(url)
	if err != nil {
		return
	}

	body, err := resp.Body.ToString()
	if err != nil {
		return
	}

	charset, err := self.parseCharset(resp, body)

	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err == nil {
		self.parseRespByDocument(document)
	}

	self.parseResp(resp)
}

func (self *Climber) parseCharset(resp *goreq.Response, body string) string {
	var str, headerCharset, htmlCharset string
	var matchs []string

	r, _ := regexp.Compile("charset=([a-zA-Z0-9\055]+)|charset='([a-zA-Z0-9\055]+)'|charset=\"([a-zA-Z0-9\055]+)\"")

	str = resp.Header.Get("Content-type")
	matchs = r.FindStringSubmatch(str)
	if len(matchs) > 0 {
		strings.Trim(" \n")
		headerCharset = matchs[1]
	}

	matchs = r.FindStringSubmatch(body)
	if len(matchs) > 0 {
		htmlCharset = matchs[1]
	}

	if htmlCharset != "" {
		return htmlCharset
	}
	if headerCharset != "" {
		return headerCharset
	}

	return "utf-8"
}

func (self *Climber) parseResp(resp *goreq.Response) (title string) {
	return
}

func (self *Climber) parseRespByDocument(doc *goquery.Document) (title string) {
	title = doc.Find("head title").Text()
	description = doc.Find("head description").Text()

	return
}

func (self *Climber) fetch(url string) (*goreq.Response, error) {
	resp, err := goreq.Request{Uri: url, MaxRedirects: 5}.Do()

	if err != nil {
		return nil, err
	}

	return resp, nil
}
