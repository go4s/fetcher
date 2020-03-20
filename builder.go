package fetcher

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	urlpkg "net/url"
	"strings"
	"sync"
)

type Modifier func(r *http.Request) error

type Builder interface {
	Build(...Modifier) FetchBuilder
}

type builder struct {
}

func (b *builder) Build(modifiers ...Modifier) FetchBuilder {
	req := &http.Request{
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
	}

	for _, modifierFn := range modifiers {
		if err := modifierFn(req); err != nil {
			return nil
		}
	}
	fetcher := manager{req: req, pool: sync.Pool{New: func() interface{} {
		return &http.Client{}
	}}}
	return &fetcher
}

// Method with check
func SetMethod(m string) Modifier {
	var method string
	var err error
	switch strings.ToUpper(m) {
	case http.MethodGet:
		method = http.MethodGet
	case http.MethodHead:
		method = http.MethodHead
	case http.MethodPost:
		method = http.MethodPost
	case http.MethodPut:
		method = http.MethodPut
	case http.MethodPatch:
		method = http.MethodPatch // RFC 5789
	case http.MethodDelete:
		method = http.MethodDelete
	case http.MethodConnect:
		method = http.MethodConnect
	case http.MethodOptions:
		method = http.MethodOptions
	case http.MethodTrace:
		method = http.MethodTrace
	default:
		err = fmt.Errorf("net/http: invalid method %q", m)
	}
	return func(r *http.Request) error {
		if err != nil {
			return err
		}
		r.Method = method
		return nil
	}
}

func NewBuilder() Builder {
	return &builder{}
}

// Url together
func SetUrl(u string) Modifier {
	url, err := urlpkg.Parse(u)
	return func(r *http.Request) error {
		if err != nil {
			return err
		}
		url.Host = removeEmptyPort(url.Host)
		r.URL = url
		r.Host = url.Host
		return nil
	}
}

// body
func SetBody(b io.Reader) Modifier {
	rc, ok := b.(io.ReadCloser)
	if !ok && b != nil {
		rc = ioutil.NopCloser(b)
	}
	return func(req *http.Request) error {
		req.Body = rc
		if b != nil {
			switch v := b.(type) {
			case *bytes.Buffer:
				req.ContentLength = int64(v.Len())
				buf := v.Bytes()
				req.GetBody = func() (io.ReadCloser, error) {
					r := bytes.NewReader(buf)
					return ioutil.NopCloser(r), nil
				}
			case *bytes.Reader:
				req.ContentLength = int64(v.Len())
				snapshot := *v
				req.GetBody = func() (io.ReadCloser, error) {
					r := snapshot
					return ioutil.NopCloser(&r), nil
				}
			case *strings.Reader:
				req.ContentLength = int64(v.Len())
				snapshot := *v
				req.GetBody = func() (io.ReadCloser, error) {
					r := snapshot
					return ioutil.NopCloser(&r), nil
				}
			default:
				// This is where we'd set it to -1 (at least
				// if body != NoBody) to mean unknown, but
				// that broke people during the Go 1.8 testing
				// period. People depend on it being 0 I
				// guess. Maybe retry later. See Issue 18117.
			}
			// For client requests, Request.ContentLength of 0
			// means either actually 0, or unknown. The only way
			// to explicitly say that the ContentLength is zero is
			// to set the Body to nil. But turns out too much code
			// depends on NewRequest returning a non-nil Body,
			// so we use a well-known ReadCloser variable instead
			// and have the http package also treat that sentinel
			// variable to mean explicitly zero.
			if req.GetBody != nil && req.ContentLength == 0 {
				req.Body = http.NoBody
				req.GetBody = func() (io.ReadCloser, error) { return http.NoBody, nil }
			}
		}
		return nil
	}

}

func SetBasicAuth(u, p string) Modifier {
	return func(r *http.Request) error {
		r.SetBasicAuth(u, p)
		return nil
	}
}

// -------------
func hasPort(s string) bool { return strings.LastIndex(s, ":") > strings.LastIndex(s, "]") }

// removeEmptyPort strips the empty port in ":port" to ""
// as mandated by RFC 3986 Section 6.2.3.
func removeEmptyPort(host string) string {
	if hasPort(host) {
		return strings.TrimSuffix(host, ":")
	}
	return host
}
