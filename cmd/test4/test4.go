package main

import (
	. "BookStudy/internal/StudyInterface02"
	"log"
	"net/http"
)

func main() {
	db := Databases{"鞋子": 50, "短袖": 5}
	log.Fatal(http.ListenAndServe("localhost:23500", db))
}
