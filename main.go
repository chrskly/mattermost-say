
package main

import (
    "fmt"
    "stdinreader"
)

func main() {
    input, err := stdinreader.Read()
    _ = err
    fmt.Println(input)
}
