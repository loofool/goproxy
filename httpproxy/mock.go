package httpproxy

import (
	"bytes"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

type MockRequestFilter struct {
}

func (f *MockRequestFilter) HandleRequest(h *Handler, args *FilterArgs, rw http.ResponseWriter, req *http.Request) (*http.Response, error) {
	statusCode, err := args.GetInt("StatusCode")
	if err != nil {
		return nil, err
	}
	header, err := args.GetHeader("Header")
	if err != nil {
		return nil, err
	}
	body, err := args.GetString("Body")
	if err != nil {
		body = ""
	}
	resp := &http.Response{
		StatusCode:    statusCode,
		ProtoMajor:    1,
		ProtoMinor:    1,
		Header:        *header,
		ContentLength: int64(len(body)),
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		Request:       req,
	}
	glog.Infof("%s \"MOCK %s %s %s\" %d %s", req.RemoteAddr, req.Method, req.URL.String(), req.Proto, resp.StatusCode, resp.Header.Get("Content-Length"))
	return resp, nil
}
