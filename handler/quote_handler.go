package handler

import (
	"encoding/json"
	"github.com/dilaragorum/api-quote-reverser/service"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type quoteHandler struct {
	service service.IQuoteService
}

func NewQuoteHandler(qs service.IQuoteService) *quoteHandler {
	return &quoteHandler{service: qs}
}

func (qh *quoteHandler) GetQuotesReversed(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	groupsReversedQuotes, err := qh.service.GetQuotesReversed()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytesStr, err := json.Marshal(groupsReversedQuotes) // Encoding
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytesStr)
}
