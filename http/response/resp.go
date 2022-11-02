package response

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/mySingleLive/requi/http"
	"time"
)

type Resp struct {
	RestyResp  *resty.Response
	ResultText string
	Error      error
}

func (r *Resp) IsError() bool {
	return r.Error != nil
}

func (r *Resp) IsSuccess() bool {
	return !r.IsError()
}

func (r *Resp) Text() string {
	if r.ResultText == "" {
		r.ResultText = r.RestyResp.String()
	}
	str := r.ResultText
	if len(str) > 10000 {
		str = fmt.Sprintf("%s...", str[0:9999])
	}
	return str
}

func (r *Resp) Result() string {
	if r.ResultText == "" {
		r.ResultText = r.RestyResp.String()
	}
	str := r.ResultText
	return str
}

func (r *Resp) Protocol() string {
	return r.RestyResp.Proto()
}

func (r *Resp) Status() string {
	return r.RestyResp.Status()
}

func (r *Resp) Headers() []http.Header {
	var hm map[string][]string = r.RestyResp.Header()
	var headers []http.Header
	for name := range hm {
		values := hm[name]
		if values != nil && len(values) > 0 {
			for i := range values {
				value := values[i]
				headers = append(headers, http.Header{
					Name:  name,
					Value: value,
				})
			}
		}
	}
	return headers
}

func (r *Resp) Time() time.Duration {
	return r.RestyResp.Time()
}
