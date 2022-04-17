package model

/*
	{
		"text": "Genius is one percent inspiration and ninety-nine percent perspiration.",
		"author": "Thomas Edison"
	}
*/

type Quote struct {
	Text   string `json:"text"`
	Author string `json:"author"`
}
