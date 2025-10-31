There are many ways that a program gets input. Programs
1. generate their own data
2. read from a file
3. read from a network connection
4. from the output of another program
5. a user at a keyboard
6. command line arguments etc


Command line arguments are available to a program in a variable named os.Args. It is a slice of strings.
The first element of os.Args, os.Args[0], is the name of the command itself. The other elements are the arguments that were presented to the program when it started execution. 


String concatenation using + operator is a memory intensive operation. strings in go are immutable. When you concatenate strings, Golang creates a new string and the old strings are garbage collected. Because of the memory allocation and the subsequent garbage collections, concatenation using + is not preferred. Instead, use strings.Builder type
strings.Builder allocates a buffer and copies the string content to buffer.

```go
package main

import (
    "fmt"
    "strings"
)

func main() {
    var s strings.Builder
    s.WriteString("write what you want here")
    fmt.Println(s.String())
}
```
Run the benchmarks in the exercise3 to know how fast strings.Builder is compared to string concatenation