import type { NullLiteral } from "typescript";
import { State } from "./state";

export class EngineNFA {
  states: Object;
  initialState: null;
  endingStates: null;

  constructor() {
    this.states = {};
    this.initialState = null;
    this.endingStates = null;
  }

  setInitialState(state: null) {
    this.initialState = state;
  }

  setEndingStates(states: null) {
    this.endingStates = states;
  }

  addState(name: string | number) {
    this.states[name] = new State(name);
  }

  declareStates(...names: any[]) {
    for (const n of names) {
      this.addState(n);
    }
  }

  addTransition(
    fromState: string | number,
    toState: string | number,
    matcher: any
  ) {
    this.states[fromState].addTransition(this.states[toState], matcher);
  }

  unshiftTransition(fromState, toState, matcher) {
    this.states[fromState].unshiftTransition(this.states[toState], matcher);
  }

  compute(string) {
    //TODO
  }
}
