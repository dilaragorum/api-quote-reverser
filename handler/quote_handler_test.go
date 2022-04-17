package handler

import (
	"encoding/json"
	"errors"
	"github.com/dilaragorum/api-quote-reverser/model"
	"github.com/dilaragorum/api-quote-reverser/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_quoteHandler_GetQuotesReversed(t *testing.T) {
	t.Run("Error when getting quotes", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/reverse", http.NoBody)
		rec := httptest.NewRecorder()

		mockService := service.NewMockIQuoteService(gomock.NewController(t))
		mockService.EXPECT().
			GetQuotesReversed().
			Return([]model.GroupQuote{}, errors.New("")).
			Times(1)

		qh := NewQuoteHandler(mockService)

		qh.GetQuotesReversed(rec, req, nil)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

	})

	t.Run("Success when getting quotes", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/api/reverse", http.NoBody)
		rec := httptest.NewRecorder()

		var groupedQuotes []model.GroupQuote
		groupedQuotes = append(groupedQuotes, model.GroupQuote{
			Author: "Thomas Edison",
			Quotes: []string{".noitaripsrep tnecrep enin-ytenin dna noitaripsni tnecrep eno si suineG"},
		})

		mockService := service.NewMockIQuoteService(gomock.NewController(t))
		mockService.
			EXPECT().
			GetQuotesReversed().
			Return(groupedQuotes, nil).
			Times(1)

		qh := NewQuoteHandler(mockService)
		qh.GetQuotesReversed(rec, req, nil)
		assert.Equal(t, http.StatusOK, rec.Code)

		var returnedQuotes []model.GroupQuote
		json.NewDecoder(rec.Body).Decode(&returnedQuotes)

		assert.Equal(t, groupedQuotes, returnedQuotes)
	})
}
