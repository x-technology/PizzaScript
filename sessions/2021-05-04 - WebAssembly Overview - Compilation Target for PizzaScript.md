# WebAssembly Overview - Compilation Target for PizzaScript

- Introduction
- WebAssembly
  - Design
  - Key Features
  - Syntax
  - Examples
- Demo
  - What about PizzaScript?
- Summary
  - Feedback

- Get to know `WebAssembly`, understand goals & definitions
- See `WebAssembly` programmatic usage with Web

# PizzaScript

![pizzascript](/assets/pizzascript.jpg)

- Learn `Go` language, and key libraries like `RxGo`
- Understand how programming languages & interpreters work
- Experiment with `WebAssembly`

# WebAssembly

![web assembly](/assets/web-assembly-logo.png)

Bringing other languages to Web since 2014?!

[Windows 95 with WebAssembly](https://archive.org/details/win95_in_dosbox)

### WebAssembly (abbreviated Wasm) is a **binary instruction format** for **a stack-based virtual machine**. Wasm is designed as a **portable target** for **compilation of high-level languages** like `C/C++/Rust`, enabling deployment on the web for client and server applications.

![web assembly](/assets/web-assembly-compile-target-architecture.png)

## [Code First](https://mbebenita.github.io/WasmExplorer/)

```c
int add(int a, int b) {
  return a + b;
}
```

```wat
(module
  (func $add (param $lhs i32) (param $rhs i32) (result i32)
    get_local $lhs
    get_local $rhs
    i32.add)
  (export "add" (func $add))
)
```

```javascript
WebAssembly.instantiateStreaming(fetch(`program.wasm`))
  .then(prog => {
    console.log(prog.instance.exports.add(1, 2))
  })
```

[WasmExplorer Explorer](https://mbebenita.github.io/WasmExplorer/) - nice playground to compile C/C++ to `WebAssembly`

## [History](https://www.youtube.com/watch?v=6r0NKEQqkz0)

- `9 December 2011`, Google Native Client (`NaCl` and `PNaCl`) 😅
- `11 October 2013` - `18 August 2014`, Mozilla [`asm.js`](http://asmjs.org/) - `JavaScript` subset
- `2015 - 2017`, WebAssembly Project Announce
  - `WebAssembly Working Group`
  - `Core Specification`, 2016
  - Official logo, 2017 😂 
  - *Browser Preview*

![web assembly](/assets/web-assembly-logo.png)

- `2017 --> WebAssembly 1.0 MVP -->` [Proposals and WIP](https://github.com/WebAssembly/proposals)

> The initial (MVP) WebAssembly API and binary format is complete to the extent that no further design work is possible without implementation experience and significant usage

## [Use-Cases](https://webassembly.org/docs/use-cases/) and usage examples

- [Windows 95 with WebAssembly](https://archive.org/details/win95_in_dosbox)
- `ZIP` for `WebAssembly` ?!
  - [Almost](https://github.com/YuJianrong/node-unrar.js)

- [Doom 3](http://www.continuation-labs.com/projects/d3wasm/)
  - [Demo](http://wasm.continuation-labs.com/d3demo/)

- Dynamic [`Polyfills` not only for `Web`](https://developer.mozilla.org/en-US/docs/WebAssembly/existing_C_to_wasm)
- [Games](https://hackernoon.com/games-build-on-webassembly-3679b3962a19)

### [Languages and Features Support](https://github.com/appcypher/awesome-wasm-langs)

![webassembly-new-features-browser-support](/assets/webassembly-new-features-browser-support.png)

![browsers support](/assets/wasm-browser-support.png)

# [Design Goals](https://webassembly.github.io/spec/core/intro/introduction.html#design-goals)

- **Fast**: executes with near native code performance
- **Safe**: code is validated and executes in a memory-safe, sandboxed environment
- **Compact**: has a binary format that is fast to process
- **Modular**: programs can be split up in smaller parts
- **Efficient**: can be decoded, validated, and compiled in a fast single pass
- **Platform-independent**

# Definitions & [MVP](https://webassembly.org/docs/mvp/)

- `Modules` and `JavaScript API` in secure environment
- `Binary format (wasm)` - fast binary encoded format
- `Text format` - text format for debugging
- `wasm` engine design to be implemented by browsers and other environments
- [WebAssembly High-Level Goals](https://webassembly.org/docs/high-level-goals/)
  - *execute in the same semantic universe as JavaScript* 🤔

- **No Support**
  - Garbage Collector (proposal...) 
  - DOM Access
  - Old Browsers...

And more!
- [x] Threads
  
## `Stack-based Virtual Machine`?

![web assembly actors](/assets/web-assembly-actors.png)

![Stack-based Virtual Machine](/assets/Stack_3.png)  

## Capabities

- Data Types
  - void i32 i64 f32 f64
- Data Operations
  - i32: + - * / % << >> >>> etc
  - i64: + - * / % << >> >>> etc
  - f32: + - * / sqrt ceil floor
  - f64: + - * / sqrt ceil floor
- Structured Control Flow
  - if loop block br switch
- Functions
- State: linear memory

# JavaScript API

- `WebAssembly` object
  - `Module`, `Instance`, `Memory`, `Table`
  - `validate()` 
  - `compile()`
  - `instantiate()`

```js
var importObj = {js: {
  import1: () => console.log("hello,"),
  import2: () => console.log("world!"),
}};

fetch('demo.wasm').then(response =>
  response.arrayBuffer()
).then(buffer =>
  WebAssembly.instantiate(buffer, importObj)
).then(({module, instance}) =>
  instance.exports.f()
);
```

## [Modules](https://webassembly.org/docs/modules/)

> The distributable, loadable, and executable unit of code in WebAssembly

- `imports`: `js, env, table, memory`

```wat
(module
  (import "js" "import1" (func $i1))
  (import "js" "import2" (func $i2))
  (func $main (call $i1))
  (start $main)
  (func (export "f") (call $i2))
)
```

- <a href="https://github.com/WebAssembly/proposals/issues/12">🛤 ECMAScript module integration</a>

```html
<script type="module" src="proposal.wasm"></script>
```

## Tooling & Compilation

- `Emscripten`

```
C/C++/Rust -> AST -> Binary (.wasm) -> AST -> ...Module
```

[![web assembly compile flow diagram](/assets/webassembly-v8-js-vs-wasm.png)](https://youtu.be/njt-Qzw0mVY?t=1135)

![web assembly compile flow diagram](/assets/web-assembly-compile-flow-diagram.png)

```bash
# cmake
PATH="/Applications/CMake.app/Contents/bin":"$PATH"
cmake
# wabt
git clone --recursive https://github.com/WebAssembly/wabt
cd wabt
mkdir build
cd build
cmake ..
cmake --build .
PATH=$PATH:$(pwd)
wasm-decompile --help
# emsdk
git clone https://github.com/emscripten-core/emsdk.git
cd emsdk
git pull
./emsdk install latest
./emsdk activate latest
source ./emsdk_env.sh
```


### [WebAssembly System Interface (WASI)](https://github.com/bytecodealliance/wasmtime/blob/master/docs/WASI-intro.md)

[![wasi software architecture](/assets/wasi-software-architecture.png)](https://github.com/bytecodealliance/wasmtime/blob/master/docs/wasi-software-architecture.png)

# Performance

![web-assembly-performance](/assets/web-assembly-performance1.png)
![web-assembly-performance](/assets/web-assembly-performance2.png)

# [Demo](otus/webassembly/index.html)

- [Simple Add Function](otus/webassembly/add.wat) 
- [Call Imported API](otus/webassembly/import.wat) 
- [Store API](otus/webassembly/store.wat) 

## [Fibonacci in c, js, and performance](/Users/rd25xo/Developer/experiments/otus/webassembly/fibonacci.sh)

## [Go Hello World](https://github.com/golang/go/wiki/WebAssembly#getting-started)

```bash
mkdir docs
GOOS=js GOARCH=wasm go build -o docs/main.wasm wasm.go
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" docs
npx http-serve -p 8081 docs
```

## What About PizzaScript

- [Output wat](https://webassembly.studio/)

```wat
(module
  (import "console" "log" (func $log (param i32)))
  (func $add (result i32)
    i32.const 13
    i32.const 13
    i32.add
    i32.const 13
    i32.const 13
    i32.add
    i32.add)
  (export "add" (func $add))
  (func (export "logIt")
    call $add
    call $log)
)
```

- Make a simple module
- Wasm output by `wabt`

## Summary

- `WebAssembly` is a highly effective inter programming language protocol, which can be executed in browser and other environments in a secure context

## Links

- [Build Your Own WebAssembly Compiler - Colin Eberhardt, QCon San Francisco 2019](https://www.youtube.com/watch?v=OsGnMm59wb4)
- [Compiling for the Web with WebAssembly (Google I/O '17)](https://www.youtube.com/watch?v=6v4E6oksar0)
- [WebAssembly](https://webassembly.org/)
- [WebAssembly: Disrupting JavaScript - Dan Callahan](https://www.youtube.com/watch?v=7mBf3Gig9io)
- [Why we Need WebAssembly - An Interview with Brendan Eich - Eric Elliott](https://medium.com/javascript-scene/why-we-need-webassembly-an-interview-with-brendan-eich-7fb2a60b0723)
- [WebAssembly Explorer](https://mbebenita.github.io/WasmExplorer/)
- [WebAssembly for Web Developers (Google I/O ’19)](https://www.youtube.com/watch?v=njt-Qzw0mVY)
- [WebAssembly Threads - HTTP 203](https://www.youtube.com/watch?v=x9RP-M6q2Mg)
