//Package trakt ...
package trakt

import (
	"github.com/jingweno/go-sawyer"
	"github.com/jingweno/go-sawyer/mediatype"
)

func newRequest(client *Client, urlStr string) (req *Request, err error) {
	sawyerReq, err := client.Client.NewRequest(urlStr)
	if err != nil {
		return
	}

	req = &Request{client: client, Request: sawyerReq}

	return
}
//Request ...
type Request struct {
	*sawyer.Request
	client *Client
}
//Head ...
func (r *Request) Head(output interface{}) (*Response, error) {
	return r.createResponse(r.Request.Head(), output)
}
//Get ...
func (r *Request) Get(output interface{}) (*Response, error) {
	return r.createResponse(r.Request.Get(), output)
}
//Post ...
func (r *Request) Post(input interface{}, output interface{}) (*Response, error) {
	r.setBody(input)
	return r.createResponse(r.Request.Post(), output)
}
//Put ...
func (r *Request) Put(input interface{}, output interface{}) (*Response, error) {
	r.setBody(input)
	return r.createResponse(r.Request.Put(), output)
}
//Delete ...
func (r *Request) Delete(output interface{}) (*Response, error) {
	return r.createResponse(r.Request.Delete(), output)
}
//Patch ...
func (r *Request) Patch(input interface{}, output interface{}) (*Response, error) {
	r.setBody(input)
	return r.createResponse(r.Request.Patch(), output)
}
//Options ...
func (r *Request) Options(output interface{}) (*Response, error) {
	return r.createResponse(r.Request.Options(), output)
}

func (r *Request) setBody(input interface{}) {
	mtype, _ := mediatype.Parse(DefaultMediaType)
	r.Request.SetBody(mtype, input)
}

func (r *Request) createResponse(sawyerResp *sawyer.Response, output interface{}) (resp *Response, err error) {
	resp, err = NewResponse(sawyerResp)
	if err == nil {
		err = sawyerResp.Decode(output)
	}

	return
}
