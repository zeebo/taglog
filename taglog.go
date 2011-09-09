package taglog

import (
	"log"
	"strings"
)

type Logger struct {
	l *MultiLogger
	tags map[string]bool
}

func New(l *log.Logger) {
	
}

//Tagging

func (l *Logger) Enable(tag string) {
	l.SetTag(tag, true)
}

func (l *Logger) Disable(tag string) {
	l.SetTag(tag, false)
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
	l.l.Fatal(v...)
}
func (l *Logger) Fatalf(tag string, format string, v ...interface{}) {
	
}
func (l *Logger) Fatalln(tag string, v ...interface{}) {
	
}
func (l *Logger) Panic(tag string, v ...interface{}) {
	
}
func (l *Logger) Panicf(tag string, format string, v ...interface{}) {
	
}
func (l *Logger) Panicln(tag string, v ...interface{}) {
	
}

//Normal logging

func (l *Logger) Print(tag string, v ...interface{}) {
	
}
func (l *Logger) Printf(tag string, format string, v ...interface{}) {
	
}
func (l *Logger) Println(tag string, v ...interface{}) {
	
}

//Helpers

func (l *Logger) splitTags(tag string) []string {
	
}

func (l *Logger) tagEnabled(tag string) bool {
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


type Taglist []string

func Tags(tags ...string) Taglist {
	return Taglist(tags)
}
