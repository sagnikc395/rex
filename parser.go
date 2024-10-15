package main

func parse(regex string) *parseContext {
	ctx := &parseContext{
		pos:    0,
		tokens: []token{},
	}

	for ctx.pos < len(regex) {
		processElement(regex, ctx)
		ctx.pos++
	}

	return ctx
}

//process each element

func processElement(regex string, ctx *parseContext) {
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
