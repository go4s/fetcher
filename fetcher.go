package fetcher

import (
	"net/http"
	"sync"
)

type Fetcher interface {
	Fetch(interface{}) error
}

type manager struct {
	//inflight        int64
	//success, failed int64
	pool sync.Pool
	*builder
}

func (m *manager) getClient() *http.Client {
	client := m.pool.Get()
	cli, ok := client.(*http.Client)
	if !ok {
		cli = &http.Client{}
	}
	// todo decorate cli
	return cli
}

func (m *manager) closeClient(cli *http.Client, err error) {
	// todo clean cli
	m.pool.Put(cli)
}

func (m *manager) Fetch(ret interface{}) (err error) {
	//atomic.AddInt64(&m.inflight, 1)
	var resp *http.Response
	cli := m.getClient()
	defer m.closeClient(cli, err)
	if resp, err = cli.Do(m.builder.request); err != nil {
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
	_ Fetcher = &manager{}
)
