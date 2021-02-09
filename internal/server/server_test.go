package server_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/Garius6/blog/internal/server"
	"github.com/Garius6/blog/internal/storage/teststorage"
	"github.com/stretchr/testify/assert"
)

func TestArticleCreateHandler(t *testing.T) {
	testCases := []struct {
		name         string
		title        string
		information  string
		expectedCode int
	}{
		{
			name:         "valid data",
			title:        "Valid",
			information:  "Valid",
			expectedCode: http.StatusSeeOther,
		},
		{
			name:         "invalid data",
			title:        "",
			information:  "",
			expectedCode: http.StatusBadRequest,
		},
	}

	store := teststorage.New()
	server := server.New(store)
	for _, tc := range testCases {
		body := url.Values{}
		body.Set("title", tc.title)
		body.Set("information", tc.information)

		req, _ := http.NewRequest(
			"POST", "/article/create",
			bytes.NewBufferString(body.Encode()),
		)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		rr := httptest.NewRecorder()
		server.ServeHTTP(rr, req)
		assert.Equal(t, tc.expectedCode, rr.Code, tc.name)
	}
}
