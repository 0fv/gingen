package gingen

import (
	"fmt"
	"go/ast"
	"strings"
)

type comment string

var (
	tag = "!"

	method     = "method"
	group      = "group"
	parent     = "parent"
	middleware = "middleware"
	route      = "route"
	underline  = true
)

const (
	GET     = "GET"
	POST    = "POST"
	DELETE  = "DELETE"
	PATCH   = "PATCH"
	PUT     = "PUT"
	OPTIONS = "OPTIONS"
	HEAD    = "HEAD"
	Any     = "Any"
)

var methodList = []string{
	GET,
	POST,
	DELETE,
	PATCH,
	PUT,
	OPTIONS,
	HEAD,
	Any,
}

//UnderlineSet ...
func UnderlineSet(i string) {
	if i == "f" || i == "false" {
		underline = false
	}
}

var ErrNotDefined = func(name, abbr string) error {
	return fmt.Errorf("%v:%v not defined", name, abbr)
}

func commentCheck(c *ast.CommentGroup) bool {
	return strings.HasPrefix(c.Text(), tag)
}

func (c comment) String() string {
	return string(c)
}

func (c comment) checkGroup(structName string) (r, p string) {
	//split content
	content := strings.Split(c.String(), " ")
	for _, abbr := range content {
		k, v := kvget(abbr)
		switch k {
		case group:
			r = v
		case parent:
			p = v
		}
	}
	//check route,if route not defined,use structName
	if r == "" {
		r = "/" + structName
		if underline {
			r = snakeString(r)
		}
	}

	return
}

func (c comment) routeFuncProcess(funcName string) (r string, m []string, mw bool, err error) {
	//split content
	content := strings.Split(c.String(), " ")
	for _, abbr := range content {
		k, v := kvget(abbr)
		switch k {
		case method:
			m = strings.Split(v, ",")
		case route:
			r = v
		case middleware:
			mw = true
			return
		}
	}
	//check method,if method not defined,use funcName
	if len(m) == 0 {
		for _, v := range methodList {
			if strings.ToLower(v) == strings.ToLower(funcName) {
				m = append(m, v)
				return
			}
		}
	}
	if len(m) == 0 {
		err = ErrNotDefined(funcName, "request method")
	}
	return
}

func kvget(abbr string) (k, v string) {
	abbr = strings.ReplaceAll(abbr, "\n", "")
	abbr = strings.ReplaceAll(abbr, "\r", "")
	abbr = strings.TrimSpace(abbr)
	m := strings.Split(abbr, "=")
	if len(m) != 2 {
		return m[0], ""
	}
	return m[0], m[1]
}

func snakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]

		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data[:]))
}
