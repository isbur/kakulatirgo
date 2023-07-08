package main

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/brandenc40/romannumeral"
	"os"
	"regexp"
	"strconv"
)


const ArabicRegExpStr string = `^\s*([1-9]|10)\s*([\*/\+-])\s*([1-9]|10)\s*$`
const SingleRomanNumberRegExpStr string = `\b(?:I{1,3}|IV|VI{0,3}|I?X)\b`
const RomanRegExpStr string = `^\s*(` + SingleRomanNumberRegExpStr + `)\s*([\*/\+-])\s*(` + SingleRomanNumberRegExpStr + `)\s*$`


func parse(s, RegExpString string) []string {
	r, _ := regexp.Compile(RegExpString)
	return r.FindStringSubmatch(s)
}


func main() {
    var RomanModeEnabled bool = false
    var parsedExpression []string
	reader := bufio.NewReader(os.Stdin)
    s, _ := reader.ReadString('\n')
	parsedExpression = parse(s, ArabicRegExpStr)
    if parsedExpression == nil {
        parsedExpression = parse(s, RomanRegExpStr)
        if parsedExpression == nil {
            fmt.Print(errors.New("Inappropriate input. Please, check whether there are only numbers from 1 to 10, single sign /*+- between them and whether you haven't mixed up Arabic and Roman digits."))
            return
        } else {
            RomanModeEnabled = true
        }
    }

    if RomanModeEnabled {
        RomanToArabic := map[string]string {
            "I":"1",
            "II":"2",
            "III":"3",
            "IV":"4",
            "V":"5",
            "VI":"6",
            "VII":"7",
            "VIII":"8",
            "IX":"9",
            "X":"10",
        }
        parsedExpression[1] = RomanToArabic[parsedExpression[1]]
        parsedExpression[3] = RomanToArabic[parsedExpression[3]]
    }

    a, _ := strconv.Atoi(parsedExpression[1])
    b, _ := strconv.Atoi(parsedExpression[3])
    sign := parsedExpression[2]
    var res int
    
    switch sign {
        case "*":
            res = a * b
        case "/":
            res = a / b
        case "+":
            res = a + b
        case "-":
            res = a - b
    }

    if RomanModeEnabled {
        // First try
        // ArabicToRoman := map[int]string{
        //     1: "I",
        //     2: "II",
        //     3: "III",
        //     4: "IV",
        //     5: "V",
        //     6: "VI",
        //     7: "VII",
        //     8: "VIII",
        //     9: "IX",
        //     10: "X",
        //     12: "",
        //     14: "",
        //     15: "",
        //     16: "",
        //     18: "",
        //     20: "",
        //     21: "",
        //     24: "",
        //     27: "",
        //     28: "",
        //     30: "",
        //     32: "",
        //     36: "",
        // }
        rom_res, err := romannumeral.IntToString(res)
        if err != nil {
            fmt.Println(errors.New("Result seems to be not representable by Roman digits; try another input."))
            return
        }
        fmt.Println(rom_res)
    }

    if !RomanModeEnabled {
        fmt.Println(res)
    }
    
}
