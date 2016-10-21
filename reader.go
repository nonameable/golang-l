package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)


func caesar(r rune, shift int) rune {
    // Shift character by specified number of places.
    // ... If beyond range, shift backward or forward.
    s := int(r) + shift
    if s > 'z' {
    return rune(s - 26)
    } else if s < 'a' {
    return rune(s + 26)
    }
    return rune(s)
}


func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter encrypted text: ")
    text, _ := reader.ReadString('\n')
    fmt.Println(text)

    fmt.Println("The possible charchains that thing could make are: ")
    
    for j := 0; j < 26; j++ {
        result := strings.Map(func(r rune) rune {
        return caesar(r, j)
        }, text)
        fmt.Println(result)
    }
}