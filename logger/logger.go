package logger

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/color"
	"net/http"
	"strings"
	"time"
)

type LogType int

const (
	info = iota
	warn
	panic
)

func log(logType LogType, format string, args ...interface{}) {
	var prefix string

	switch logType {
	case info:
		prefix = fmt.Sprintf("%s[INFO]%s", color.BoldBlue, color.Reset)
		break
	case warn:
		prefix = fmt.Sprintf("%s[WARN]%s", color.BoldYellow, color.Reset)
		break
	case panic:
		prefix = fmt.Sprintf("%s[ERR*]%s", color.BoldRed, color.Reset+color.Red)
		break
	}
	var body string

	if format == "" {
		body = fmt.Sprint(args)
	} else {
		body = fmt.Sprintf(format, args...)
	}

	formatted := fmt.Sprintf("%s %s", prefix, body)
	strings.ReplaceAll(formatted, "Submit", "")
	fmt.Println(formatted)
}

func Middleware(handler http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		t1 := time.Now()
		defer func() {
			Info("%s%s %s'%s', %s", color.Green, r.Method, color.Reset, r.RequestURI, time.Since(t1))
		}()
		handler.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}

func Info(format string, args ...interface{}) {
	log(info, format, args...)
}

func Warn(args ...interface{}) {
	log(warn, "", args...)
}

func Error(args ...interface{}) {
	log(panic, "", args...)
}
