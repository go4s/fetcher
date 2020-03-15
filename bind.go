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

// HandlerBind returns the appropriate Binding instance based on the HTTP method
// and the content type.
func HandlerBind(response *http.Response, bean interface{}) error {
	contentType := response.Header.Get(contentTypeKey)
	b := binding.Get(contentType)
	if b == nil {
		return fmt.Errorf("err failed get bind for Content-Type %s ", contentType)
	}
	return b.Bind(response, bean)
}
