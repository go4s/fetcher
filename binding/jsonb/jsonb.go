package jsonb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/go4s/fetcher/binding"
	"io"
	"net/http"
)

func init() {
	binding.RegisterBinding(jsonBinding{})
}

var (
	_ binding.Binding = &jsonBinding{}
)

type jsonBinding struct{}

func (jsonBinding) Name() string {
	return binding.MIMEJSON
}

func (jsonBinding) Bind(req *http.Response, obj interface{}) error {
	if req == nil || req.Body == nil {
		return fmt.Errorf("invalid response")
	}
	return decodeJSON(req.Body, obj)
}

func (jsonBinding) BindBody(body []byte, obj interface{}) error {
	return decodeJSON(bytes.NewReader(body), obj)
}

func decodeJSON(r io.Reader, obj interface{}) error {
	decoder := json.NewDecoder(r)
	if err := decoder.Decode(obj); err != nil {
		return err
	}
	return nil
}
