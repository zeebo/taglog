package taglog

import (
	"fmt"
	"log"
	"os"
)

type MultiLogger struct {
	loggers []*log.Logger
}

func NewMultiLogger() *multiLogger {
	return &multiLogger{
		loggers: make([]*log.Logger, 0),
	}
}

func (m *MultiLogger) AddLogger(l *log.Logger) {
	m.loggers = append(m.loggers, l)
	//Update the just added logger to be consistent with the style of the
	//rest of the loggers.
	l.SetFlags(m.loggers[0].Flags())
	l.SetPrefix(m.loggers[0].Prefix())
}

func (m *MultiLogger) Flags() int {
	return m.loggers[0].Flags()
}

func (m *MultiLogger) SetFlags(flag int) {
	for _, l := range m.loggers {
		l.SetFlags(flag)
	}
}

func (m *MultiLogger) Prefix() string {
	return m.loggers[0].Prefix()
}

func (m *MultiLogger) SetPrefix(prefix string) {
	for _, l := range m.loggers {
		l.SetPrefix(prefix)
	}
}

func (m *MultiLogger) Output(level int, data string) {
	for _, l := range m.loggers {
		l.Output(level, data)
	}
}

func (m *MultiLogger) Printf(format string, v ...interface{}) {
	m.Output(2, fmt.Sprintf(format, v...))
}

// Print calls m.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func (m *MultiLogger) Print(v ...interface{}) {
	m.Output(2, fmt.Sprint(v...))
}

// Println calls m.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func (m *MultiLogger) Println(v ...interface{}) {
	m.Output(2, fmt.Sprintln(v...))
}

func (m *MultiLogger) Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to Printf() followed by a call to os.Exit(1).
func (m *MultiLogger) Fatalf(format string, v ...interface{}) {
	m.Output(2, fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to Println() followed by a call to os.Exit(1).
func (m *MultiLogger) Fatalln(v ...interface{}) {
	m.Output(2, fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to Print() followed by a call to panic().
func (m *MultiLogger) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	m.Output(2, s)
	panic(s)
}

// Panicf is equivalent to Printf() followed by a call to panic().
func (m *MultiLogger) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	m.Output(2, s)
	panic(s)
}

// Panicln is equivalent to Println() followed by a call to panic().
func (m *MultiLogger) Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	m.Output(2, s)
	panic(s)
}