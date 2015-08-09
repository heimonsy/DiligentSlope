package climber

import (
	"io"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

type testBody struct {
	readed int
	body   []byte
}

func (self *testBody) Read(p []byte) (n int, err error) {
	bufLen := len(p)
	bodyLen := len(self.body)

	for self.readed < bodyLen && n < bufLen {
		p[n] = self.body[self.readed]
		self.readed++
		n++
	}

	if self.readed == bodyLen {
		err = io.EOF
	}

	return
}

func TestParseByDocument(t *testing.T) {
	titleExpect := "TestTitle"
	climber := Climber{}

	document, err := goquery.NewDocumentFromReader(&testBody{body: []byte("<html><head><title>" + titleExpect + "</title></head></html>")})
	if err != nil {
		t.Errorf(err.Error())
		return
	}

	title := climber.parseRespByDocument(document)
	if title != titleExpect {
		t.Errorf("title parse error want: %s, got: %s", titleExpect, title)
	}

}
