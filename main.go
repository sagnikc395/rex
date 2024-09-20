package main

import "fmt"

//state -> set of transitions and a transition is just a matcher and the state
//where the transition goes

type State struct {
	name        string
	startsGroup []string
	endsGroup   []string
	transitions []Transition
}

// transition represents a transition from one state to another
type Transition struct {
	toState string
	matcher Matcher
}

func NewState(name string) *State {
	return &State{
		name:        name,
		transitions: []Transition{},
		startsGroup: []string{},
		endsGroup:   []string{},
	}
}

// interface for matching condition
type Matcher interface {
	Match(input string) bool
}

func (s *State) addTransition(toState string, matcher Matcher) {
	s.transitions = append(s.transitions, Transition{toState: toState, matcher: matcher})
}

func (s *State) unshiftTransition(toState string, matcher Matcher) {
	s.transitions = append([]Transition{{toState: toState, matcher: matcher}}, s.transitions...)
}

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
