# `PizzaScript`

![pizzascript](/assets/pizzascript.jpg)

Writing PizzaScript Lexer & Parser with `Golang` and `RxGo`

- PizzaScript
  - Goals
  - Examples
- Theory
  - Programming Languages
  - Lexer
  - Parser
    - Pratt Parsing Algorithm
- Code Overview
- Summary
  - Next Steps

### Goals

- Learn `Go` language, and key libraries like `RxGo`
- Understand how programming languages & interpreters work
- Experiment with `WebAssembly`

### PizzaScript Key Features Plans

- Cool name!
- `JavaScript` like syntax
- Variables, integers, floats and strings
- Arithmetic expressions, built-in functions
- Functions, first class functions, closures
- Dynamic types and coercions
- Modules & Standard library
- Compiles to `WebAssembly` and produces modules in `wasm` or `wat` formats

## Examples

The `PizzaScript` goal is to take everyone equally

```ps
5 + 5 * 10 // = 100
```

All operators should have same priority

```ps
var h1679 = "1"
val g2788 = 2

h1679 + g2788 == "12"
g2788 + h1679 == 3
```

In mathematics, commutative property means

> "changing the order of the operands does not change the result"

```ps
fun sum(var a1573: string, b7232): int {
  a1573 + b7232
}

sum(1, 2) // 12
```

Standard library, ported to `WebAssembly`

```ps
c6572 = 1
print c6572 // 1.0
c6572: float = 3
```

All variables are hoisted by default, and it is possible to switch it off

```json
{
  "Hoisting": "false"
}
```

# Programming Languages

- `Alphabet` all available symbols

- `Chains` of characters makes a text in `alphabet`

- A specific set of `chains` from the `alphabet` form a `language`

## How can we specify the language?

- `Grammars` are rules defining a language, a written in a special form algorithm.

The most known is [the `Backus–Naur` form](https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_form)

![](../assets/bnf.png)


