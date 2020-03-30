package fetcher

import (
	"context"
	"net/http"
)

type Fetcher interface {
	Fetch(interface{}) error
}

type FetchBuilder interface {
	Fetcher
	Builder
}
type HttpClientPoolManager interface {
	Get() *http.Client
	Close(cli *http.Client, err error)
}

type fetcher struct {
	HttpClientPoolManager
	req *http.Request
}

func (m *fetcher) Build(modifiers ...Modifier) FetchBuilder {

	for _, modifierFn := range modifiers {
		if err := modifierFn(m.req); err != nil {
			return nil
		}
	}
	return m
}
func (m *fetcher) Fetch(ret interface{}) (err error) {
	//atomic.AddInt64(&m.inflight, 1)
	var resp *http.Response
	cli := m.Get()
	defer m.Close(cli, err)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if resp, err = cli.Do(m.req.Clone(ctx)); err != nil {
		goto ERR
	}
	if err = m.Bind(resp, ret); err != nil {
		goto ERR
	}
	//atomic.AddInt64(&m.inflight, -1)
	//atomic.AddInt64(&m.success, 1)
	return nil
ERR:
	//atomic.AddInt64(&m.inflight, -1)
	//atomic.AddInt64(&m.failed, 1)
	return
}

var (
	_ FetchBuilder = &fetcher{}
)
