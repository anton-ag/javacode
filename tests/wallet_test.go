package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	nethttp "net/http"
	"net/http/httptest"

	"github.com/anton-ag/javacode/internal/http"
)

func (s *APITestSuite) TestCheck() {
	r := s.Require()

	router := s.handler.Init()

	req, _ := nethttp.NewRequest("GET", fmt.Sprintf("/api/v1/wallets/%s", s.uuid), nil)

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	r.Equal(nethttp.StatusOK, resp.Result().StatusCode)

	amount, err := s.repo.Wallet.Check(s.uuid)
	s.NoError(err)

	var b http.Balance
	err = json.NewDecoder(resp.Body).Decode(&b)
	s.NoError(err)
	r.Equal(amount, b.Amount)
}

func (s *APITestSuite) TestDeposit() {
	r := s.Require()

	router := s.handler.Init()

	payload := http.Payload{
		ID:        s.uuid,
		Operation: "DEPOSIT",
		Amount:    100,
	}
	marshalledPayload, err := json.Marshal(payload)
	s.NoError(err)

	req, _ := nethttp.NewRequest("POST", "/api/v1/wallet/", bytes.NewReader(marshalledPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	r.Equal(nethttp.StatusOK, resp.Result().StatusCode)

	amount, err := s.repo.Wallet.Check(s.uuid)
	s.NoError(err)
	r.Equal(1100, amount)
}

func (s *APITestSuite) TestWithDraw() {
	r := s.Require()

	router := s.handler.Init()

	payload := http.Payload{
		ID:        s.uuid,
		Operation: "WITHDRAW",
		Amount:    1200,
	}
	marshalledPayload, err := json.Marshal(payload)
	s.NoError(err)

	req, _ := nethttp.NewRequest("POST", "/api/v1/wallet/", bytes.NewReader(marshalledPayload))
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)
	r.Equal(nethttp.StatusOK, resp.Result().StatusCode)

	amount, err := s.repo.Wallet.Check(s.uuid)
	s.NoError(err)
	r.Equal(-100, amount)
}
