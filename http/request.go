package http

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Request structure
type Request struct {
	client      *Client
	headers     map[Header]string
	queryParams map[string]string
	pathParams  map[string]string
	context     context.Context
	body        interface{}
}

// NewRequest return a new instance of `Request` using the given `Client`
func NewRequest(client *Client) *Request {
	if client == nil {
		client = DefaultClient
	}

	return &Request{
		client:      client,
		headers:     make(map[Header]string),
		queryParams: make(map[string]string),
		pathParams:  make(map[string]string),
		context:     context.Background(),
		body:        nil,
	}
}

func (r *Request) SetQueryParam(name, value string) *Request {
	r.queryParams[name] = value
	return r
}

func (r *Request) SetQueryParams(params map[string]string) *Request {
	for name, value := range params {
		r.SetQueryParam(name, value)
	}
	return r
}

func (r *Request) SetPathParam(name, value string) *Request {
	r.pathParams[name] = value
	return r
}

func (r *Request) SetPathParams(params map[string]string) *Request {
	for name, value := range params {
		r.SetPathParam(name, value)
	}
	return r
}

func (r *Request) SetHeader(name Header, value string) *Request {
	r.headers[name] = value
	return r
}

func (r *Request) SetHeaders(headers map[Header]string) *Request {
	for name, value := range headers {
		r.SetHeader(name, value)
	}
	return r
}

func (r *Request) SetBearerToken(token string) *Request {
	r.SetHeader(HeaderAuthorization, fmt.Sprintf("Bearer %s", strings.Trim(token, " ")))
	return r
}

func (r *Request) SetBasicAuth(userName, password string) *Request {
	user := base64.StdEncoding.EncodeToString([]byte(userName + ":" + password))
	r.SetHeader(HeaderAuthorization, fmt.Sprintf("Basic %s", user))
	return r
}

func (r *Request) SetBody(body interface{}) *Request {
	r.body = body
	return r
}

func (r *Request) marshal(val interface{}) ([]byte, error) {
	if val == nil {
		return make([]byte, 0), nil
	}

	mime, ok := r.headers[HeaderContentType]
	if !ok {
		return nil, fmt.Errorf("%s is not defined", HeaderContentType.String())
	}

	switch NewMIME(mime) {
	case MIMEApplicationJSON, MIMEApplicationJSONCharsetUTF8:
		return json.Marshal(val)
	case MIMEApplicationXML, MIMEApplicationXMLCharsetUTF8:
		return xml.Marshal(val)
	default:
		return nil, fmt.Errorf("serializing format '%s' is not supported", val)
	}
}

func (r *Request) unmarshall(mime MIME, reader io.Reader, out interface{}) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return err
	}

	switch mime {
	case MIMEApplicationJSON, MIMEApplicationJSONCharsetUTF8:
		return json.Unmarshal(body, out)
	case MIMEApplicationXML, MIMEApplicationXMLCharsetUTF8:
		return xml.Unmarshal(body, out)
	default:
		return fmt.Errorf("deserializing format '%s' is not supported", mime.String())
	}
}

func (r *Request) Do(method Method, URL string, out interface{}) error {
	body, err := r.marshal(r.body)
	if err != nil {
		return err
	}

	params := make([]string, 0)
	for name, value := range r.queryParams {
		params = append(params, fmt.Sprintf("%s=%s", url.QueryEscape(name), url.QueryEscape(value)))
	}

	URL += fmt.Sprintf("?%s", strings.Join(params, "&"))
	for name, value := range r.pathParams {
		URL = strings.Replace(URL, name, url.QueryEscape(value), -1)
	}

	r.SetHeader(HeaderContentLength, fmt.Sprintf("%d", len(body)))
	req, err := http.NewRequest(method.String(), URL, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req = req.WithContext(r.context)
	for name, value := range r.headers {
		req.Header.Set(name.String(), value)
	}

	res, err := r.client.Do(req)
	if err != nil {
		return err
	}

	status := NewStatus(res.StatusCode)
	if !status.IsSuccess() {
		return fmt.Errorf("http request failed, got status: %s", status.String())
	}

	buf := new(bytes.Buffer)
	if _, err = io.Copy(buf, res.Body); err != nil {
		return err
	}

	if err = res.Body.Close(); err != nil {
		return err
	}

	if out == nil {
		return nil
	}

	val := res.Header.Get(HeaderContentType.String())
	if val == "" {
		return fmt.Errorf("could not auto-detect response format, header '%s' is not set", HeaderContentType.String())
	}

	return r.unmarshall(NewMIME(val), buf, out)
}

func (r *Request) Head(URL string, out interface{}) error {
	return r.Do(MethodHead, URL, out)
}

func (r *Request) Options(URL string, out interface{}) error {
	return r.Do(MethodOptions, URL, out)
}

func (r *Request) Get(URL string, out interface{}) error {
	return r.Do(MethodGet, URL, out)
}

func (r *Request) Post(URL string, out interface{}) error {
	return r.Do(MethodPost, URL, out)
}

func (r *Request) Put(URL string, out interface{}) error {
	return r.Do(MethodPut, URL, out)
}

func (r *Request) Patch(URL string, out interface{}) error {
	return r.Do(MethodPatch, URL, out)
}

func (r *Request) Delete(URL string, out interface{}) error {
	return r.Do(MethodDelete, URL, out)
}
