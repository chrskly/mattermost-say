
package main

import (
    "os"
    "fmt"
    "stdinreader"

    "mmsay"
)

func main() {
    input, err := stdinreader.Read()
    _ = err
    //fmt.Println(input)

    resp, err := mmsay.Say(input)
    _ = resp
    if err != nil {
        fmt.Fprintf(os.Stdout, "could not post message: %v\n", err)
    }
}
