package esi

import (
	"regexp"

	"github.com/fastly/compute-sdk-go/fsthttp"
)

const remove = "remove"

var closeRemove = regexp.MustCompile("</esi:remove>")

type removeTag struct {
	*baseTag
}

func (r *removeTag) process(b []byte, req *fsthttp.Request) ([]byte, int) {
	closeIdx := closeRemove.FindIndex(b)
	if closeIdx == nil {
		return []byte{}, len(b)
	}

	r.length = closeIdx[1]

	return []byte{}, r.length
}
