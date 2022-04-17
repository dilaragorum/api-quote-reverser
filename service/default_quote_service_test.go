package service

import (
	"errors"
	"github.com/dilaragorum/api-quote-reverser/model"
	"github.com/dilaragorum/api-quote-reverser/repository"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_quoteService_GetQuotesReversed(t *testing.T) {
	t.Run("Error when get quote reverse", func(t *testing.T) {
		repoErr := errors.New("an error occurred in repository")
		mockRepository := repository.NewMockIQuoteRepository(gomock.NewController(t))
		mockRepository.
			EXPECT().
			GetQuotes().
			Return(nil, repoErr).
			Times(1)
		sq := NewQuoteService(mockRepository)

		_, err := sq.GetQuotesReversed()
		assert.ErrorIs(t, err, repoErr)
	})
	t.Run("foo", func(t *testing.T) {
		quotes := []model.Quote{
			{
				Text:   "foo",
				Author: "Thomas Edison",
			},
			{
				Text:   "bar",
				Author: "Lao Tzu",
			},
			{
				Text:   "baz",
				Author: "Thomas Edison",
			},
		}
		expectedReversedQuotes := []model.GroupQuote{
			{
				Author: "Thomas Edison",
				Quotes: []string{"oof", "zab"},
			},
			{
				Author: "Lao Tzu",
				Quotes: []string{"rab"},
			},
		}

		mockRepository := repository.NewMockIQuoteRepository(gomock.NewController(t))
		mockRepository.
			EXPECT().
			GetQuotes().
			Return(quotes, nil).
			Times(1)

		svc := NewQuoteService(mockRepository)
		returnedReverseGroupedQuotes, err := svc.GetQuotesReversed()

		assert.Nil(t, err)
		assert.Equal(t, expectedReversedQuotes, returnedReverseGroupedQuotes)
	})
}
