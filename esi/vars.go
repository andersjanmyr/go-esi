package esi

import (
	"regexp"
	"strings"

	"github.com/fastly/compute-sdk-go/fsthttp"
)

const (
	httpAcceptLanguage = "HTTP_ACCEPT_LANGUAGE"
	httpCookie         = "HTTP_COOKIE"
	httpHost           = "HTTP_HOST"
	httpReferrer       = "HTTP_REFERER"
	httpUserAgent      = "HTTP_USER_AGENT"
	httpQueryString    = "QUERY_STRING"

	vars = "vars"
)

var (
	interpretedVar   = regexp.MustCompile(`\$\((.+?)(\{(.+)\}(.+)?)?\)`)
	defaultExtractor = regexp.MustCompile(`\|('|")(.+?)('|")`)
	stringExtractor  = regexp.MustCompile(`('|")(.+)('|")`)

	closeVars = regexp.MustCompile("</esi:vars>")
)

func parseVariables(b []byte, req *fsthttp.Request) string {
	interprets := interpretedVar.FindSubmatch(b)

	if interprets != nil {
		switch string(interprets[1]) {
		case httpAcceptLanguage:
			if strings.Contains(req.Header.Get("Accept-Language"), string(interprets[3])) {
				return "true"
			}
		case httpCookie:
			if c, e := req.Cookie(string(interprets[3])); e == nil && c.Value != "" {
				return c.Value
			}
		case httpHost:
			return req.Host
		case httpReferrer:
			return req.Header.Get("referer")
		case httpUserAgent:
			return req.Header.Get("user-agent")
		case httpQueryString:
			if q := req.URL.Query().Get(string(interprets[3])); q != "" {
				return q
			}
		}

		if len(interprets) > 3 {
			defaultValues := defaultExtractor.FindSubmatch(interprets[4])

			if len(defaultValues) > 2 {
				return string(defaultValues[2])
			}

			return ""
		}
	} else {
		strs := stringExtractor.FindSubmatch(b)

		if len(strs) > 2 {
			return string(strs[2])
		}
	}

	return string(b)
}

type varsTag struct {
	*baseTag
}

// Input (e.g. comment text="This is a comment." />).
func (c *varsTag) process(b []byte, req *fsthttp.Request) ([]byte, int) {
	found := closeVars.FindIndex(b)
	if found == nil {
		return nil, len(b)
	}

	c.length = found[1]

	return interpretedVar.ReplaceAllFunc(b[5:found[0]], func(b []byte) []byte {
		return []byte(parseVariables(b, req))
	}), c.length
}
