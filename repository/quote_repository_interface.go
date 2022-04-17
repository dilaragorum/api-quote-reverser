package repository

import "github.com/dilaragorum/api-quote-reverser/model"

type IQuoteRepository interface {
	GetQuotes() ([]model.Quote, error)
}
