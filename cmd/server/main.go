package main

import (
	"github.com/dilaragorum/api-quote-reverser/handler"
	"github.com/dilaragorum/api-quote-reverser/repository"
	"github.com/dilaragorum/api-quote-reverser/service"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
)

func main() {
	router := httprouter.New()

	quoteRepository := repository.NewQuoteClient()
	quoteService := service.NewQuoteService(quoteRepository)
	quoteHandler := handler.NewQuoteHandler(quoteService)

	router.GET("/api/reverse", quoteHandler.GetQuotesReversed)

	log.Println("http server runs on :8080")
	err := http.ListenAndServe(":8080", router)
	log.Fatal(err)
}
