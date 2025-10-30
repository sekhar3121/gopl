Run the program using the following command
```shell
go run chapter1/section1/helloworld/main.go
```
**go run** command compiles the source code from one or more source files whose names end in .go, links it with libraries, then runs the resulting executable file.

```go
fmt.Println("Hello, 世界")
```
Go natively handles Unicode. This means that the Go programming language has built-in support for working with Unicode text, which includes characters from virtually every written language in the world - not just ASCII or English.

```shell
go build main.go
```
go build command compiles the .go files, links it with libraries, and creates an executable binary file.