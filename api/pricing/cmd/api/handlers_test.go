package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/agpelkey/combustion-sorc/internal/data"
)

type mockItem struct {
	item *data.Item
}

func TestHandleGetPriceByID(t *testing.T) {

	var app application

	rr := httptest.NewRecorder()
	req, err := http.NewRequest(http.MethodGet, "", nil)
	if err != nil {
		t.Error(err)
	}

	app.handleGetPriceByID(rr, req)

	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 but got %d", rr.Result().StatusCode)
	}
	defer rr.Result().Body.Close()

	expected := "item"
	b, err := ioutil.ReadAll(rr.Result().Body)

	if err != nil {
		t.Error(err)
	}

	if string(b) != expected {
		t.Errorf("expected %s but we got %s", expected, string(b))
	}

}
