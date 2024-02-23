package main

import "fmt"

type State struct {
	isEnd              bool
	transition         map[string]*State
	epsilonTransitions []*State
}

func createState(isEnd bool) *State {
	return &State{isEnd: isEnd, transition: make(map[string]*State), epsilonTransitions: []*State{}}
}

func addEpsilonTransition(from, to *State) {
	from.epsilonTransitions = append(from.epsilonTransitions, to)
}

func addTransition(from, to *State, symbol string) {
	from.transition[symbol] = to
}

type NFATransitionFunc func(*State, string) (*State, error)

type NFA struct {
	Start *State
	End   *State
}

func fromEpsilon() NFA {
	start := createState(false)
	end := createState(true)
	addEpsilonTransition(start, end)

	return NFA{
		Start: start,
		End:   end,
	}
}

func fromSymbol(symbol string) NFA {
	start := createState(false)
	end := createState(true)
	addTransition(start, end, symbol)

	return NFA{
		Start: start,
		End:   end,
	}
}

func main() {
	fmt.Println("rex 🦖")
}
