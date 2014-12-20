express
=======

Not really a programming language that compiles to Go binary.

NOTE: Currently in experimental stage!

## What?
* It just removes curly-braces `{` and `}` to declare an `interface`, `struct`, and `func` and closes it with `end` keyword.
* It just removes the `type` keyword to declare an `interface` and `struct`.
* It just replaces `func` keyword to `fun` to declare a function.
* Then everything is just Go ^_^

## Example
`hello.xps` file
```go
package main

import "fmt"

interface Message
  say() string
end

struct Greeting
  name string
  msg string
end

fun (g Greeting) say() string
  return g.msg + ", " + g.name
end

fun main()
  c := make(chan string)

  go fun()
    g := Greeting{"Gohan", "Hello"}
    c <- g.say()
  end()

  fmt.Println(<-c)
end
```

Then execute the command below to create Go binary `hello` or `hello.exe` (based on your platform). :
```sh
express hello.xps
-> hello
```
Then execute the application:
```sh
hello
-> Hello, Gohan
```

## License
The MIT License (MIT)

Copyright (c) 2014 Maiah Macariola
