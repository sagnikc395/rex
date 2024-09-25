class State {
  constructor(name) {
    this.name = name;
    this.transitions = [];
    this.startsGroup = [];
    this.endsGroup = [];
  }

  addTransition(toState, matcher) {
    this.transitions.push([matcher, toState]);
  }

  unshiftTransition(toState, matcher) {
    this.transitions.unshift([matcher, toState]);
  }
}
