package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	. "github.com/FlorentinDUBOIS/go.util/http"
)

func TestDefaultClientGet(testing *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderContentType.String(), MIMEApplicationJSONCharsetUTF8.String())
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"foo":"bar"}`))
	}))

	defer server.Close()
	mapper := make(map[string]interface{})
	if err := DefaultClient.R().Get(server.URL, &mapper); err != nil {
		testing.Error(err)
		testing.Fail()
	}
}

func TestDefaultClientPost(testing *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set(HeaderContentType.String(), MIMEApplicationJSONCharsetUTF8.String())
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"foo":"bar"}`))
	}))

	defer server.Close()
	mapper := make(map[string]interface{})
	err := DefaultClient.
		R().
		SetHeader(HeaderContentType, MIMEApplicationJSONCharsetUTF8.String()).
		SetBody(map[string]interface{}{"foo": "bar"}).
		Post(server.URL, &mapper)

	if err != nil {
		testing.Error(err)
		testing.Fail()
	}
}
