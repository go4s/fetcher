package protob

import (
	"github.com/go4s/fetcher/binding"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
)

func init() {
	binding.RegisterBinding(protobufBinding{})
}

var (
	_ binding.Binding = &protobufBinding{}
)

type protobufBinding struct{}

func (protobufBinding) Name() string {
	return binding.MIMEPROTOBUF
}

func (b protobufBinding) Bind(req *http.Response, obj interface{}) error {
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	return b.BindBody(buf, obj)
}

func (protobufBinding) BindBody(body []byte, obj interface{}) error {
	if err := proto.Unmarshal(body, obj.(proto.Message)); err != nil {
		return err
	}
	return nil
}
