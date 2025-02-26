package apiserver

import (
	"net/http"
	"net/http/httptest"
	"restapi/internal/app/store/teststore"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer_HandleUsersCre(t *testing.T) {
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/users", nil)
	s := newServer(teststore.New())
	s.ServeHTTP(rec, req)
	assert.Equal(t, rec.Code, http.StatusOK)
}
