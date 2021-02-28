# PizzaScript

> Programming language that fucks up (c)

`PizzaScript` or `ps` is a new cool programming language with no-ordinary paradigm. 

`PizzaScript` compiles to `WebAssembly`, which makes it portable and suitable target for execution both on client and server sides. 

## Goals

- learn Go language, including key libraries like RxGo
- deeper understanding of WebAssembly
- understand how programming languages & interpreters work

## Technologies

- programming languages 
- lexer, parser
- rxgo operators
- webassembly

## Key Features 

- Cool name (and logo)!

That's maybe the most important of all features. We are proud of our choice.

- `JavaScript` like syntax, kind of. That's true, we are inspiring by official browser language
- Variables, integers, floats and strings
- Arithmetic expressions, built-in functions
- Functions, first class functions, closures
- Dynamic types and coercions
- Modules & Standard library
- Awesome ideas and interpreter flows

```ps
var h1679 = "1"
val g2788 = 2

h1679 + g2788 == "12"
g2788 + h1679 == 3
```

- Parentheses and curly braces are treated appropriately (meaning, they can be omitted in most cases)

> You write your code, and let `PizzaScript` do the rest!

## Way of working

- We follow the best programming techiques and standarts we want (know? can?)
- We release every month
- We compile to `WebAssembly`
- We publish changelogs

## Syntax

### Examples

```ps
export "pizza/io/print"

fun factorial(n: int): int {
  // curly braces can be omitted if context has single meaning
  if n = 0
    1
  else
    n * factorial n - 1
  // there is no need for return statement inside function body
}

for var n := 0; n <= 16; n++ {
  print n, '! = ', factorial n
  // you can choose to specify parentheses
  // print(n, '! = ', factorial n)
}
```

### Common

Comments until end of line are declared with double slash `//` sign.

### Arrays

> Arrays are immutable in Monkey

### Control Flow

We need to discuss the default language behavior for normal expressions like

```ps
5 + 5 * 10 // 100
```

The goal for the PizzaScript language is to take everyone equally, so for all operators the compiler treats every expression with the same priority.

### Variables

> Declaration operators

As we know, `Kotlin` has amazing operators `var, val`, which is very similar to `JavaScript` `let` and `const` respectively. We [decided](todo link to decision) to borrow best parts from these examples: we going to use `Kotlin` notation, and we keep hoisting.

```ps
export "pizza/io/print"

a1573 = "1"
print a1573 // "1"
var a1573: string = "null"

// b2217 = 1
print b2217 // undefined
val b2217: int = 34
```

> All variables do hoist. No, hoisting is not to be discussed

Both `var` and `val` operator variables are hoisted. Meanining, a variable is accessible within an entire function or global scope. Therefore, print could output variables before actual declaration.

```ps
c6572 = 1
print c6572 // 1.0
c6572: float = 3
```

Great thing, that you can omit a declaration operator, that should help to keep program code cleaner. When no operator is declared, by default the `PizzaScript` interpreter will treat it as `var`, but it is possible to change the setting in [a language configuration](todo language configuration file).

```json
{
  "Hoisting": "false"
}
```

```ps
print d1317 // 3.0
d1317 = 3
```

In case a declaration is missing, interpreter needs to execute all expressions with such a variable in order to know variable's type. Value comes for free in such scenario.

> It is possible to use semicolons, but not obligatory

### If Else Statements

Control flow in `PizzaScript` language is very similar to common 

```ps
if 5 < 10 {
  true
} else {
  false
}
```

No parentheses or curly braces are required inside the `if-else` statement.

### Data Types

That's simple, as `WebAssembly` supports 4 types - `i32, i64, f32, f64` we decided to support also `string`. In our case all types are `int, float, and string` and type size is decided on compilation time. At least, in the beginning. That should be enough for now. We [have plans](./todo json) to support `json` and [other types](./todo other types).

Of course, we also need `undefined` type. To point out that a variable is not defined yet.

Variable types are leaded to the result type according to coercion priority table ...

> #PizzaScript 🍕programming language appreciates your help, however can handle the thing itself

### Functions

Functions are declared with the `fun` keyword.

```ps
fun sum(var a1573: string, b7232): int {
  a1573 + b7232
}

sum(1, 2) // 12
```

