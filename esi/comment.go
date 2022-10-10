package esi

import (
	"regexp"

	"github.com/fastly/compute-sdk-go/fsthttp"
)

const comment = "comment"

var closeComment = regexp.MustCompile("/>")

type commentTag struct {
	*baseTag
}

// Input (e.g. comment text="This is a comment." />).
func (c *commentTag) process(b []byte, req *fsthttp.Request) ([]byte, int) {
	found := closeComment.FindIndex(b)
	if found == nil {
		return nil, len(b)
	}

	return []byte{}, found[1]
}
