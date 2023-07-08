package main

import (
//    "bufio"
	"fmt"
//    "os"
    "regexp"
)


func parse(s, RegExpString string) (string, error) {
    r, err := regexp.Compile(RegExpString)
    if err != nil {
        fmt.Println("Well, your RegExp", RegExpString, "sucks with error: ", err)
    }
    s = r.FindString(s)
    fmt.Println(s)
    return s, err
}


func main() {
    s := "foo"
    // Arabic
    parse(s, `^\s*(\d0?)\s*([\*\\\+=])\s*(\d0?)\s*$`)
    // Roman
    parse(s, `(I{1,3}|IV|VI{0,3}|I?X)\s*([\*\\\+=])\s*(I{1,3}|IV|VI{0,3}|I?X)`)
    
    
    
	fmt.Println("Hello, World!")
}
