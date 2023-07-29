package main

import (
	"database/sql"
	"encoding/json"
	//"encoding/json"
	//"io/ioutil"
	"net/http"
	"net/http/httptest"

	//"fmt"
	//"io/ioutil"
	"log"
	//"net/http"
	//"net/http/httptest"
	"os"
	"testing"

	"github.com/agpelkey/combustion-sorc/internal/data"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
)

type mockItem struct {
	ID    int
	Name  string
	Price float64
}

const (
	dbDriver = "postgres"
)

var testRepo *data.PostgresDBRepo

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, os.Getenv("PRICING_DB_DSN"))
	if err != nil {
		log.Fatal("cannot connect to db")
	}

	testRepo = data.New(conn)

	m.Run()
}

func TestGetItemByID(t *testing.T) {
	item1 := &mockItem{
		1,
		"testItem",
		5,
	}
	result, err := testRepo.GetItemByID(1)
	require.NoError(t, err)
	require.NotEmpty(t, result)

	require.Equal(t, item1.ID, result.ID)
	require.Equal(t, item1.Name, result.Name)
	require.Equal(t, item1.Price, result.Price)

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

	expected := &data.Item{}

	b := &data.Item{}
	err = json.NewDecoder(rr.Result().Body).Decode(b)

	/*
		b, err := ioutil.ReadAll(rr.Result().Body)
		if err != nil {
			t.Error(err)
		}
	*/

	if b != expected {
		t.Errorf("expected does not equal returned")
	}

}
