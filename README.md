express
=======

Not really a programming language that compiles to Go binary

NOTE: Currently in experimental stage!

## What?
It just removes curly-braces and removes `type` keyword to declare `interface`s and `struct`s in Go code. Then everything is just Go ^_^

## Example
`hello.exp` file
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

fun (g Greeting, a string, b string, s string) say() string
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

Then execute the command below:
```sh
express hello.exp
```

This will create Go binary `hello` or `hello.exe` (based on your platform). Then execute the binary:
```sh
hello
-> Hello, Gohan
```

## License
The MIT License (MIT)

Copyright (c) 2014 Maiah Macariola
