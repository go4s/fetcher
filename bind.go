package fetcher

import (
	"fmt"
	"github.com/go4s/fetcher/binding"
	_ "github.com/go4s/fetcher/binding/jsonb"
	_ "github.com/go4s/fetcher/binding/protob"
	"net/http"
)

const (
	contentTypeKey = "Content-Type"
)

func (m manager) Bind(response *http.Response, bean interface{}) error {
	contentType := response.Header.Get(contentTypeKey)
	b := binding.Get(contentType)
	if b == nil {
		return fmt.Errorf("err failed get bind for Content-Type %s ", contentType)
	}
	return b.Bind(response, bean)
}
