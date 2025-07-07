'''
a state if a set of transitions, and a transition is just a matcher and the
state where the transition goes
'''

class State:
    def __init__(self,name):
        self.name = name
        self.transitions = []
        self.starts_grp = []
        self.ends_grp = []

    def add_transition(self,to_state,matcher):
        self.transitions.append([matcher,to_state])

    # nfa's needs a way to define of choosing a path when there are multiple options
    # our machine will prioritize transations and give this meaning
    # add this transition and give it the highest priority
    def prepend_transition(self,to_state,matcher):
        self.transitions.insert([matcher,to_state],0)
