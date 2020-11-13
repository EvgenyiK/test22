package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

var correctString = []string{"John Daggett, 341 King Road, Plymouth MA", "Alice Ford, 22 East Broadway, Richmond VA"}
var correctStringAnswer = "||Massachusec: John Daggett, 341 King Road, Plymouth MA...||Virginia: Alice Ford, 22 East Broadway, Richmond VA..."
var empty = []string{"John Daggett, 341 King Road, Plymouth MA", ""}
var emptyAnswer = "||Massachusec: John Daggett, 341 King Road, Plymouth MA..."
var emptyState = []string{"John Daggett, 341 King Road, Plymouth XX"}

func TestAdress(t *testing.T) {
	type args struct{
		list []string
	}
	tests:= []struct{
		name string
		args args
		want string
		wantErr bool
	}{
		{"CorrectCase", args{list: correctString}, correctStringAnswer, false},
		{"EmptyItem", args{list: empty}, emptyAnswer, false},
		{"IncorrectStateName", args{list: emptyState}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			serv:= httptest.NewServer(http.HandlerFunc(createAdress))
			defer serv.Close()
			res, err:= http.Post(serv.URL,"" ,nil)
			if err != nil {
				log.Fatal(err)
			}
			res.Body.Close()
		})
	}
}
