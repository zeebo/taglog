package taglog

import (
	"log"
	"strings"
)

type Logger struct {
	l    *MultiLogger
	tags map[string]bool
}

func New(l *log.Logger) *Logger {
	logger := &Logger{
		l:    NewMultiLogger(),
		tags: make(map[string]bool),
	}
	logger.AddLogger(l)
	return logger
}

//Tagging

func (l *Logger) Enable(tags string) {
	for _, tag := range l.splitTags(tags) {
		l.SetTag(tag, true)
	}
}

func (l *Logger) Disable(tags string) {
	for _, tag := range l.splitTags(tags) {
		l.SetTag(tag, false)
	}
}

func (l *Logger) SetTag(tag string, value bool) {
	l.tags[tag] = value
}

//Adding more loggers

func (l *Logger) AddLogger(ol *log.Logger) {
	l.l.AddLogger(ol)
}

//Error logging

func (l *Logger) Fatal(tag string, v ...interface{}) {
	tags := l.filterTags(l.splitTags(tag))
	if len(tags) > 0 {
		l.l.Fatal(v...)
	}
}
func (l *Logger) Fatalf(tag string, format string, v ...interface{}) {
	tags := l.filterTags(l.splitTags(tag))
	if len(tags) > 0 {
		l.l.Fatalf(format, v...)
	}
}
func (l *Logger) Fatalln(tag string, v ...interface{}) {
	tags := l.filterTags(l.splitTags(tag))
	if len(tags) > 0 {
		l.l.Fatalln(v...)
	}
}
func (l *Logger) Panic(tag string, v ...interface{}) {
	tags := l.filterTags(l.splitTags(tag))
	if len(tags) > 0 {
		l.l.Panic(v...)
	}
}
func (l *Logger) Panicf(tag string, format string, v ...interface{}) {
	tags := l.filterTags(l.splitTags(tag))
	if len(tags) > 0 {
		l.l.Panicf(format, v...)
	}
}
func (l *Logger) Panicln(tag string, v ...interface{}) {
	tags := l.filterTags(l.splitTags(tag))
	if len(tags) > 0 {
		l.l.Panicln(v...)
	}
}

//Normal logging

func (l *Logger) Print(tag string, v ...interface{}) {
	tags := l.filterTags(l.splitTags(tag))
	if len(tags) > 0 {
		l.l.Print(v...)
	}
}
func (l *Logger) Printf(tag string, format string, v ...interface{}) {
	tags := l.filterTags(l.splitTags(tag))
	if len(tags) > 0 {
		l.l.Printf(format, v...)
	}
}
func (l *Logger) Println(tag string, v ...interface{}) {
	tags := l.filterTags(l.splitTags(tag))
	if len(tags) > 0 {
		l.l.Println(v...)
	}
}

//Helpers

func (l *Logger) splitTags(tags string) []string {
	splits := strings.Split(tags, ",")
	ret := make([]string, 0, len(splits))
	for _, tag := range splits {
		stripped := strings.TrimSpace(tag)
		if len(stripped) > 0 {
			ret = append(ret, stripped)
		}
	}
	return ret
}

func (l *Logger) filterTags(tags []string) []string {
	ret := make([]string, 0, len(tags))
	for _, tag := range tags {
		if l.TagEnabled(tag) {
			ret = append(ret, tag)
		}
	}
	return ret
}

func (l *Logger) TagEnabled(tag string) bool {
	if val, ok := l.tags[tag]; ok {
		return val
	}
	return false
}

//Pass throughs

func (l *Logger) Flags() int {
	return l.l.Flags()
}
func (l *Logger) SetFlags(flag int) {
	l.l.SetFlags(flag)
}
func (l *Logger) Prefix() string {
	return l.l.Prefix()
}
func (l *Logger) SetPrefix(prefix string) {
	l.l.SetPrefix(prefix)
}
