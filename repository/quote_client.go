package repository

import (
	"encoding/json"
	"fmt"
	"github.com/dilaragorum/api-quote-reverser/model"
	"net/http"
)

type quoteClient struct{}

func NewQuoteClient() *quoteClient {
	return &quoteClient{}
}

func (qc *quoteClient) GetQuotes() ([]model.Quote, error) {
	httpResponse, err := http.Get("https://type.fit/api/quotes")
	if err != nil { //isteği atarken herhangi bir problem yaşadım mı? :thinking:
		return []model.Quote{}, err
	}

	// isteği başarılı bir şekilde attım fakat isteğin sonucunda problem var mı?
	if httpResponse.StatusCode >= 299 {
		return []model.Quote{}, fmt.Errorf("response has status code %d", httpResponse.StatusCode)
	}

	var quotes []model.Quote
	err = json.NewDecoder(httpResponse.Body).Decode(&quotes)
	if err != nil {
		return []model.Quote{}, err
	}
	defer httpResponse.Body.Close() //Body'i kapatmak. Memory leak () önlenmesi için.

	return quotes, nil
}
