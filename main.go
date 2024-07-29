package main


type State struct {
	name string
	transitions [][]string
	startsGroups []string
	endsGroup []string
}

func NewState(name string, transitions,startsGroups,endsGroup []string) *State {
	return &State{
		name : name,
		transitions: make([][]string,1024),
		startsGroups: make([]string,1024),
		endsGroup: make([]string,1024),
	}
}

func (s *State) addTransition(toState , matcher string){
	s.transitions = append(s.transitions, []string{matcher,toState})
}

func (s *State) unshiftTransition(toState, matcher string){
	s.transitions = append([]string{toState,matcher},s.transitions...)
}






func main() {
	nfa := EngineNFA()
	nfa.declareStates("q0","q1")
	nfa.setInitialState("q0")
	nfa.setEndingStates(["q1"])
	nfa.addTransition("q0","q1",*CharacterMatcher("a"))
	fmt.Printf(nfa.compute("a"))
}
