package main

import (
	"net/http"

	"github.com/darkweak/go-esi/esi"
	"github.com/fastly/compute-sdk-go/fsthttp"
)

func main() {
	rq, _ := fsthttp.NewRequest(http.MethodGet, "domain.com/", nil)
	esi.Parse([]byte{}, rq)
}
