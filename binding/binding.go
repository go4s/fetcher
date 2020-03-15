package binding

import (
	"net/http"
	"strings"
	"sync"
)

// Content-Type MIME of the most common data formats.
const (
	MIMEJSON              = "application/json"
	MIMEHTML              = "text/html"
	MIMEXML               = "application/xml"
	MIMEXML2              = "text/xml"
	MIMEPlain             = "text/plain"
	MIMEPOSTForm          = "application/x-www-form-urlencoded"
	MIMEMultipartPOSTForm = "multipart/form-data"
	MIMEPROTOBUF          = "application/x-protobuf"
	MIMEMSGPACK           = "application/x-msgpack"
	MIMEMSGPACK2          = "application/msgpack"
	MIMEYAML              = "application/x-yaml"
)

// Binding describes the interface which needs to be implemented for binding the
// data present in the request such as JSON request body, query parameters or
// the form POST.
type Binding interface {
	Name() string
	Bind(*http.Response, interface{}) error
}

// BindingBody adds BindBody method to Binding. BindBody is similar with Bind,
// but it reads the body from supplied bytes instead of req.Body.
type BindingBody interface {
	Binding
	BindBody([]byte, interface{}) error
}

func init() {
	once.Do(func() {
		bindings = make(map[string]Binding, 0)
	})
}

func RegisterBinding(binding Binding) {
	bindings[strings.ToUpper(binding.Name())] = binding
}

var (
	once     = sync.Once{}
	bindings map[string]Binding
)

func Get(contentType string) Binding {
	contentTypes := strings.Split(contentType, ";")
	return bindings[strings.ToUpper(contentTypes[0])]
}
