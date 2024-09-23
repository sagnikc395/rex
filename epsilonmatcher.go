package main

// EpsilonMatcher struct
type EpsilonMatcher struct {
	BaseMatcher
}

func NewEpsilonMatcher() *EpsilonMatcher {
	return &EpsilonMatcher{}
}

// An epsilon matcher always matches
func (em *EpsilonMatcher) Matches(char rune) bool {
	return true
}

func (em *EpsilonMatcher) IsEpsilon() bool {
	return true
}

func (em *EpsilonMatcher) Label() string {
	return "ε"
}
