package main

import "fmt"


// implementation of matcher for a single character
type CharacterMatcher struct {
	char rune
}

func NewCharacterMatcher(char rune) *CharacterMatcher {
	return &CharacterMatcher{char: char}
}

func (cm *CharacterMatcher) Match(input string) bool {
	if len(input) == 1 {
		return rune(input[0]) == cm.char
	}
	return false
}

func main() {
	nfa := NewEngineNFA()
	nfa.declareStates("q0", "q1")
	nfa.setInitialStates("q0")
	endStates := []string{"q1"}
	nfa.setEndingStates(endStates[:])
	charMatcher := NewCharacterMatcher("a")
	nfa.addTransition("q0", "q1", charMatcher)
	fmt.Println(nfa.compute("a"))

}
