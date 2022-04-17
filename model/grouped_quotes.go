package model

type GroupQuote struct {
	Author string   `json:"author"`
	Quotes []string `json:"quotes"`
}
