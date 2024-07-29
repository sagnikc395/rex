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

type Matcher interface {
	Matches(char rune) bool
	IsEpsilon() bool
	Label() string
}

//characterMatcher class
type CharacterMatcher struct {
	c rune
}
//constructor for CharacterMatcher
func NewCharacterMatcher(c rune) *CharacterMatcher{
	return &CharacterMatcher{c:c}
}

func (cm *CharacterMatcher) Matches(char rune)bool {
	return cm.c == char
}

func (cm *CharacterMatcher) IsEpsilon() bool {
    return false
}

func (cm *CharacterMatcher) Label() string {
    return string(cm.c)
}


//epsilonmatcher class
type EpsilonMatcher struct {}

func NewEpsilonMatcher() *EpsilonMatcher {
    return &EpsilonMatcher{}
}

// An epsilon matcher always matches
func (em *EpsilonMatcher) Matches(char rune) bool {
    return true
}

func (em *EpsilonMatcher) IsEpsilon() bool {
    return true
}

func (em *EpsilonMatcher) Label() string {
    return "ε"
}



func main() {
	nfa := EngineNFA()
	nfa.declareStates("q0","q1")
	nfa.setInitialState("q0")
	nfa.setEndingStates(["q1"])
	nfa.addTransition("q0","q1",*CharacterMatcher("a"))
	fmt.Printf(nfa.compute("a"))
}
