package request

import (
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/mySingleLive/requi/http"
	"github.com/mySingleLive/requi/http/response"
	url2 "net/url"
)

type Type uint8

// Type of http request
const (
	GET Type = iota
	POST
	PUT
	HEAD
	DELETE
	OPTIONS
	TRACE
	PATCH
)

func (t Type) Name() string {
	switch t {
	case GET:
		return "GET"
	case POST:
		return "POST"
	case PUT:
		return "PUT"
	case HEAD:
		return "HEAD"
	case DELETE:
		return "DELETE"
	case OPTIONS:
		return "OPTIONS"
	case TRACE:
		return "TRACE"
	case PATCH:
		return "PATCH"
	}
	return "unspecified"
}

type restySend func(url string) (*resty.Response, error)

func (t Type) RestyRequestSend(request *resty.Request) restySend {
	switch t {
	case GET:
		return request.Get
	case POST:
		return request.Post
	case PUT:
		return request.Post
	case HEAD:
		return request.Head
	case DELETE:
		return request.Delete
	case OPTIONS:
		return request.Options
	case PATCH:
		return request.Patch
	}
	return nil
}

type ReqState uint8

const (
	Initialized ReqState = iota
	Sending
	Success
	Error
)

type OnEnd func(req *Req, resp *response.Resp)

type Req struct {
	Type    Type
	State   ReqState
	URL     *url2.URL
	Headers []http.Header
	Body    http.Body
	Resp    *response.Resp
	onEnd   OnEnd
}

func New(typ Type) *Req {
	return &Req{
		Type:  typ,
		State: Initialized,
	}
}

func (r *Req) OnEnd(success OnEnd) {
	r.onEnd = success
}

func (r *Req) ParseURL(rawUrl string) {
	url, err := url2.Parse(rawUrl)
	if err == nil {
		r.URL = url
	}
}

func (r *Req) Send() error {
	if r.URL == nil {
		return errors.New("no valid URL")
	}
	r.State = Sending
	// Create a Resty Client
	client := resty.New()
	request := client.R()
	if r.URL.Scheme == "" {
		r.URL.Scheme = "http"
	}
	urlText := r.URL.String()
	send := r.Type.RestyRequestSend(request)

	go func() {
		restyResp, err := send(urlText)
		resp := &response.Resp{
			RestyResp: restyResp,
			Error:     err,
		}
		r.Resp = resp
		if r.onEnd != nil {
			r.onEnd(r, resp)
		}
		r.State = Success
	}()
	return nil
}
