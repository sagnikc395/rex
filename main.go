package main

import "fmt"






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
