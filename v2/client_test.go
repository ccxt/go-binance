package binance

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type baseTestSuite struct {
	suite.Suite
	client    *mockedClient
	apiKey    string
	secretKey string
}

func (s *baseTestSuite) r() *require.Assertions {
	return s.Require()
}

func (s *baseTestSuite) SetupTest() {
	s.apiKey = "dummyAPIKey"
	s.secretKey = "dummySecretKey"
	s.client = newMockedClient(s.apiKey, s.secretKey)
}

func (s *baseTestSuite) mockDo(data []byte, err error, statusCode ...int) {
	s.client.Client.do = s.client.do
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	s.client.On("do", anyHTTPRequest()).Return(newHTTPResponse(data, code), err)
}

func (s *baseTestSuite) assertDo() {
	s.client.AssertCalled(s.T(), "do", anyHTTPRequest())
}

func (s *baseTestSuite) assertReq(f func(r *request)) {
	s.client.assertReq = f
}

func (s *baseTestSuite) assertRequestEqual(e, a *request) {
	s.assertURLValuesEqual(e.query, a.query)
	s.assertURLValuesEqual(e.form, a.form)
}

func (s *baseTestSuite) assertURLValuesEqual(e, a url.Values) {
	var eKeys, aKeys []string
	for k := range e {
		eKeys = append(eKeys, k)
	}
	for k := range a {
		aKeys = append(aKeys, k)
	}
	r := s.r()
	r.Len(aKeys, len(eKeys))
	for k := range a {
		switch k {
		case timestampKey, signatureKey:
			r.NotEmpty(a.Get(k))
			continue
		}
		r.Equal(e[k], a[k], k)
	}
}

func anythingOfType(t string) mock.AnythingOfTypeArgument {
	return mock.AnythingOfType(t)
}

func newContext() context.Context {
	return context.Background()
}

func anyHTTPRequest() mock.AnythingOfTypeArgument {
	return anythingOfType("*http.Request")
}

func newHTTPResponse(data []byte, statusCode int) *http.Response {
	return &http.Response{
		Body:       io.NopCloser(bytes.NewBuffer(data)),
		StatusCode: statusCode,
	}
}

func newRequest() *request {
	r := &request{
		query: url.Values{},
		form:  url.Values{},
	}
	return r
}

func newSignedRequest() *request {
	return newRequest().setParams(params{
		timestampKey: "",
		signatureKey: "",
	})
}

type assertReqFunc func(r *request)

type mockedClient struct {
	mock.Mock
	*Client
	assertReq assertReqFunc
}

func newMockedClient(apiKey, secretKey string) *mockedClient {
	m := new(mockedClient)
	m.Client = NewClient(apiKey, secretKey)
	return m
}

func (m *mockedClient) do(req *http.Request) (*http.Response, error) {
	if m.assertReq != nil {
		r := newRequest()
		r.query = req.URL.Query()
		if req.Body != nil {
			bs := make([]byte, req.ContentLength)
			for {
				n, _ := req.Body.Read(bs)
				if n == 0 {
					break
				}
			}
			form, err := url.ParseQuery(string(bs))
			if err != nil {
				panic(err)
			}
			r.form = form
		}
		m.assertReq(r)
	}
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

func (s *baseTestSuite) assertTradeV3Equal(e, a *TradeV3) {
	r := s.r()
	r.Equal(e.ID, a.ID, "ID")
	r.Equal(e.Symbol, a.Symbol, "Symbol")
	r.Equal(e.OrderID, a.OrderID, "OrderID")
	r.Equal(e.OrderListId, a.OrderListId, "OrderListId")
	r.Equal(e.Price, a.Price, "Price")
	r.Equal(e.Quantity, a.Quantity, "Quantity")
	r.Equal(e.QuoteQuantity, a.QuoteQuantity, "QuoteQuantity")
	r.Equal(e.Commission, a.Commission, "Commission")
	r.Equal(e.CommissionAsset, a.CommissionAsset, "CommissionAsset")
	r.Equal(e.Time, a.Time, "Time")
	r.Equal(e.IsBuyer, a.IsBuyer, "IsBuyer")
	r.Equal(e.IsMaker, a.IsMaker, "IsMaker")
	r.Equal(e.IsBestMatch, a.IsBestMatch, "IsBestMatch")
}

func TestFormatTimestamp(t *testing.T) {
	tm, _ := time.Parse("2006-01-02 15:04:05", "2018-06-01 01:01:01")
	assert.Equal(t, int64(1527814861000), FormatTimestamp(tm))
}
