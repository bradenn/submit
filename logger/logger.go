package logger

import (
	"fmt"
	"github.com/go-git/go-git/v5/plumbing/color"
	"net/http"
	"time"
)

var logger Logger = Logger{}

type Logger struct {
}

func log(format string, args ...string) {
	fmt.Printf(format, args)
}

func Info(format string, args ...interface{}) {
	timestamp := time.Now().Format(time.Stamp)
	fmt.Printf("%s[INFO]%s %s - %s%s\n", color.BoldCyan, color.Reset+color.Faint, timestamp, color.Reset,
		fmt.Sprintf(format,
			args...))
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

func Warn(format string, args ...string) {

}

func Error(format string, args ...interface{}) {
	timestamp := time.Now().Format(time.Stamp)
	fmt.Printf("%s[ERR*]%s %s - %s%s\n", color.BoldRed, color.Reset+color.Faint, timestamp, color.Reset+color.Red,
		fmt.Sprintf(format,
			args...))
}

func Panic(format string, args ...string) {

}
