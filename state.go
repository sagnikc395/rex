package main

type State struct {
	name        string
	startsGroup []string
	endsGroup   []string
	transitions []Transition
}

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

//methods on transitions

func (s *State) addTransition(toState string, matcher Matcher) {
	s.transitions = append(s.transitions, Transition{toState: toState, matcher: matcher})
}

func (s *State) unshiftTransition(toState string, matcher Matcher) {
	s.transitions = append([]Transition{{toState: toState, matcher: matcher}}, s.transitions...)
}
