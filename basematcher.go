package main

//base Matcher

type BaseMatcher struct{}

func (m *BaseMatcher) Matches(char rune) bool {
	return false
}

func (m *BaseMatcher) isEpsilon() bool {
	return false
}

func (m *BaseMatcher) IsEpsilon() bool {
	return false
}

func (m *BaseMatcher) Label() string {
	return "undefined-matcher"
}
