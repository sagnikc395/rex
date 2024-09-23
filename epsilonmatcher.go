package main

// EpsilonMatcher struct
type EpsilonMatcher struct {
	BaseMatcher
}

// Match implements Matcher.
func (em *EpsilonMatcher) Match(input string) bool {
	return true
}

// isEpsilon implements Matcher.
// Subtle: this method shadows the method (BaseMatcher).isEpsilon of EpsilonMatcher.BaseMatcher.
func (em *EpsilonMatcher) isEpsilon() bool {
	return true
}

func NewEpsilonMatcher() *EpsilonMatcher {
	return &EpsilonMatcher{}
}

func (em *EpsilonMatcher) Label() string {
	return "ε"
}
