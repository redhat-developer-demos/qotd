package main

type Quote struct {
	Id        int    `json:"id"`
	Quotation string `json:"quotation"`
	Author    string `json:"author"`
}

type Quotes []Quote
