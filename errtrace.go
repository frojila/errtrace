package errtrace

import (
	"errors"
	"fmt"
	"runtime"
	"strings"
)

var (
	CallerSkip int = 1
)

type errorTrace struct {
	err  error
	msg  string
	fn   string
	file string
	line int
}

func (e *errorTrace) Error() string {
	var (
		err error
	)

	s := strings.Builder{}

	s.WriteString(fmt.Sprintf("\terror in \"%s\": %s \n", e.fn, e.msg))
	s.WriteString(fmt.Sprintf("\t\t\t\tat %s:%d\n", e.file, e.line))

	end := false
	err = e

	for !end {
		err = errors.Unwrap(err)

		if err != nil {
			switch errx := err.(type) {
			case *errorTrace:
				if errx.msg != "" {
					s.WriteString(fmt.Sprintf("\t\t\tcaused by \"%s\": %s \n", errx.fn, errx.msg))
				} else {
					s.WriteString(fmt.Sprintf("\t\t\tcaused by \"%s\"\n", errx.fn))
				}
				s.WriteString(fmt.Sprintf("\t\t\t\tat %s:%d", errx.file, errx.line))
			default:
				s.WriteString(fmt.Sprintf("\t\t\toriginal message: %s", err.Error()))
			}

			s.WriteString("\n")
			continue
		}

		end = true
	}

	return s.String()
}

func (e *errorTrace) Unwrap() error {
	return e.err
}

func New(msg string) *errorTrace {
	pc, file, line, _ := runtime.Caller(1)
	fn := runtime.FuncForPC(pc)
	return &errorTrace{
		err:  nil,
		msg:  msg,
		fn:   fn.Name(),
		file: file,
		line: line,
	}
}

func Wrap(err error) error {
	pc, file, line, _ := runtime.Caller(CallerSkip)
	fn := runtime.FuncForPC(pc)
	return &errorTrace{
		err:  err,
		fn:   fn.Name(),
		file: file,
		line: line,
	}
}

func Valid(err error) (isValid bool) {
	switch err.(type) {
	case *errorTrace:
		isValid = true

	}
	return
}

func Message(msg string) Wrapper {
	return &wrapper{message: msg}
}

type Wrapper interface {
	Wrap(err error) error
}

type wrapper struct {
	_       struct{}
	message string
}

func (w *wrapper) Wrap(err error) error {
	pc, file, line, _ := runtime.Caller(CallerSkip)
	fn := runtime.FuncForPC(pc)
	return &errorTrace{
		err:  err,
		msg:  w.message,
		fn:   fn.Name(),
		file: file,
		line: line,
	}
}
