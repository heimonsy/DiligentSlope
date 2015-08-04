package climber

import (
	"fmt"

	"github.com/franela/goreq"
	"github.com/sky31/DiligentSlope/logger"
)

type Climber struct {
	workChan chan *Work
}

func (self *Climber) Start(workChan chan *Work) {
	self.workChan = workChan
	go func() {
		for {
			work, ok := <-self.workChan
			if !ok {
				return
			}
			self.fetch(work.Url, make(map[string]bool), 0)
		}
	}()
}

func (self *Climber) Stop() {
	close(self.workChan)
}

func (self *Climber) fetch(url string, fetched map[string]bool, depth int) {
	res, err := goreq.Request{Uri: url}.Do()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if res.StatusCode == 301 || res.StatusCode == 302 || res.StatusCode == 307 {
		redirectUrl := res.Header.Get("Location")
		if url == "" {
			logger.Info(url + ": location empty")
			return
		}
		if fetched[redirectUrl] {
			logger.Debug(url + ": redirect cycle")
			return
		}
		self.fetch(url, depth+1)

	} else if res.StatusCode == 200 {
	}
}
