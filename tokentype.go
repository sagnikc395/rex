package main

type tokenType uint8

const (
	group           tokenType = iota
	bracket         tokenType = iota
	or              tokenType = iota
	repeat          tokenType = iota
	literal         tokenType = iota
	groupUncaptured tokenType = iota
)

type token struct {
	tokenType tokenType
	//payload for each token will be different
	value interface{}
}

type parseContext struct {
	//index of the character we are processing in the regex string
	pos    int
	tokens []token
}

type repeatPayload struct {
	min   int
	max   int
	token token
}
