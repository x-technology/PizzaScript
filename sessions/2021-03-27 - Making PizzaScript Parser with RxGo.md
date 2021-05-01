# Writing `PizzaScript` Parser with `Golang` and `RxGo`

![pizzascript](/assets/pizzascript.jpg)

- About us
- PizzaScript
- Theory
  - Programming Languages & Compilers
  - Parser
- Summary
  - Code Review
  - Feedback

# About us

**XTechnology** - *Educational Programming and Data Science* open source project

## Goals

- learn new technologies
- share ideas and knowledge
- build online projects and community
- and have fun! 👋

Copyright (c) 2021 x-technology

# PizzaScript

- Learn `Go` language, and key libraries like `RxGo`
- Understand how programming languages & interpreters work
- Experiment with `WebAssembly`

## Features & Examples

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

- Standard `pizza` 🍕library ported to `WebAssembly`

```ps
export "pizza/io/print"

fun factorial(n: int): int {
  // curly braces can be omitted if context has unique meaning
  if n = 0
    1
  else
    n * factorial n - 1
  // there is no need for return statement inside function body
}

for var n := 0; n <= 16; n++ {
  print n, '! = ', factorial n
  // also possible to specify parentheses
  // print(n, '! = ', factorial n)
}
```

# Questions

- What's the difference between an interpreter and a compiler?
- What is a token?
- How can we specify a language?
- What is an AST? 
- What is the difference between a parser and a lexer?

# Programming Languages & Compilers

- `Alphabet` all available symbols

- `Chains` of characters makes a text in `alphabet`

- A specific set of `chains` from the `alphabet` form a `language`

- `Grammars` are rules defining a language, a written in a special form algorithm.

