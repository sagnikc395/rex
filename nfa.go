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

// Compute checks if the given string matches the NFA
func (e *EngineNFA) Compute(input string) bool {
	// Stack holds the current position in the string and the current state
	type StackFrame struct {
		i            int
		currentState *State
	}

	stack := []StackFrame{}

	// Push the initial state to the stack
	stack = append(stack, StackFrame{i: 0, currentState: e.InitialState})

	for len(stack) > 0 {
		// Pop the top of the stack
		frame := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		currentState := frame.currentState
		i := frame.i

		// Check if the current state is one of the ending states
		for _, endState := range e.EndingStates {
			if currentState.name == endState.name {
				return true
			}
		}

		// If there's more input to process
		if i < len(input) {
			currentChar := input[i]

			// Iterate over transitions in reverse order
			for t := len(currentState.transitions) - 1; t >= 0; t-- {
				transition := currentState.transitions[t]
				matcher := transition.matcher
				toState := e.States[transition.toState]

				// Check if the matcher matches the input
				if matcher.isEpsilon() || matcher.Match(string(currentChar)) {
					// Epsilon transitions don't consume input
					nextI := i
					if !matcher.isEpsilon() {
						nextI = i + 1
					}
					// Push the next state and position onto the stack
					stack = append(stack, StackFrame{i: nextI, currentState: toState})
				}
			}
		}
	}

	return false
}
