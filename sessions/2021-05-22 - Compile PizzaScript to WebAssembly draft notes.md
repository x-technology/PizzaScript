# 2021-05-22 - Compile PizzaScript to WebAssembly

# Changelog

## PizzaScript Variables

First, I wanted to review and refactor my parser implementation a little bit.

It misses the important extension mechanism. As a language compiler writer, I'd like to have a relatively easy way to add operators to my language. And what is more important, it is not only about mathematical operations, but also about all kind of PizzaScript language power.

Taht's right, it's time to introduce first "real" language feature to PizzaScript. Let's think of variables. The requirements we had declared before are:

- `Kotlin`-like variables operators `var, val`

```ps
export "pizza/io/print"

var a1573: string = "null"
a1573 = "1"
print a1573 // "1"

val b2217: int = 34
print b2217 // 34
```

- Variables are hoisted by default, unless it is disabled in the compiler's configuration

```ps
c6572 = 1
print c6572 // 1.0
c6572: float = 3
```

```json
{
  "Hoisting": "false"
}
```

Also, in case a type declaration is missing, interpreter needs to execute all expressions with such a variable in order to know variable's type. Value comes for free in such scenario. That is something called type inference. 

> [This is a real language type inference example in TypeScript](https://www.typescriptlang.org/docs/handbook/type-inference.html)

## Implementation

Wow, that's actually a lot already. Let's start with something small, like `var` variable type for instance. 

And actually, let's start with a bit of refactoring here

## Extending PizzaScript

An original implementation function for Pratt parsing algorithm looks like following

```go
func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}
```

```ps
// rbp - right binding power
fun expr(rbp = 0) {
  // null denotation
	var left = nud(next())
	while bp(peek()) > rbp
    // left denotation
    left = led(left, next())
  left
}
```

It misses one detail though - methods bound to operators recognition. However, we do support basic mathematical operations, we didn't generalize it yet. And it was an intention in the Pratt parse algorithm.

Now, in `PizzaScript`

```go
fun parse(precedence int) {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}
```