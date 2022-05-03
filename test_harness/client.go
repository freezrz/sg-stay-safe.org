package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var (
	MyRequest []string
)

type MyHttpRequest struct {
	URL         string
	Method      string
	ContentType string
	Headers     map[string]string
	Params      map[string]string
	NoEncoding  bool
	Body        string
	Cookie      http.Cookie
	FormBody    bytes.Buffer
}

func New() *MyHttpRequest {
	req := &MyHttpRequest{}
	return req
}

func (r *MyHttpRequest) Build() *MyHttpRequest {
	if r.ContentType == "" {
		r.ContentType = "application/json; charset=utf-8"
	}
	if r.Method == "" {
		r.Method = http.MethodGet
	}
	if r.URL == "" {
		r.URL = "http://live-checkin.sg-stay-safe.org/"
	}
	return r
}

func (r *MyHttpRequest) Send(i interface{}) ([]byte, int, map[string][]string) {
	if r.URL == "" {
		log.Println("error: request URL empty")
		return make([]byte, 0), -1, nil
	}
	log.Println(fmt.Sprintf("############# %s REQUEST #############", r.Method))
	client := &http.Client{}
	var req *http.Request
	if r.FormBody.String() != "" {
		req, _ = http.NewRequest(r.Method, r.URL, &r.FormBody)
	} else {
		req, _ = http.NewRequest(r.Method, r.URL, bytes.NewBuffer([]byte(r.Body)))
	}
	if r.ContentType != "" {
		req.Header.Set("Content-Type", r.ContentType)
	}
	if r.Cookie.String() != "" {
		req.AddCookie(&r.Cookie)
	}
	for key, value := range r.Headers {
		req.Header.Set(key, value)
	}

	if len(r.Params) > 0 {
		query := req.URL.Query()
		for param, value := range r.Params {
			query.Add(param, value)
		}
		if r.NoEncoding == false {
			req.URL.RawQuery = query.Encode()
		}
	}

	log.Printf("sending request: %v\n", req.URL)
	if r.Body != "" {
		log.Printf("With body: %s", r.Body)
	}
	if len(r.Headers) > 0 {
		for k, v := range r.Headers {
			if strings.Contains(strings.ToLower(k), "auth") && len(v) > 100 {
				log.Printf("With auth: %s = %s", k, v[:100]+"...")
			} else {
				log.Printf("With header: %s = %s", k, v)
			}
		}
	}
	if len(r.Params) != 0 {
		log.Printf("With params: %s", r.Params)
	}
	var resp *http.Response
	var err error
	resp, err = client.Do(req)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	var body []byte
	var reader io.ReadCloser
	reader = resp.Body
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println("############# RESPONSE #############")
	log.Printf("response code: %s", resp.Status)
	if !strings.Contains(string(content), "DOCTYPE html") {
		log.Printf("response: %s\n", content)
	}
	body = content
	json.Unmarshal(body, i)
	return body, resp.StatusCode, resp.Header
}

func (r *MyHttpRequest) WithURL(s string) *MyHttpRequest {
	r.URL = s
	return r
}

func (r *MyHttpRequest) WithMethod(s string) *MyHttpRequest {
	r.Method = s
	return r
}

func (r *MyHttpRequest) WithContentType(s string) *MyHttpRequest {
	r.ContentType = s
	return r
}

func (r *MyHttpRequest) WithHeaders(s map[string]string) *MyHttpRequest {
	r.Headers = s
	return r
}

func (r *MyHttpRequest) WithParams(s map[string]string) *MyHttpRequest {
	r.Params = s
	return r
}

func (r *MyHttpRequest) WithNoEncoding(b bool) *MyHttpRequest {
	r.NoEncoding = b
	return r
}

func (r *MyHttpRequest) WithBody(s string) *MyHttpRequest {
	r.Body = s
	return r
}

func (r *MyHttpRequest) WithCookie(s http.Cookie) *MyHttpRequest {
	r.Cookie = s
	return r
}
