package main

// interface for matching a condition
type Matcher interface {
	Match(input string) bool

	Label() string

	isEpsilon() bool
}
