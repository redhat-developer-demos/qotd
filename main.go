package main

import (
	"log"
	"net/http"
)

var quotes = Quotes{
	Quote{Id: 0, Quotation: "Whatever can happen at any time can happen today.", Author: "Seneca"},
	Quote{Id: 1, Quotation: "Yeah, well, that's just like, your opinion ... man.", Author: "The Dude"},
	Quote{Id: 2, Quotation: "Life is really simple, but we insist on making it complicated.", Author: "Confucius"},
	Quote{Id: 3, Quotation: "This above all, to thine own self be true.", Author: "William Shakespeare"},
	Quote{Id: 4, Quotation: "Never complain. Never explain.", Author: "Katharine Hepburn"},
	Quote{Id: 5, Quotation: "It is not only what you do, but also the attitude you bring to it, that makes you a success.", Author: "Don Schenck"},
}

func main() {
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":10000", router))
}
