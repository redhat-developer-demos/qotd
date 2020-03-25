package main

import (
	"log"
	"net/http"
)

var quotes = Quotes{
	Quote{Id: 0, Quotation: "40 is the old age of youth, while 50 is the youth of old age.", Author: "Victor Hugo"},
	Quote{Id: 1, Quotation: "Knowledge is power.", Author: "Francis Bacon"},
	Quote{Id: 2, Quotation: "Life is really simple, but we insist on making it complicated.", Author: "Confucius"},
	Quote{Id: 3, Quotation: "This above all, to thine own self be true.", Author: "William Shakespeare"},
	Quote{Id: 4, Quotation: "Never complain. Never explain.", Author: "Katharine Hepburn"},
	Quote{Id: 5, Quotation: "Do be do be dooo.", Author: "Frank Sinatra"},
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
