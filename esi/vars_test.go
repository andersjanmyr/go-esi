package esi

import (
	"testing"

	"github.com/fastly/compute-sdk-go/fsthttp"
)

func Test_parseVars(t *testing.T) {
	t.Parallel()
	parseVariables(logicalAndTest, newRequest(fsthttp.MethodGet, "http://domain.com", nil))
}
