# interface of the machine uses string to make it more readable,
#
class EngineNFA:
    def __init__(self):
        self.states = {}
        self.initial_state = None
        self.ending_states = None

    def set_initial_state(self,state):
        self.initial_state = state

    def set_ending_states(self,states):
        self.ending_states = states
