package service

import (
	"github.com/dilaragorum/api-quote-reverser/model"
	"github.com/dilaragorum/api-quote-reverser/repository"
)

type quoteService struct {
	repository repository.IQuoteRepository
}

func NewQuoteService(qr repository.IQuoteRepository) *quoteService {
	return &quoteService{repository: qr}
}

func (qs *quoteService) GetQuotesReversed() ([]model.GroupQuote, error) {
	returnedQuotes, err := qs.repository.GetQuotes()
	if err != nil {
		return nil, err
	}

	authorQuoteMap := make(map[string][]string)

	for _, quote := range returnedQuotes {
		reversedQuoteText := reverseString(quote.Text)
		authorQuoteMap[quote.Author] = append(authorQuoteMap[quote.Author], reversedQuoteText)
	}

	var quotes []model.GroupQuote
	for key, value := range authorQuoteMap {
		quotes = append(quotes, model.GroupQuote{
			Author: key,
			Quotes: value,
		})
	}

	return quotes, nil
}

func reverseString(str string) string {
	output := ""
	for _, char := range str {
		output = string(char) + output
	}
	return output
}
