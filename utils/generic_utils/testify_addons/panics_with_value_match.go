package testify_addons

import (
	"fmt"
	"regexp"
	"runtime/debug"
)
import "github.com/stretchr/testify/assert"

type tHelper interface {
	Helper()
}

func didPanic(f assert.PanicTestFunc) (didPanic bool, message interface{}, stack string) {
	didPanic = true

	defer func() {
		message = recover()
		if didPanic {
			stack = string(debug.Stack())
		}
	}()

	// call the target function
	f()
	didPanic = false

	return
}

func matchRegexp(rx interface{}, str interface{}) bool {
	var r *regexp.Regexp
	if rr, ok := rx.(*regexp.Regexp); ok {
		r = rr
	} else {
		r = regexp.MustCompile(fmt.Sprint(rx))
	}

	return r.FindStringIndex(fmt.Sprint(str)) != nil

}

func PanicsWithValueMatch(t assert.TestingT, expected interface{}, f assert.PanicTestFunc, msgAndArgs ...interface{}) bool {
	if h, ok := t.(tHelper); ok {
		h.Helper()
	}

	funcDidPanic, panicValue, panickedStack := didPanic(f)
	if !funcDidPanic {
		return assert.Fail(t, fmt.Sprintf("func %#v should panic\n\tPanic value:\t%#v", f, panicValue), msgAndArgs...)
	}
	if !matchRegexp(expected, panicValue) {
		return assert.Fail(t, fmt.Sprintf("func %#v should panic with value:\t%#v\n\tPanic value:\t%#v\n\tPanic stack:\t%s", f, expected, panicValue, panickedStack), msgAndArgs...)
	}

	return true
}
