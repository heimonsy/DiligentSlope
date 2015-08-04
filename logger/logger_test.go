package logger

import (
	"strings"
	"testing"
)

type writeTest struct {
	conent string
}

func (self *writeTest) Write(p []byte) (n int, err error) {
	self.conent = string(p)
	return len(p), nil
}

func TestLogger(t *testing.T) {
	logFile := new(writeTest)

	InitGlobal("TEST_LOGGER", logFile)

	Debug("TEST MESSAGE")
	strs := strings.Split(logFile.conent, " -- ")
	if strs[1] != string(LOG_DEBUG) {
		t.Errorf("Debug Level Error")
	}
	if strs[3] != "TEST MESSAGE" {
		t.Errorf("Debug Message Error")
	}
}