The most known is [the `Backus–Naur` form](https://en.wikipedia.org/wiki/Backus%E2%80%93Naur_form)

[An example of a calculator language](https://bnfplayground.pauliankline.com/?bnf=%3Cexpr%3E%20%3A%3A%3D%20%3Coperand%3E%20((%22%2B%22%20%7C%20%22-%22%20%7C%20%22*%22%20%7C%20%22%2F%22)%20%3Coperand%3E)*%0A%3Coperand%3E%20%3A%3A%3D%20(%22%2B%22%20%7C%20%22-%22)*%20%3Cnum%3E%0A%3Cnum%3E%20%3A%3A%3D%20%3Cdigit%3E%2B%20%0A%3Cdigit%3E%20%3A%3A%3D%20%220%22%20%7C%20%221%22%20%7C%20%222%22%20%7C%20%223%22%20%7C%20%224%22%20%7C%20%225%22%20%7C%20%226%22%20%7C%20%227%22%20%7C%20%228%22%20%7C%20%229%22%0A&name=PizzaScript%200.0.3).

```
<expr> ::= <operand> (("+" | "-" | "*" | "/") <operand>)*
<operand> ::= ("+" | "-")* <num>
<num> ::= <digit>+ 
<digit> ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```

Or, a better way to define a language - write a compiler!

- `Lexer` - split text into lexems or tokens

```go
type TokenType string
type Token struct {
  Type    TokenType
  Literal string
}
```

- `Parser` - program syntax analysis

- `Compile`, `Eval`, or `Interpret`. `PizzaScript` compiles to `WebAssembly`, and we will produce an output in `wasm` or `wat` format.

> `Compiler` - a program that also answers the question - if text is part of a language

## Implementation

- [`RxGo`](https://github.com/ReactiveX/RxGo) implementation of `ReactiveX` - a common pattern for asyncronous data flow, including collections, iterators, operators
  - Unify API, data flow & interaction

# Parser

- `Parser` - works with tokens, checks language syntax and produces an *Abstract Syntax Tree* (ast).

![](../assets/ast.png)

- `AST` - a data object, representing programm's text as a declarative tree structure. 

Check the [AST Explorer](https://astexplorer.net/)

![compiler flow](../assets/compiler-flow.png)

- [Parsing Algorithms](https://www.tutorialspoint.com/compiler_design/compiler_design_bottom_up_parser.htm)

![top down parsing](/assets/top-down-parsing.jpg)

BNF can be converted into a program by applying the following steps:
- the left side (non-terminal) of a BNF rule is a function declaration 
- the right side is the function body (checking for terminal ones)
- quantifiers `(*, +)` are translated into like `while, for, if` and so on

- **Recursive descent top down** technique, [wikipedia shows a good example](https://en.wikipedia.org/wiki/Recursive_descent_parser)

![bottom up parsing](/assets/bottom-up-parsing.jpg)

### Definition Quirks

> In computer science, an **operator precedence parser** is a **bottom-up parser** that interprets an operator-precedence grammar

> Another **precedence parser** known as Pratt parsing was first described by Vaughn Pratt in the 1973 paper **Top down operator precedence**, based on **recursive descent**. Though it predates precedence climbing, it can be viewed as a generalization of **precedence climbing**

## Algorithm

- 1st Step

```ps
fun expr {
  let left = nud(next())
  while (!eof)
    left = led(left, next())
  return left
}
```

- Precedence though

The canonical example `1 + 2 * 3` should result in a tree like 

```bash
  +
 / \
1   *
   / \
  2   3

// not

    *
   / \
  +   3
 / \
1   2
```

> In fact, there is a `PizzaScript` proposal to take all operators equally 🤒 😸

```ps
5 + 5 * 10 // = 100
```

- 2nd Step

```
          ┌────┐
          │ q0 │
          └─╥──┘
            ║  c ← nud;
            ║  advance;
            ║  left ← run c
         ┌──╨─┐
    ╔═▶▶═╡ q1 │
    ║    └──╥─┘
    ║       ║  rbp < lpb/
    ║       ║  c ← led;
    ║       ║  advance;
    ║       ║  left ← run c
    ╚═══◀◀══╝
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

// Pratt Parsing - Jonathan Apodaca
// https://dev.to/jrop/pratt-parsing
```

- **nud** (or prefix) - `null denotation`, no left-context `+1, -1` or even `-+-1`

- **led** (or infix) - `left denotation` - `1+2`

- **bp** - binding power `bp(*) > bp(+)`

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

## [The Pyramid of Doom](https://korzio.medium.com/pizzascript-parser-with-rxgo-the-pyramid-of-doom-36e574f129dc)

![pizzascript](/assets/pyramid-of-pizza.png)

# Summary

## Code Review

- Homogeneous AST
- Scan operator
- Golang `any` type, really?

```go
func nud(acc interface{}, next token.Token) interface{} {  
  // wtf return different types, really????
}
```

- [Code Review](https://github.com/x-technology/PizzaScript/compare/b4791695cbf6ed33c9937fafc2c23224373d3a47..main)
- [Github Actions](https://github.com/x-technology/PizzaScript/actions/runs/651968919)
- [Changelog](/CHANGELOG.md)

## Feedback

- Where do we go from here?
  - Extend with variables
  - Port to WebAssembly
- [What's our next steps?](https://forms.gle/nWSJnX6uH8rLk4iP7)
- Is `RxGo` a good choice?
  - Not sure 🤷‍♀️ probably not...

## Links

- [Pratt Parsing by Jonathan Apodaca](https://dev.to/jrop/pratt-parsing)
- [PizzaScript Lexer with RxGo](https://korzio.medium.com/writing-pizzascript-lexer-with-rxgo-a-saga-in-iii-slices-3790dc6099e7)
- [PizzaScript Parser with RxGo - The Pyramid of Doom](https://korzio.medium.com/pizzascript-parser-with-rxgo-the-pyramid-of-doom-36e574f129dc)
- [How to choose operator?](http://xgrommx.github.io/rx-book/content/which_operator_do_i_use/instance_operators.html)