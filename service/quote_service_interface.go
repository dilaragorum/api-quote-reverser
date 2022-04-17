package service

import "github.com/dilaragorum/api-quote-reverser/model"

type IQuoteService interface {
	GetQuotesReversed() ([]model.GroupQuote, error)
}
