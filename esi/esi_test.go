package esi

import (
	"fmt"
	"io"
	"os"
	"testing"

	"github.com/fastly/compute-sdk-go/fsthttp"
)

func loadFromFixtures(name string) []byte {
	b, e := os.ReadFile("../fixtures/" + name)
	if e != nil {
		panic("The file " + name + " doesn't exist.")
	}

	return b
}

// tryMock     = loadFromFixtures("try")

func newRequest(method, uri string, body io.Reader) *fsthttp.Request {
	req, _ := fsthttp.NewRequest(method, uri, body)
	return req
}

func Test_Parse_includeMock(t *testing.T) {
	t.Parallel()
	fmt.Println(string(Parse(loadFromFixtures("include"), newRequest(fsthttp.MethodGet, "/", nil))))
}

func Test_Parse_commentMock(t *testing.T) {
	t.Parallel()
	fmt.Println(string(Parse(loadFromFixtures("comment"), newRequest(fsthttp.MethodGet, "/", nil))))
}

func Test_Parse_chooseMock(t *testing.T) {
	t.Parallel()
	fmt.Println(string(Parse(loadFromFixtures("choose"), newRequest(fsthttp.MethodGet, "/", nil))))
}

func Test_Parse_escapeMock(t *testing.T) {
	t.Parallel()
	fmt.Println(string(Parse(loadFromFixtures("escape"), newRequest(fsthttp.MethodGet, "/", nil))))
}

func Test_Parse_removeMock(t *testing.T) {
	t.Parallel()
	fmt.Println(string(Parse(loadFromFixtures("remove"), newRequest(fsthttp.MethodGet, "/", nil))))
}

func Test_Parse_varsMock(t *testing.T) {
	t.Parallel()
	fmt.Println(string(Parse(loadFromFixtures("vars"), newRequest(fsthttp.MethodGet, "/", nil))))
}

func Test_Parse_fullMock(t *testing.T) {
	t.Parallel()
	fmt.Println(string(Parse(loadFromFixtures("full.html"), newRequest(fsthttp.MethodGet, "/", nil))))
}

func Test_Parse_choose_match(t *testing.T) {
	t.Parallel()
	fmt.Println(string(Parse(loadFromFixtures("choose_match"), newRequest(fsthttp.MethodGet, "/se/sv/this-is-ikea/", nil))))
}
