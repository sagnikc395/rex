class Matcher:
    # will return true if 'char' can be consumed by the matcher
    def matches(self,char):
        return False

    # need to know when it's an epsilon transition to avoid consuming inputs
    def is_epsilon():
        return None

    def label(self):
        return "undefined-matcher"

# now extends the other matchers on top of the base class

class CharacterMatcher(Matcher):
    def __init__(self,c):
        super().__init__()
        self.c = c

    def matches(self,char):
        return self.c == char

    def is_epsilon(self):
        return False

    def label(self):
        return self.c

class EpsilonMatcher(Matcher):
    # an epsilon matcher will always match
    def matches(self):
        return True

    def is_epsilon(self):
        return True

    def label(self):
        return "ε"
