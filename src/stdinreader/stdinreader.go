
package stdinreader

import (
    "bufio"
    "os"
    "strings"
)

func Read() (string, error) {
    allInput := ""
    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        allInput = allInput + string(scanner.Text()) + "\n"
    }
    if scanner.Err() != nil {
        return allInput, scanner.Err()
    }
    return strings.TrimSpace(allInput), nil
}
