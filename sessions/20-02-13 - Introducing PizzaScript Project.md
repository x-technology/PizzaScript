# Topic Ideas

- Compare JS reactive vs golang
- Benefits of using streams?
  - standard operators, iterations, flows
- Practice Make and explain Lexer example

# Agenda

- Introduction (10 minutes)
- Theory (30 minutes)
  - Programming Languages Intro
  - Lexer, Parser and so on
- Code Review (20 minutes)
  - Go needed concepts - modules, goroutines, channels
  - ReactiveX, rxgo
    - Operators - filter, map, create
    - How to filter with order?

```
>> 1+2
{INT 1}
{+ +}
{INT 2}
>> 1+2 
{+ +}
{INT 1}
{INT 2}
>> 1+2
{INT 1}
{+ +}
{INT 2}
```

    - Different states

```go
func (l *Lexer) Tokens() rxgo.Observable {
	return rxgo.Merge([]rxgo.Observable{
		// operators
		// l.observable.
		// Filter(func(i interface{}) bool {
		// 	var str = i.(string)

		// 	sort.Strings(token.ALL_OPERATORS)
		// 	var index = sort.SearchStrings(token.ALL_OPERATORS, str)

		// 	return token.ALL_OPERATORS[index] == str
		// }).
		// Map(func(_ context.Context, i interface{})(interface{}, error) {
		// 	var str = i.(string)
		// 	ch := []byte(str)[0]
		// 	var index = sort.SearchStrings(token.ALL_OPERATORS, str)

		// 	var tok token.Token = newToken(token.TokenType(token.ALL_OPERATORS[index]), ch)
		// 	return tok, nil
		// }),
		// numbers
		l.observable.
		// Filter(func(i interface{}) bool {
		// 	var str = i.(string)
		// 	ch := []byte(str)[0]

		// 	return isNumber(ch)
		// }).
		DistinctUntilChanged(func(_ context.Context, i interface{}) (interface{}, error) {
			var str = i.(string)
			ch := []byte(str)[0]
			fmt.Println(str)
	
			return isNumber(ch), nil
		}).
		// Reduce(func(_ context.Context, acc interface{}, elem interface{}) (interface{}, error) {
		// 	if acc == nil {
		// 		return elem, nil
		// 	}
		// 	return acc.(string) + elem.(string), nil
		// }),
		// SumInt64()
		Map(func(_ context.Context, i interface{})(interface{}, error) {
			var str = i.(string)
			ch := []byte(str)[0]

			var tok token.Token = newToken(token.INT, ch)
			return tok, nil
		}),
	})
}
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

other problems and final flow with Filter, Scan, Filter (swap)

11=00

scan 1: save 1
scan 11: save 11
scan 11=: return 11, save =
scan =0: return =, save 0
scan 00: save 00
scan .: return 00

  - Explain language techniques

# TODO

- why to use observables?

- check syntax

```go
// type conversion
i.V.(string)
```

- check concept channels, goroutine

# Feedback

# Summary

- Learn lots of Go specifics
- Learn Rx operators
- Program

## Links

- [How to choose operator?](http://xgrommx.github.io/rx-book/content/which_operator_do_i_use/instance_operators.html)