package main

//main nfa engine class

type EngineNFA struct {
	States       map[string]*State
	InitialState *State
	EndingStates []*State
}

// create a new engine NFA
func NewEngineNFA() *EngineNFA {
	return &EngineNFA{
		States: make(map[string]*State),
	}
}

// SetInitialState sets the initial state of the NFA
func (e *EngineNFA) SetInitialState(state *State) {
	e.InitialState = state
}

// SetEndingStates sets the ending states of the NFA
func (e *EngineNFA) SetEndingStates(states []*State) {
	e.EndingStates = states
}

// AddState adds a new state to the NFA
func (e *EngineNFA) AddState(name string) {
	e.States[name] = NewState(name)
}

// DeclareStates adds multiple states to the NFA
func (e *EngineNFA) DeclareStates(names ...string) {
	for _, name := range names {
		e.AddState(name)
	}
}

// AddTransition adds a new transition to the NFA
func (e *EngineNFA) AddTransition(fromState, toState string, matcher Matcher) {
	from := e.States[fromState]
	from.addTransition(toState, matcher)
}

// UnshiftTransition adds a new transition with high priority to the NFA
func (e *EngineNFA) UnshiftTransition(fromState, toState string, matcher Matcher) {
	from := e.States[fromState]
	from.unshiftTransition(toState, matcher)
}
