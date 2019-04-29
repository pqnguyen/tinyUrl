package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pqnguyen/tinyUrl/api/server"
	"github.com/pqnguyen/tinyUrl/config/env"
	"github.com/pqnguyen/tinyUrl/models"
	"github.com/pqnguyen/tinyUrl/services/usecase"
	"github.com/pqnguyen/tinyUrl/types/enums"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func before() {
	env.InitEnvironment(enums.Testing)
	models.InitModels()
	usecase.InitUseCase()
}

func after() {
	if err := models.ClearDB(); err != nil {
		log.Fatalf("[Error] can not drop database: %s", err)
	}
	log.Print("[INFO] drop datebase successfully!")
	if err := models.ClearCache(); err != nil {
		log.Fatalf("[Error] can not flushall cache: %s", err)
	}
	log.Print("[INFO] flushall cache successfully!")
}

func ConvertToPostBody(data map[string]interface{}) io.Reader {
	buf, err := json.Marshal(data)
	if err != nil {
		return nil
	}
	return bytes.NewReader(buf)
}

func TestRouting(t *testing.T) {
	before()
	router := server.CreateServer()
	srv := httptest.NewServer(router)

	var hash string
	handler := usecase.UrlUCase
	if u, err := handler.CreateFreeUrl("https://www.facebook.com/"); err == nil {
		hash = u.Hash
	} else {
		log.Fatalf("[ERROR] can not create free tiny url, get error: %s", err)
	}

	tt := []struct {
		name   string
		url    string
		method string
		body   map[string]interface{}
	}{
		{
			name:   "create free tiny url",
			url:    fmt.Sprintf("%s/create_free_url", srv.URL),
			method: "POST",
			body:   map[string]interface{}{"url": "https://www.facebook.com/"},
		},
		{
			name:   "redirect url",
			url:    fmt.Sprintf("%s/%s", srv.URL, hash),
			method: "GET",
		},
	}

	var (
		resp *http.Response
		err  error
	)
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.method == "POST" {
				resp, err = http.Post(tc.url, "application/json", ConvertToPostBody(tc.body))
				if err != nil {
					t.Fatalf("could not send Post request create tiny free url: %v", err)
				}
			} else {
				resp, err = http.Get(tc.url)
				if err != nil {
					t.Fatalf("could not send Get request create tiny free url: %v", err)
				}
			}
			defer resp.Body.Close()
			if resp.StatusCode != http.StatusOK {
				t.Errorf("expected status OK; got %v", resp.Status)
			}
		})
	}
	after()
}