A function definition can contain any number of arguments written inside parenthesis, separated by comma.

If arguments and function types are specified, that affects type coercion and affects the end result.

> Function body is the result! (😲)

If a function is called in a simple context, parenthesis can be omitted.

### Async Await

As all cool languages, `PizzaScript` is asyncronous. 

```ps
await print a1573
var a1573: string = "null"
```

### Generics

We declare and reserve generics functionality for the future without description and obligations.

### White space

In some languages like `Python`, whitespaces are strict. Not in `PizzaScript` - a developer can use any indent.

## Standard library

Every language needs standard library, and `PizzaScript` has one.
We start with the simplest and at the same time, key functionality.

### pizza/io

`pizza/io` contains functions to input & output values somewhere. 
`pizza/io/print` is a `print` function, useful to debug your code. 

## Compiler

- lexer
- parser
- abstract syntax tree (AST)
- webassembly transformation
- webassembly glue code, including standard library

## Next steps

As soon as the first article gets 10 different people upvotes, we immediately start preparing a next release. Otherwise, we will stick with formal way of working. 

## Tags

- 100DaysOfCode
- 100daysofcodechallenge
- CodeNewbies
- mentoring
- developer
- golang
- ReactiveX
- webdevelopment
- OnlineEvents
- OnlineSeminars
- programming
- development
- introduction
- algorithms
- compiler
- interpreter
- programminglanguages

## TODO

- pizzascript as a topic:
  + learn new programming language
  + learn theory
  - not an expert in language
  - not an expert in theory

- [Crockford](http://crockford.com/javascript/)
- [`8` Cool introduction to PL grammars](http://www.mollypages.org/page/grammar/index.mp)

- Use existing language??
- Pratt parser
- regular activities - twitter start in January

- Запланировать твит о книге и купить го интерпретатор книгу - as promotion to first session

twit as group messages about the book
- twit "book gives an impression on how languages work, especially like javascript and go - builtin functions, environments, etc"

> That’s because the whole idea behind our Pratt parser hinges on the idea of precedences and we haven’t defined the precedence of our index

- twit "learning an awesome programming language even more, thanks for inspiration for this technical topic"
- twit "Tony Hoare introduced null references to the ALGOL W language in 1965 and called this his “billion-dollar mistake”"
- twit "cumbersome java example"
- twit "for _, value := range pow" very interesting decision to hide not needed variables https://tour.golang.org/moretypes/17
- twit also surprise decisions like global delete, append, etc. functions

- decompose projects into 2 hours sessions

- https://interpreterbook.com/

./assets/Writing An Interpreter In Go by Thorsten Ball (z-lib.org).pdf

- > Unlike other JVM languages, it's not that straightforward to create simple “Hello, World!” program in Clojure.

Should be something amazing!

- implementation detail - rxjs / rxgo? streams to split into tokens, then organize into ast

- github account for pizza script, see typescript
- language specification repository, compiler repository
- articles, changelog
- imgs
- WOW language specification contribution
- twitter
- standart library `pizza`

# Links

- [wat](https://www.destroyallsoftware.com/talks/wat)
- [wtfjs](https://github.com/denysdovhan/wtfjs)
- [Pixel - Language Study](http://rigaux.org/language-study/syntax-across-languages/)
- [Terence Parr - ANTLR4](https://vimeo.com/59285751)

- [Introduction to JVM Languages - Vincent van der Leun](...), [code](https://github.com/PacktPublishing/Introduction-to-JVM-Languages)
- [Language Implementation Patterns](https://www.oreilly.com/library/view/language-implementation-patterns/9781680500097/)
- [Emerging Programming Languages](https://www.oreilly.com/library/view/emerging-programming-languages/9781492082590/)
- [Top Down Operator Precedence - Vaughan Pratt](...)
- [Pratt Parsers: Expression Parsing Made Easy by Bob Nystrom](http://journal.stuffwithstuff.com/2011/03/19/pratt-parsers-expression-parsing-made-easy/)

TODO links

Rob Pike - Lexical Scanning in Go - https://www.youtube.com/watch?v=HxaD_trXwRE

https://www.youtube.com/watch?v=PXoG0WX0r_E

TODO https://dev.to/jrop/pratt-parsing

- [WebAssembly AST Playground](http://ast.run/)
