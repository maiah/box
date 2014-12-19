box
===

Programming language that compiles to Go binary

## Example
`hello.box` file
```go
parcel main

deploy "fmt"

cargo Message iopen
	say() string
close

cargo Greeting sopen
	name string
	msg string
close

box (g Greeting) say() string open
	return g.msg + ", " + g.name
close

box main() open
  c := make(chan string)

  go box() open
    g := Greeting open "Gohan", "Hello" close
    c<- g.say()
  close()

  fmt.Println(<-c)
close
```

Then execute the command below:
```sh
box hello.box
```

This will create Go binary `hello` or `hello.exe` (based on your platform). Then execute the binary:
```sh
hello
-> Hello, Gohan
```
