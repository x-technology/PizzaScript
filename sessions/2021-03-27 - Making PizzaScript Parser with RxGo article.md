# Making PizzaScript Parser with RxGo

Hello there 👋! In the previous chapter, [we introduced a toy programming language PizzaScript 🍕](https://korzio.medium.com/writing-pizzascript-lexer-with-rxgo-a-saga-in-iii-slices-3790dc6099e7) to start our journey into Lexers, Parsers, `Reactive` patterns and many other interesting topics. 

In this series of meetups and articles we learn Go language, including key libraries like [RxGo](https://github.com/ReactiveX/RxGo), explore how programming languages and interpreters work, and experiment with WebAssembly.

![](../assets/pizzascript.jpg)

All articles in this series

- [Introduction to the PizzaScript project](https://korzio.medium.com/introducing-pizzascript-educational-go-open-source-project-d7a15128db94)
- [Writing PizzaScript Lexer with RxGo](https://korzio.medium.com/writing-pizzascript-lexer-with-rxgo-a-saga-in-iii-slices-3790dc6099e7)
- [🍕PizzaScript Parser with RxGo (**this article**)](TODO)
- 🍕Slice 4 - WebAssembly as a compilation target
- 🍕Slice 5 - Transforming AST Tree into WebAssembly

Today, we are going to focus on programming language parsers, overview and practice with a calculator use case. We implement *Pratt's top down operator precedence* algorithm in Go using RxGo library.

---

We already discussed the nature of programming languages and compilers. 
Lexer is the first step of analyzing a program. It splits original text into simple structures, aka lexems or tokens. We are familiar with PizzaScript tokens

```go
// https://gist.github.com/korzio/76492131637d5fb12f2b92cf26009bb2#file-token-go
package token

type TokenType string
type Token struct {
  Type TokenType
  Literal string
}
```

During the syntax analysis or parser step we check if a program is valid against defined language syntax. 

Below you can see [an example of a calculator language](https://bnfplayground.pauliankline.com/?bnf=%3Cexpr%3E%20%3A%3A%3D%20%3Coperand%3E%20((%22%2B%22%20%7C%20%22-%22%20%7C%20%22*%22%20%7C%20%22%2F%22)%20%3Coperand%3E)*%0A%3Coperand%3E%20%3A%3A%3D%20%3Cnum%3E%20%7C%20(%22%2B%22%20%7C%20%22-%22)%20%3Coperand%3E%0A%3Cnum%3E%20%3A%3A%3D%20%3Cdigit%3E%2B%20%0A%3Cdigit%3E%20%3A%3A%3D%20%220%22%20%7C%20%221%22%20%7C%20%222%22%20%7C%20%223%22%20%7C%20%224%22%20%7C%20%225%22%20%7C%20%226%22%20%7C%20%227%22%20%7C%20%228%22%20%7C%20%229%22&name=PizzaScript%200.0.3) written in [the Backus–Naur form](https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_form).

```
<expr> ::= <operand> (("+" | "-" | "*" | "/") <operand>)*
<operand> ::= <num> | ("+" | "-") <operand>
<num> ::= <digit>+ 
<digit> ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```

An expression can contain one or multiple operands combined with arithmetic operations. Each operand is a number or a set of digits, possibly preceded by a "+" or "-" signs. 

> Hello 👋 does PizzaScript support it yet? 📢 It does!

```bash
➜ pizzascript git:(main) ✗ go run main.go
Hello alex! This is the PizzaScript programming language!
Feel free to type in commands

>> +1-2/2
0

>> +++2
2

>> ---2
-2
```

## The Algorithm

> Now, how are we going to transform lexems into an the AST?

Good question! A nice thing about BNF grammars (language representations), as [we mentioned last time](https://korzio.medium.com/writing-pizzascript-lexer-with-rxgo-a-saga-in-iii-slices-3790dc6099e7) is that rules can be transformed into a parsing algorithm. Especially when this notation is following a special "context-free" principle. It is a kind of grammar's limitation in which the left part of a rule can contain only non-terminal tokens, the right part can contain both non-terminal and terminal ones. With that in mind, a BNF can be converted into a working program by applying the following steps
- the left side of a BNF rule is going to be a function declaration, defined for each non-terminal token, 
- the right side is the function body, consequentially calling functions and checking for terminal ones,
- quantifiers `(*, +)` and other syntax is translated into control flow constructions like `while, for, if` and so on.

```ps
// <expr> ::= <operand> (("+" | "-" | "*" | "/") <operand>)*
fun expr {
  // remember, in pizzascript parenthesis can be omitted in most cases
  operand

  while operator
    operand
}

fun operator {
  // take next token from input
  next == "+" || next == "-" || next == "*" || next == "/"
}

fun next {
  input[i + 1]
}
```

This way of constructing a parser is known a recursive descent top down technique, [wikipedia shows a good example](https://en.wikipedia.org/wiki/Recursive_descent_parser) of it.

> The alternative are bottom-up parsers, which are out of scope for this article. 😄 Luckily, we skip language and parser categories, too 😅

It's straightforward, clean and nice. It has a disadvantage though - you have to implement each combination of operators, operands and expressions separately. As you can imagine, as `pizzascript` grammar grows, that would be a hell amount of work to maintain. There is where Pratt parser comes into place. It encapsulates the parser's logic into a simple shape.

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

**nud** (or prefix) stands for `null denotation`, and can operate to the right with no left-context (for expressions like `+1, -1` or even `-+-1`) 

**led** (or infix) is `left denotation` and operate on two operands from left to right (for normal expressions like `1+2`)

> If you wonder 😸 expressions produce values, statements don't.

```ps
fun nud(operator) {
  node(operator)
}

fun led(left, operator) {
  node(left, operator, expr(bp(operator))
}
```

The one not explained yet function is `bp` - binding power. It is a key for Pratt algorithm as it solves another important problem of which token has more priority. Say, `*` has more priority than `+`. 

> There is a proposal in `PizzaScript` to treat every expression with the same priority 😀

```ps
fun bp(operator) {
  precendence[operator]
}
```

> We recommend to go through [Pratt Parsing by Jonathan Apodaca](https://dev.to/jrop/pratt-parsing) - a concise, practical, and very informative article about it.

## Node type

Next, no matter how we proceed with our interpreter, we will have a intermediate representation (IR is a commong term for language parsing programs). Parser takes input stream of tokens and build an hieararchical structure representing the original program - the abstract syntax tree (AST). 

> There is also a general notion of a parse tree. Let's take it as a tree of algorithm decisions made to parse an input. Whereas AST is a final representation of a program. It's really complex matter we are going through and we need to carefully stop before falling into the abyss. 

The `AST` format is non standard across different interpreters, and generally nodes inside the tree can be of different types. We are going to use a *Homogeneous AST* model in which all the nodes are of a same *Node* type. That would help us to simplify implementation for now, and we can refactor and extend it later. 

```go
type Node struct {
	Token *token.Token
	Left  *Node
	Right *Node
}

func (n *Node) ToString() string {
	res := "{" + n.Token.Literal

	if n.Left != nil {
		res += "," + n.Left.ToString()
	}

	if n.Right != nil {
		res += "," + n.Right.ToString()
	}

	res += "}"
	return res
}
```

As one can see, *Node* is a potentially recursive structure. Although we stated it as homogeneous, its' interpretation or even evaluation still would differ depending on nodes and tokens it contains.

> And we will explore it in future chapters 🤫

## Implementation & Problems

---

Usually, compilers act more in an "imperative" way, by controlling processed text and position. With `RxGo` and `Observable` pattern this concept changes from top to bottom. And here is the challenge. Now, a stream of asynchronous events is the compiler's input and handlers have to deal with it. We saw how it happened with lexer before. Instead of saving and incrementing the current position while reading text, we operate on given text chunks. We had to save the aggregated object. todo code.

    We need to support:
    - operator precedence

    > operator precedence. An alternative term for this is order of operations, which should make clearer what operator precedence describes: which priority do different operators have. The canonical example is this one, which we saw earlier
    > 5 + 5 * 10

    - infix and prefix expressions

    > Pratt parsing works by scanning the input tokens, and classifying them into two categories:
    - no left-context (null denotation / prefix)
    - left-to-right (infix)

    compare with wikipedia c implemenation 

    write in a pseudocode

    ```go
    // 1 + 2 * 3
    // -> {1, +, {2, *, 3}}
    // 
    // 1 * 2 + 3
    // -> {{1, *, 2}, + , 3}
    func (p *Parser) parseExpression(precedence int) ast.Expression {
      // first precedence LOWEST
      // second precedence SUM
      prefix := p.prefixParseFns[p.curToken.Type]
      // first parseIntegerLiteral 1
      // second parseIntegerLiteral 2
      if prefix == nil {
        p.noPrefixParseFnError(p.curToken.Type)
        return nil
      }
      leftExp := prefix()

      // second precedence SUM, SUM > LOWEST
      for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
        infix := p.infixParseFns[p.peekToken.Type]
        // parseInfixExpression +
        if infix == nil {
          return leftExp
        }

        p.nextToken()

        leftExp = infix(leftExp)
      }

      return leftExp
    }
    ```

    so it's a set of lookahead functions,

    ```go
    func (p *Parser) parsePrefixExpression() ast.Expression {
      expression := &ast.PrefixExpression{
        Token:    p.curToken,
        Operator: p.curToken.Literal,
      }

      p.nextToken()

      expression.Right = p.parseExpression(PREFIX)

      return expression
    }
    ```

    we need to find appropriate reactive implementations for them. probably there will be same scan with additional structures. And actually it looks like a good idea to addd intermediate observables inside temporary tokens

    the task requires recursion
    actually, let's do the recursion with state
    we'll push every valid item, and pop when it is over

    write in pseudocode

    CALCULATOR

    first step

```js
function expr() {
    let left = nud(next())
    while (!eof)
        left = led(left, next())
    return left
}

// LED(left, operator) = Tree(left, operator, right=next())
```

```
// todo change to expr(bp(operator))
right := nud(next)
it.left = led(it.left, *it.operator, right)
```

that just eats everything from left to right

second step

```js
function expr(rbp = 0) {
	let left = nud(next())
	while (bp(peek()) > rbp)
    left = led(left, next())
    // LED(left, operator) = Tree(left, operator, right=expr(bp(operator)))
	return left
}

```

current implementation

```go
accToken, isAccToken := acc.(interToken)
currentToken := elem.(token.Token)

// if acc is not defined, save current state and continue
if !isAccToken {
  accNode := &ast.Node{Token: currentToken}
  accToken.states = append(accToken.states, accNode)

  return accToken, nil
}

// if acc is prefixExpression of a first item in state
currentState := accToken.states[len(accToken.states)-1]
accToken.states = accToken.states[:len(accToken.states)-1]

prefix := p.prefixParseFns[currentState.Token.Type]
if prefix == nil {
  log.Err("No prefix parse function found for token", currentState.Token.Type)
  return nil, nil
}

newState := prefix(currentState)
newState.Token = currentToken
// save states back
accToken.states = append(accToken.states, newState)

return accToken, nil

// func (p *Parser) parseExpression(precedence int) ast.Node {
// 	prefix := p.prefixParseFns[p.curToken.Type]
// 	if prefix == nil {
// 		p.noPrefixParseFnError(p.curToken.Type)
// 		return nil
// 	}
// 	leftExp := prefix()
// 	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
// 		infix := p.infixParseFns[p.peekToken.Type]
// 		if infix == nil {
// 			return leftExp
// 		}

// 		p.nextToken()

// 		leftExp = infix(leftExp)
// 	}

// 	return leftExp
// }

// .Map(func(_ context.Context, i interface{}) (interface{}, error) {
// 	log.Debug("parser", i)
// 	tok := i.(token.Token)
// 	prefix := p.prefixParseFns[tok.Type]
// 	if prefix == nil {
// 		log.Info("Could not find prefix for token", tok)
// 		return nil, nil
// 	}
// 	leftExp := prefix()
// 	return leftExp, nil
// })
```

- ReactiveX, rxgo
  - Benefits of using streams
    - standard operators, 
    - iterators, 
    - flows,
    - asyncronous,
    - immutable
  - Operators Overview - filter, map, create, scan

  > Does Scala have ReactiveX analogue?

- Code Review (20 minutes)
  - Oh, btw logger
  - Explain language techniques

- CHANGELOG

# TODO

- http://rigaux.org/language-study/syntax-across-languages/Mthmt.html

- change all lists to reactive collections?

- why to use observables?

- check syntax

```go
// type conversion
i.V.(string)
```

- check concept channels, goroutine

# Feedback

- Where do we go from here?
  - WebAssembly - try a compile target, explain WebAssembly basics
  - Programming languages - extending PizzaScript and the interpreter
  - More on Golang? 
  - More on ReactiveX operators?

# Summary

- Use appropiate tools
- First implement draft algorithm, than polish

- Learn lots of Go specifics
- Learn Rx operators
- Program

## Links

- [How to choose operator?](http://xgrommx.github.io/rx-book/content/which_operator_do_i_use/instance_operators.html)

TODO hand-written vs automatic parsers

TODO check *ast.Node - * vs &

TODO why this comment? // TODO change to reduce

- [ ] Complete https://ruslanspivak.com/lsbasi-part3/

TODO twit about all links