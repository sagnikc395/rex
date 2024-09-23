package main

type CharacterMatcher struct {
	BaseMatcher
	char rune
}

func NewCharacterMatcher(char rune) *CharacterMatcher {
	return &CharacterMatcher{char: char}
}

func (cm *CharacterMatcher) Matches(char rune) bool {
	return cm.char == char
}

func (cm *CharacterMatcher) IsEpsilon() bool {
	return false
}

func (cm *CharacterMatcher) Label() string {
	return string(cm.char)
}
