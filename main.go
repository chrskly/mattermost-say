
package main

import (
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
        fmt.Errorf("could not post message: %v", err)
    }
}
