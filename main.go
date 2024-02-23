package main

import "fmt"

type NFAState struct {
	isEnd               bool
	transition          map[string]NFAState
	epsilionTransitions []NFAState
}

func addEpsilonTransition(from NFAState, to NFAState) {
	from.epsilionTransitions = append(from.epsilionTransitions, to)
}

func addTransition(from NFAState, to NFAState, symbol string) {
	from.transition[symbol] = to
}

func createState(env bool) NFAState {
	return NFAState{isEnd: env}
}

func fromEpsilon() (NFAState, NFAState) {
	var start NFAState = createState(false)
	var end NFAState = createState(true)
	addEpsilonTransition(start, end)
	return start, end
}
func fromSymbol(symbol string) (NFAState, NFAState) {
	var start NFAState = createState(false)
	var end NFAState = createState(true)

	addTransition(start, end, symbol)

	return start, end
}

func main() {
	fmt.Println("rex 🦖")
}
