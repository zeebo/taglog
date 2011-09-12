package taglog

import (
	"testing"
	"bytes"
	"log"
)

func createLogger() (*Logger, *bytes.Buffer) {
	buf := new(bytes.Buffer)
	logger := log.New(buf, "", 0)
	tagLog := New(logger)

	return tagLog, buf
}

func compare(t *testing.T, buf *bytes.Buffer, expected string) {
	if expected != buf.String() {
		t.Errorf("Expected %q, got %q", expected, buf.String())
	}
	buf.Reset()
}

func TestTagged(t *testing.T) {
	//create buffer
	log, buf := createLogger()

	log.Print("footag", "foo")
	compare(t, buf, "")

	log.Enable("footag")
	log.Print("footag", "foo")
	compare(t, buf, "foo\n")

	log.Print("bartag", "bar")
	compare(t, buf, "")

	log.Enable("bartag")
	log.Print("bartag", "bar")
	compare(t, buf, "bar\n")
}

func TestDisableTag(t *testing.T) {
	log, buf := createLogger()

	log.Enable("footag")
	log.Print("footag", "foo")
	compare(t, buf, "foo\n")

	log.Disable("footag")
	log.Print("footag", "foo")
	compare(t, buf, "")
}

func TestMultiLog(t *testing.T) {
	firstlog, buf := createLogger()
	otherbuf := new(bytes.Buffer)
	other := log.New(otherbuf, "", 0)

	firstlog.AddLogger(other)
	firstlog.Enable("footag")
	firstlog.Print("footag", "foo")
	compare(t, buf, "foo\n")
	compare(t, otherbuf, "foo\n")
}

func TestMultitag(t *testing.T) {
	log, buf := createLogger()

	log.Enable("footag")
	log.Enable("bartag")
	log.Print("notag, other", "foo")
	compare(t, buf, "")

	log.Print("notag, footag", "foo")
	compare(t, buf, "foo\n")

	log.Print("footag, bartag", "foo")
	compare(t, buf, "foo\n")

	log.Print("footag,bartag", "foo")
	compare(t, buf, "foo\n")

	log.Disable("footag,bartag")
	log.Print("footag", "foo")
	compare(t, buf, "")
	log.Print("bartag", "foo")
	compare(t, buf, "")
}
