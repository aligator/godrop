package log

import (
	"fmt"
	"log"
)

// StdLogger describes a logger that is compatible with the standard
// log.Logger but also logrus and others. As not to limit which loggers can and
// can't be used with the API.
//
// This interface is from https://godoc.org/github.com/Sirupsen/logrus#StdLogger
type StdLogger interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})

	Panic(...interface{})
	Panicf(string, ...interface{})
	Panicln(...interface{})
}

// GoDropLogger describes a logger that is based on the StdLogger but
// provides some additional utility methods used by GoDrop.
// A StdLogger can be wrapped using NewLogger.
type GoDropLogger interface {
	StdLogger
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
}

// NewLogger can wrap a StdLogger and just adds some more utility methods
// needed by the GoDropLogger.
// It can be used with any standard-lib compatible logger.
func NewLogger(stdLogger StdLogger) GoDropLogger {
	return goDropLogger{stdLogger}
}

// DefaultLogger just returns a new GoDropLogger based on log.Default.
func DefaultLogger() GoDropLogger {
	return NewLogger(log.Default())
}

// goDropLogger is a simple wrapper around a StdLogger which
// adds some utility methods.
type goDropLogger struct {
	StdLogger
}

func (l goDropLogger) Error(v ...interface{}) {
	l.Printf("error: %s", fmt.Sprint(v...))
}

func (l goDropLogger) Errorf(format string, v ...interface{}) {
	l.Printf(fmt.Sprintf(format, v...))
}

func (l goDropLogger) Errorln(v ...interface{}) {
	l.Errorf("%s/n", fmt.Sprint(v...))
}
