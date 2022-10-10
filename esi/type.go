package esi

import "github.com/fastly/compute-sdk-go/fsthttp"

type (
	tag interface {
		process([]byte, *fsthttp.Request) ([]byte, int)
	}

	baseTag struct {
		length int
	}
)

func newBaseTag() *baseTag {
	return &baseTag{length: 0}
}

func (b *baseTag) process(content []byte, _ *fsthttp.Request) ([]byte, int) {
	return []byte{}, len(content)
}
