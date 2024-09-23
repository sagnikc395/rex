package main

import "fmt"

// implementation of matcher for a single character

func main() {
	nfa := NewEngineNFA()

	// Declare states
	nfa.DeclareStates("q0", "q1", "q2", "q3")

	// Set the initial state
	nfa.SetInitialState(nfa.States["q0"])

	// Set the ending states
	nfa.SetEndingStates([]*State{nfa.States["q3"]})

	// Add transitions
	nfa.AddTransition("q0", "q1", NewCharacterMatcher('a'))
	nfa.AddTransition("q1", "q2", NewCharacterMatcher('b'))
	nfa.AddTransition("q2", "q2", NewCharacterMatcher('b'))
	nfa.AddTransition("q2", "q3", NewEpsilonMatcher())

	// Test the NFA with different inputs
	fmt.Println(nfa.Compute("abbbbbb"))  // True
	fmt.Println(nfa.Compute("aabbbbbb")) // False
	fmt.Println(nfa.Compute("ab"))       // True
	fmt.Println(nfa.Compute("a"))        // False

}
