package entrust

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestExchangeMethodGet(t *testing.T) {
	c := &Client{
		client: &http.Client{},
	}

	requestPath := "/path"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != requestPath {
			t.Errorf("got path %q expected %s", r.URL.Path, requestPath)
		}
		if r.Method != http.MethodGet {
			t.Errorf("got method %q expected %s", r.Method, http.MethodGet)
		}
		if r.ContentLength > 0 {
			t.Errorf("got content lenght %d expected 0", r.ContentLength)
		}
		fmt.Fprintln(w, "Done")
	}))
	defer ts.Close()

	APIServer = ts.URL
	_, _ = c.exchange(requestPath, http.MethodGet, nil)
}

func TestExchangeMethodPost(t *testing.T) {
	c := &Client{
		client: &http.Client{},
	}

	payload := map[string]bool{"test": true}
	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		t.FailNow()
	}
	requestPath := "/path"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != requestPath {
			t.Errorf("got path %q expected %s", r.URL.Path, requestPath)
		}
		if r.Method != http.MethodPost {
			t.Errorf("got method %q expected %s", r.Method, http.MethodPost)
		}
		pl, err := io.ReadAll(r.Body)
		r.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
		if !bytes.Equal(jsonPayload, pl) {
			t.Errorf("got payload %s expected %s", string(pl), string(jsonPayload))
		}

		fmt.Fprintln(w, "Done")
	}))
	defer ts.Close()

	APIServer = ts.URL
	_, _ = c.exchange(requestPath, http.MethodPost, payload)
}
