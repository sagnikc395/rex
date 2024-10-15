package main

import (
	"fmt"
	"strconv"
	"strings"
)

const repeatInfinity = -1

func parse(regex string) *parseContext {
	ctx := &parseContext{
		pos:    0,
		tokens: []token{},
	}
	for ctx.pos < len(regex) {
		process(regex, ctx)
		ctx.pos++
	}

	return ctx
}

//process each element

func process(regex string, ctx *parseContext) {
	ch := regex[ctx.pos]
	if ch == '(' {
		groupCtx := &parseContext{
			pos:    ctx.pos,
			tokens: []token{},
		}
		parseGroup(regex, groupCtx)
		ctx.tokens = append(ctx.tokens, token{
			tokenType: group,
			value:     groupCtx.tokens,
		})
	} else if ch == '[' { // <2>
		parseBracket(regex, ctx)
	} else if ch == '|' { // <3>
		parseOr(regex, ctx)
	} else if ch == '*' || ch == '?' || ch == '+' { // <4>
		parseRepeat(regex, ctx)
	} else if ch == '{' { // <5>
		parseRepeatSpecified(regex, ctx)
	} else { // <6>
		// literal
		t := token{
			tokenType: literal,
			value:     ch,
		}
		ctx.tokens = append(ctx.tokens, t)
	}
}

func parseGroup(regex string, ctx *parseContext) {
	ctx.pos++
	for regex[ctx.pos] != ')' {
		process(regex, ctx)
		ctx.pos++
	}
}

func parseBracket(regex string, ctx *parseContext) {
	ctx.pos++ // get past the LBRACKET
	var literals []string
	for regex[ctx.pos] != ']' { // <1>
		ch := regex[ctx.pos]

		if ch == '-' { // <2>
			next := regex[ctx.pos+1]                                    // <3-1>
			prev := literals[len(literals)-1][0]                        // <3-1>
			literals[len(literals)-1] = fmt.Sprintf("%c%c", prev, next) // <3-2>
			ctx.pos++                                                   // to consume the 'next' char
		} else { // <4>
			literals = append(literals, fmt.Sprintf("%c", ch))
		}

		ctx.pos++ // <5>
	}

	literalsSet := map[uint8]bool{}

	for _, l := range literals { // <6>
		for i := l[0]; i <= l[len(l)-1]; i++ { // <7>
			literalsSet[i] = true
		}
	}

	ctx.tokens = append(ctx.tokens, token{ // <8>
		tokenType: bracket,
		value:     literalsSet,
	})
}
func parseOr(regex string, ctx *parseContext) {
	rhsContext := &parseContext{
		pos:    ctx.pos,
		tokens: []token{},
	}
	rhsContext.pos += 1 // get past |
	for rhsContext.pos < len(regex) && regex[rhsContext.pos] != ')' {
		process(regex, rhsContext)
		rhsContext.pos += 1
	}
	// <1:end>

	// both sides of the OR expression
	left := token{
		tokenType: groupUncaptured,
		value:     ctx.tokens, // <2>
	}

	right := token{ // <3>
		tokenType: groupUncaptured,
		value:     rhsContext.tokens,
	}
	ctx.pos = rhsContext.pos // <4>

	ctx.tokens = []token{{ // <5>
		tokenType: or,
		value:     []token{left, right},
	}}
}

func parseRepeat(regex string, ctx *parseContext) {
	ch := regex[ctx.pos]
	var min, max int
	if ch == '*' {
		min = 0
		max = repeatInfinity
	} else if ch == '?' {
		min = 0
		max = 1
	} else {
		// ch == '+'
		min = 1
		max = repeatInfinity
	}
	// we need to wrap the last token with the quantifier data
	// so that we know what the min and max apply to
	lastToken := ctx.tokens[len(ctx.tokens)-1]
	ctx.tokens[len(ctx.tokens)-1] = token{
		tokenType: repeat,
		value: repeatPayload{
			min:   min,
			max:   max,
			token: lastToken,
		},
	}
}

func parseRepeatSpecified(regex string, ctx *parseContext) {
	// +1 because we skip LCURLY { at the beginning
	start := ctx.pos + 1
	// proceed until we reach to the end of the curly braces
	for regex[ctx.pos] != '}' {
		ctx.pos++
	}

	boundariesStr := regex[start:ctx.pos]       // <1>
	pieces := strings.Split(boundariesStr, ",") // <2>
	var min, max int
	if len(pieces) == 1 { // <3>
		if value, err := strconv.Atoi(pieces[0]); err != nil {
			panic(err.Error())
		} else {
			min = value
			max = value
		}
	} else if len(pieces) == 2 { // <4>
		if value, err := strconv.Atoi(pieces[0]); err != nil {
			panic(err.Error())
		} else {
			min = value
		}

		if pieces[1] == "" {
			max = repeatInfinity
		} else if value, err := strconv.Atoi(pieces[1]); err != nil {
			panic(err.Error())
		} else {
			max = value
		}
	} else {
		panic(fmt.Sprintf("There must be either 1 or 2 values specified for the quantifier: provided '%s'", boundariesStr))
	}

	// we need to wrap the last token with the quantifier data
	// so that we know what the min and max apply to
	// <5>
	lastToken := ctx.tokens[len(ctx.tokens)-1]
	ctx.tokens[len(ctx.tokens)-1] = token{
		tokenType: repeat,
		value: repeatPayload{
			min:   min,
			max:   max,
			token: lastToken,
		},
	}
}
