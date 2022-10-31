package response

import (
	"fmt"
	"github.com/go-resty/resty/v2"
)

type Resp struct {
	RestyResp *resty.Response
	Error     error
}

func (r *Resp) IsError() bool {
	return r.Error != nil
}

func (r *Resp) IsSuccess() bool {
	return !r.IsError()
}

func (r *Resp) Text() string {
	str := r.RestyResp.String()
	if len(str) > 800 {
		str = fmt.Sprintf("%s...", str[0:799])
	}
	return str
}

func (r *Resp) Status() string {
	return r.RestyResp.Status()
}
