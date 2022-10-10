package esi

import (
	"regexp"

	"github.com/fastly/compute-sdk-go/fsthttp"
)

const escape = "<!--esi"

var (
	escapeRg    = regexp.MustCompile("<!--esi")
	closeEscape = regexp.MustCompile("-->")
)

type escapeTag struct {
	*baseTag
}

func (e *escapeTag) process(b []byte, req *fsthttp.Request) ([]byte, int) {
	closeIdx := closeEscape.FindIndex(b)

	if closeIdx == nil {
		return nil, len(b)
	}

	e.length = closeIdx[1]
	b = b[:closeIdx[0]]

	return b, e.length
}