[An example of a calculator language](https://bnfplayground.pauliankline.com/?bnf=%3Cexpr%3E%20%3A%3A%3D%20%3Coperand%3E%20((%22%2B%22%20%7C%20%22-%22%20%7C%20%22*%22%20%7C%20%22%2F%22)%20%3Coperand%3E)*%0A%3Coperand%3E%20%3A%3A%3D%20%3Cnum%3E%20%7C%20(%22%2B%22%20%7C%20%22-%22)%20%3Coperand%3E%0A%3Cnum%3E%20%3A%3A%3D%20%3Cdigit%3E%2B%20%0A%3Cdigit%3E%20%3A%3A%3D%20%220%22%20%7C%20%221%22%20%7C%20%222%22%20%7C%20%223%22%20%7C%20%224%22%20%7C%20%225%22%20%7C%20%226%22%20%7C%20%227%22%20%7C%20%228%22%20%7C%20%229%22&name=PizzaScript%200.0.3).

```
<expr> ::= <operand> (("+" | "-" | "*" | "/") <operand>)*
<operand> ::= <num> | ("+" | "-") <operand>
<num> ::= <digit>+ 
<digit> ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```

## What is described with the grammar above?

```
G = ({S},{0, 1}, {S → 0S1, S → 01},S). L(G) ?
G = ({S, A},{0, 1}, {S → 0S, S → 0A, A → 1A, A → 1},S). L(G) ?
```

Or, a better way to define a language - write a compiler!

## Stages of a Compiler

- `Lexer` - split text into lexems or tokens

Tokens are of different types, such as data types - **numbers, strings**, variables definitions - **identifiers**, operators - **+, -, =, /**, etc.

Tokens for `a + b(c)`

```
NAME "a"
PLUS "+"
NAME "b"
LEFT_PAREN "("
NAME "c"
RIGHT_PAREN ")"
```

- `Parser` - works with tokens, checks language syntax and produces an *Abstract Syntax Tree* (ast).

![](../assets/ast.png)

`AST` - a data object, representing programm's text as a declarative tree structure. 

Check the [AST Explorer](https://astexplorer.net/)

![](../assets/compiler-flow.png)

- [Parsing Algorithms](https://www.tutorialspoint.com/compiler_design/compiler_design_bottom_up_parser.htm)

![top down parsing](/assets/top-down-parsing.jpg)

BNF can be converted into a program by applying the following steps:
- the left side (non-terminal) of a BNF rule is a function declaration 
- the right side is the function body (checking for terminal ones)
- quantifiers `(*, +)` are translated into like `while, for, if` and so on

- **Recursive descent top down** technique, [wikipedia shows a good example](https://en.wikipedia.org/wiki/Recursive_descent_parser)

![bottom up parsing](/assets/bottom-up-parsing.jpg)

> In computer science, an **operator precedence parser** is a **bottom-up parser** that interprets an operator-precedence grammar

> Another **precedence parser** known as Pratt parsing was first described by Vaughn Pratt in the 1973 paper **Top down operator precedence**, based on **recursive descent**. Though it predates precedence climbing, it can be viewed as a generalization of **precedence climbing**

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

**nud** (or prefix) - `null denotation`, no left-context `+1, -1` or even `-+-1`

**led** (or infix) - `left denotation` - `1+2`

**bp** - binding power `bp(*) > bp(+)`

```ps
fun nud(operator) {
  node(operator)
}

fun led(left, operator) {
  node(left, operator, expr(bp(operator))
}

fun bp(operator) {
  precendence[operator]
}
```

- `Compile`, `Eval`, or `Interpret`. `PizzaScript` compiler will compile to `WebAssembly`, meaning we will produce an output in `wasm` or `wat` format.

> `Compiler` - a program that answers the question - if text is part of a language

## Code Review

- `Golang`
- [`RxGo`](https://github.com/ReactiveX/RxGo) implementation of `ReactiveX` - a common pattern for asyncronous data flow, including collections, iterators, operators
  - Unify API, data flow & interaction
  - [How to choose ReactiveX operator?](http://xgrommx.github.io/rx-book/content/which_operator_do_i_use/instance_operators.html)

- Demo
- Stages
  - [Lexer](https://korzio.medium.com/writing-pizzascript-lexer-with-rxgo-a-saga-in-iii-slices-3790dc6099e7)
    - Filter meaningless symbols
    - Reduce & Map combine symbols into tokens
  - Parser

- Operators overview (filter, map, create, reduce, ...)
  - [How to choose operator?](http://xgrommx.github.io/rx-book/content/which_operator_do_i_use/instance_operators.html)

- How to filter with order?

```go
DistinctUntilChanged(func(_ context.Context, i interface{}) (interface{}, error) {
  var str = i.(string)
  ch := []byte(str)[0]
  log.Info(str)

  return isNumber(ch), nil
})
```

```bash
>> 11-00
distinct:  1
distinct:  11
distinct:  11-
map:  11-
map:  11-
distinct:  11-0
map:  11-0
{INT 11-}
distinct:  11-00
{INT 11-}
{INT 11-0}
>> 11-00
distinct:  1
distinct:  11
distinct:  11-
map:  11-
map:  11-
{INT 11-}
{INT 11-}
distinct:  11-0
distinct:  11-00
map:  11-00
{INT 11-00}
```

- [The Pyramid of Doom](https://korzio.medium.com/pizzascript-parser-with-rxgo-the-pyramid-of-doom-36e574f129dc)

```go
tree := p.lexer.Tokens().
  Scan(func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
    it, isIterator := acc.(iterator)
    next := elem.(token.Token)
    if !isIterator || it.left == nil {
      // null denotation
      return nud(acc, next), nil
    }
    if it.operator != nil {
      var prevIt *iterator
      if len(it.stack) > 0 {
        prevIt = &it.stack[len(it.stack)-1]
      }
      // binding power comparison
      for len(it.stack) > 0 && prevIt != nil && bp(it.operator) < bp(prevIt.operator) {
        prevIt = &it.stack[len(it.stack)-1]
        // left denotation
        it.left = led(*prevIt.left, *prevIt.operator, *it.left)
        // in golang slices are immutable
        it.stack = it.stack[:len(it.stack)-1]
      }
      var newIt iterator
      newIt.stack = append(it.stack, it)
      // next iteration
      return nud(newIt, next), nil
    } 
    it.operator = &next
    return it, nil
  })
```

![pizzascript](/assets/pyramid-of-pizza.png)

- Homogeneous AST
- Scan operator
- Github Actions

# Feedback

- Where do we go from here?
  - Extend with variables
  - Port to WebAssembly
- Is RxGo a good choice?
- Meetups...

## Links

- [Pratt Parsing by Jonathan Apodaca](https://dev.to/jrop/pratt-parsing)
- [Introduction to the PizzaScript project](https://korzio.medium.com/introducing-pizzascript-educational-go-open-source-project-d7a15128db94)
- [PizzaScript Lexer with RxGo](https://korzio.medium.com/writing-pizzascript-lexer-with-rxgo-a-saga-in-iii-slices-3790dc6099e7)
- [PizzaScript Parser with RxGo - The Pyramid of Doom](https://korzio.medium.com/pizzascript-parser-with-rxgo-the-pyramid-of-doom-36e574f129dc)
- [PizzaScript Eventbrite #2 - Parser with RxGo](https://www.eventbrite.co.uk/e/pizzascript-2-parser-with-rxgo-tickets-145559009917)
