package main

import (
        "encoding/base64"
        "fmt"
        "strings"
)

func main() {
        //Coding Quiz
        var whatIsIt string
        secret := "aWFuZ25vVzpOQU06RU5JTDp0YTpzdTpuaW9K"
        sd, _ := base64.StdEncoding.DecodeString(secret)
        //Solve
        fmt.Println("Secret: ",secret)
        whatIsIt = (string(sd))
        fmt.Println("Base64 Decoded: ",whatIsIt)
        whatIsIt = Reverse(string(sd))
        whatIsIt = strings.Replace(whatIsIt, ":", " ", -1)
        fmt.Println("Reversed & Replaced: ",whatIsIt)
}

func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
