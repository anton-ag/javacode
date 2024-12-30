package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	nethttp "net/http"
	"net/http/httptest"

	"github.com/anton-ag/javacode/internal/http"
)

func (s *APITestSuite) TestWrongOperation() {
	r := s.Require()

	router := s.handler.Init()

	payload := http.Payload{
		ID:        s.uuid,
		Operation: "STEAL",
		Amount:    100,
	}
	marshalledPayload, err := json.Marshal(payload)
	s.NoError(err)

	req, _ := nethttp.NewRequest("POST", "/api/v1/wallet/", bytes.NewReader(marshalledPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	r.Equal(nethttp.StatusBadRequest, resp.Result().StatusCode)
}

func (s *APITestSuite) TestWrongUser() {
	r := s.Require()

	router := s.handler.Init()

	req, _ := nethttp.NewRequest("GET", fmt.Sprintf("/api/v1/wallets/%s", "s.uuid"), nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	r.Equal(nethttp.StatusNotFound, resp.Result().StatusCode)
}
