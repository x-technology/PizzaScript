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

> !Be careful, we use `PizzaScript` code to explain the problem and show `Go` implementation in the end!

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

Good question! A nice thing about BNF grammars (language representations), as [we mentioned last time](https://korzio.medium.com/writing-pizzascript-lexer-with-rxgo-a-saga-in-iii-slices-3790dc6099e7) is that rules can be transformed into a parsing algorithm. Especially when this notation is following a special "context-free" principle. It is a kind of grammar's limitation in which the left part of a rule can contain only non-terminal tokens, the right part can contain both non-terminal and terminal ones. 

> Non-terminal basically means they still can be transformed to something, while terminals are final states.

With that in mind, a BNF can be converted into a working program by applying the following steps
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

It's straightforward, clean and nice. It has a disadvantage though - you have to implement each combination of operators, operands and expressions separately. As you can imagine, as `pizzascript` grammar grows, that would be a hell amount of work to maintain. There is where [**Vaughan Pratt** parsing algorithm](https://tdop.github.io/) (so-called *Top Down Operator Precedence* algorithm) comes into place. It encapsulates the parser's logic into a simple shape.

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

The canonical example `1 + 2 * 3` should result in a tree like 

```bash
  +
 / \
1   *
   / \
  2   3

// and not

    *
   / \
  +   3
 / \
1   2
```

> In fact, there is a proposal in `PizzaScript` to treat every expression with the same priority 😀(1+2*3 == 9)

```ps
fun bp(operator) {
  precendence[operator]
}
```

![operator precedence levels](/assets/tdop-expr-bp.png)

> We recommend to go through [Pratt Parsing by Jonathan Apodaca](https://dev.to/jrop/pratt-parsing) - a concise, practical, and very informative article about it.

Let's recap the Pratt algorithm:
- each operator has a precedence, or binding power
- tokens are recognised differently in null or left position (nud or led)
- recursive parse expressions function consumes tokens from left to right, until it reaches an operator of precedence less than or equal to the previous operator

## Implementation details & Problems

Let's start with simple ones. Binding power definition and calculation in *Go* looks like

```go
// iota enumerator gives an effective way in go language
// to use constants in enum-like constructs
const (
	none = iota
	plus
	mul
)

// enum-like map of operator token keys and iota values defined above
var precendences = map[string]int{
	token.PLUS:     plus,
	token.ASTERISK: mul,
}

// calculate binding power
func bp(tok *token.Token) int {
	if precendence, ok := precendences[string(tok.Type)]; ok {
		return precendence
	}

	return none
}
```

The `led` function does nothing more

```go
func led(left ast.Node, operator token.Token, right ast.Node) *ast.Node {
	return &ast.Node{Left: &left, Token: &operator, Right: &right}
}
```

We will left the `nud` implementation for later as it is has more sense with the main parse expression finctionality.

## Node type

No matter how we proceed with our compiler further, we will have an intermediate representation (*IR*) of a processed text. Parser takes an input stream of tokens and builds an hieararchical data structure mirroring the original program - the abstract syntax tree (AST). 

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

## The main parse expression

And here is the challenge. 

Usually, compilers act more in an "imperative" way, by controlling processed text and position. With `RxGo` and `Observable` pattern this concept changes from top to bottom. Now, a stream of asynchronous events is the compiler's input and handlers have to deal with it. We saw how it happened with the lexer before. Instead of saving and incrementing the current position while reading text, we operate on asyncronously given text chunks. 

Last time, we used [the `RxGo Scan` operator](https://github.com/ReactiveX/RxGo/blob/master/doc/scan.md)

![rxgo scan](../assets/rxgo-scan.png)

It allows us to use the `look-ahead` technique - to keep previous iteration value and produce a required one based on the previos value. Again, we introduce an intermediate iterator type to make a decision on produced values.

To implement the recursive nature of the original algorithm, we had to use one another trick and save the aggregated stack. We call it 

> The pyramid of doom

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
        it.stack = it.stack[:len(it.stack)-1]
      }

      var newIt iterator
      newIt.stack = append(it.stack, it)

      return nud(newIt, next), nil
    } 

    it.operator = &next
    return it, nil
  })
```

Instead of imperatively taking the next token from stream, it saves current iterator state. For recursion analogue it uses an iterator's stack.

In the future we want to refactor this function, and use [the `Reduce` operator](https://github.com/ReactiveX/RxGo/blob/master/doc/reduce.md) instead.

![rxgo scan](/assets/rxgo-reduce.png)

And the last one - the promised `nud` function

```go
func nud(acc interface{}, next token.Token) iterator {
	// Node{Token: token.Token{Type: token.PLUS, Literal: "+"}, Left: &Node{Token: token.Token{Type: token.INT, Literal: "2"}}}
	node := ast.Node{Token: &next}
	it := iterator{nud: &node}
	accIt, isIt := acc.(iterator)

	if isIt {
		// { , {... }
		node.Left = accIt.nud
		it.stack = accIt.stack
	}

	// end
	if next.Type == token.INT {
		it.left = &node
		it.nud = nil
	}
	
	return it
}
```

It copies stack to maintain the main function's recursion.
Also, it can recognise when to stop denotation. Currently, it's a `token.Type` attribute, which should be equal to `INT`. 

Luckily, that concludes our second step of parsing expressions.

# Summary
