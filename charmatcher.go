package main

type CharacterMatcher struct {
	BaseMatcher
	char rune
}

func NewCharacterMatcher(char rune) *CharacterMatcher {
	return &CharacterMatcher{char: char}
}

func (cm *CharacterMatcher) Match(input string) bool {
	if len(input) != 1 {
		return false
	}
	return rune(input[0]) == cm.char
}

func (cm *CharacterMatcher) isEpsilon() bool {
	return false
}

func (cm *CharacterMatcher) Label() string {
	return string(cm.char)
}
