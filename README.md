express
=======

Not really a programming language that compiles to Go binary.

NOTE: Currently in experimental stage!

## What?
* It just removes curly-braces `{` and `}` to declare an `interface`, `struct`, `func`, if-else, and for-loop and closes it with `end` keyword.
* It just removes the `type` keyword to declare an `interface` and `struct`.
* It just replaces `func` keyword to `fun` to declare a function.
* Then everything is just Go ^_^

## Example
`hello.xps` file
```go
package main

import "fmt"

fun main()
  c := make(chan string)

  go fun()
    c <- "Hello, Gohan"
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
./hello
-> Hello, Gohan
```

## License
The MIT License (MIT)

Copyright (c) 2014 Maiah Macariola
