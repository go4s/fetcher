package fetcher

import (
	"encoding/json"
	"net/http"
	"strings"
	"testing"
)

type ResponseBody struct {
	Code   int             `json:"code"`
	Reason string          `json:"reason"`
	Body   json.RawMessage `json:"body"`
}

func TestNewBuilder(t *testing.T) {
	b := NewBuilder()
	f := b.Build(
		SetMethod(http.MethodGet),
		SetUrl("127.0.0.1:80/json"),
		SetBody(nil),
	)
	if f == nil {
		t.Fatal("should built")
	}
}

func TestManager_Fetch(t *testing.T) {
	b := NewBuilder()
	f := b.Build(
		SetMethod(http.MethodGet),
		SetUrl("http://127.0.0.1:8080/json"),
		SetBody(strings.NewReader("ppsp")),
	)
	if f == nil {
		t.Fatal("should built")
	}
	var resp ResponseBody
	if err := f.Fetch(&resp); err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
	var resp2 ResponseBody
	f.Build(
		SetBody(strings.NewReader("s2")),
	)
	if err := f.Fetch(&resp2); err != nil {
		t.Fatal(err)
	}
	t.Log(f.(*fetcher).req)
	t.Log(resp)
}

func BenchmarkManager_Fetch(b *testing.B) {
	bdr := NewBuilder()
	f := bdr.Build(
		SetMethod(http.MethodGet),
		SetUrl("http://127.0.0.1:8080/json"),
		SetBody(strings.NewReader("ppsp")),
	)
	if f == nil {
		b.Fatal("should built")
	}
	var resp ResponseBody
	for i := 0; i < b.N; i++ {
		if err := f.Fetch(&resp); err != nil {
			b.Fatal(err)
		}
	}
}
