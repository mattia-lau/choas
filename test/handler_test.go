package test

import (
	"chaos-go/internal/ticker"
	"chaos-go/mocks"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type HandlerSuite struct {
	suite.Suite
	r          *gin.Engine
	handler    *ticker.Handler
	repository *mocks.TickerRepository
}

func (s *HandlerSuite) SetupTest() {
	s.repository = &mocks.TickerRepository{}
	s.handler = ticker.NewHandler(s.repository)

	gin.SetMode(gin.TestMode)
	s.r = gin.Default()
	ticker.Route(s.r, s.handler)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(HandlerSuite))
}

var (
	dTicker = ticker.Ticker{
		Symbol:    "BTCUSD",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	dAggregate = ticker.Aggregates{
		Symbol: "BTCUSD",
		Price:  17000.01,
		ID:     1,
		Date:   time.Now(),
	}
	dAverage = ticker.AvgPrice{
		Average: 17123.1,
	}
)

func (s *HandlerSuite) assertAggregateResponse(a *ticker.Aggregates, b *ticker.Aggregates) {
	s.Equal(a.Symbol, b.Symbol)
	s.Equal(a.Price, b.Price)
	s.Equal(a.ID, b.ID)
}

func (s *HandlerSuite) TestGetLastPrice() {
	s.repository.On("GetTicker", dAggregate.Symbol).Return(dTicker, nil)
	s.repository.On("GetLastPrice", dAggregate.Symbol).Return(dAggregate, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/last/ticker/BTCUSD", nil)

	s.r.ServeHTTP(res, req)

	s.repository.AssertCalled(s.T(), "GetTicker", dAggregate.Symbol)
	s.repository.AssertCalled(s.T(), "GetLastPrice", dAggregate.Symbol)
	s.Equal(http.StatusOK, res.Code)
	agg := &ticker.Aggregates{}
	json.Unmarshal(res.Body.Bytes(), &agg)
	s.assertAggregateResponse(&dAggregate, agg)
}

func (s *HandlerSuite) TestGetAveragePrice() {
	s.repository.On("GetTicker", dAggregate.Symbol).Return(dTicker, nil)
	s.repository.On("GetAveragePrice", dAggregate.Symbol, mock.Anything, mock.Anything).Return(dAverage, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/average/ticker/BTCUSD?start=2022-01-01T00:03:00Z&end=2022-12-06T00:03:00Z", nil)

	s.r.ServeHTTP(res, req)

	s.repository.AssertCalled(s.T(), "GetTicker", dAggregate.Symbol)
	s.repository.AssertCalled(s.T(), "GetAveragePrice", dAggregate.Symbol, mock.Anything, mock.Anything)
	s.Equal(http.StatusOK, res.Code)
	average := &ticker.AvgPrice{}
	json.Unmarshal(res.Body.Bytes(), &average)
	s.Equal(dAverage.Average, average.Average)
}

func (s *HandlerSuite) TestSpecificDatePrice() {
	s.repository.On("GetTicker", dAggregate.Symbol).Return(dTicker, nil)
	s.repository.On("GetPriceByDate", dAggregate.Symbol, mock.Anything).Return(dAggregate, nil)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "http://localhost:8080/2022-12-06T00:03:00Z/ticker/BTCUSD", nil)

	s.r.ServeHTTP(res, req)

	s.repository.AssertCalled(s.T(), "GetTicker", dAggregate.Symbol)
	s.repository.AssertCalled(s.T(), "GetPriceByDate", dAggregate.Symbol, mock.Anything)
	s.Equal(http.StatusOK, res.Code)
	agg := &ticker.Aggregates{}
	json.Unmarshal(res.Body.Bytes(), &agg)
	s.assertAggregateResponse(&dAggregate, agg)
}
