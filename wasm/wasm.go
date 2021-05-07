package main

import (
	"syscall/js"

	"pizzascript/repl"
)

var c chan bool

func printMessage(this js.Value, inputs []js.Value) interface{} {
	message := inputs[0].String()
	println(message)

	output := repl.Compile(message)
	println(output)

	return output
}

func main() {
	println("Feel free to type in pure PizzaScript code\n")

	// js.Global.Set("test", "123")
	// js.Global().Set("printMessage", js.FuncOf(printMessage))

	c := make(chan bool)
	js.Global().Set("printMessage", js.FuncOf(printMessage))
	<-c
}
