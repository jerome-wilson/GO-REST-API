package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jerome-wilson/GO-REST-API/handlers"
)

func TestHandleBooks_GetAll(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/books/", nil)
	rw := httptest.NewRecorder()

	handlers.HandleBooks(rw, req)

	if rw.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rw.Code)
	}
}
