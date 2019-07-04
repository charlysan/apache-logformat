package apachelog

import (
	"bytes"
	"fmt"
	"net/http"
	"testing"

	"github.com/charlysan/apache-logformat/internal/logctx"
	"github.com/stretchr/testify/assert"
	"github.com/charlysan/apache-logformat/internal/httputil"
	"net/http/httptest"
)

func isDash(t *testing.T, s string) bool {
	return assert.Equal(t, "-", s, "expected dash")
}

func isEmpty(t *testing.T, s string) bool {
	return assert.Equal(t, "", s, "expected dash")
}

func TestInternalDashEmpty(t *testing.T) {
	f := func(t *testing.T, name string, dash bool, f FormatWriter) {
		var buf bytes.Buffer
		r, _ := http.NewRequest(http.MethodGet, "http://example.com", nil)
		r.Host = ""
		r.Method = ""
		r.Proto = ""
		ctx := logctx.Get(r)

		t.Run(fmt.Sprintf("%s (dash=%t)", name, dash), func(t *testing.T) {
			if !assert.NoError(t, f.WriteTo(&buf, ctx), "callback should succeed") {
				return
			}
			if dash {
				isDash(t, buf.String())
			} else {
				isEmpty(t, buf.String())
			}
		})
	}

	type dashEmptyCase struct {
		Name   string
		Dash   bool
		Format FormatWriter
	}
	cases := []dashEmptyCase{
		{Name: "Request Header", Dash: true, Format: requestHeader("foo")},
		{Name: "Response Header", Dash: true, Format: responseHeader("foo")},
		{Name: "Request Method", Dash: false, Format: requestHttpMethod},
		{Name: "Request Proto", Dash: false, Format: requestHttpProto},
		{Name: "Request RemoteAddr", Dash: true, Format: requestRemoteAddr},
		{Name: "Request Raw Query", Dash: false, Format: rawQuery},
		{Name: "Response Status", Dash: false, Format: httpStatus},
		{Name: "Request Username", Dash: true, Format: username},
		{Name: "Request Host", Dash: true, Format: requestHost},
		{Name: "Response ContentLength", Dash: true, Format: responseContentLength},
	}

	for _, c := range cases {
		f(t, c.Name, c.Dash, c.Format)
	}
}

func TestResponseWriterDefaultStatusCode(t *testing.T) {
	writer := httptest.NewRecorder()
	uut := httputil.GetResponseWriter(writer)
	if uut.StatusCode() != http.StatusOK {
		t.Fail()
	}
}
