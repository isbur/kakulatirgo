package main


import (
    "regexp"
    "testing"
)


func TestArabic (t *testing.T) {
    var s string
    var r *regexp.Regexp
    var mustMatch = []string {
        "1+1",
        "1 + 1",
        "1      + 1",
        "1 * 10",
        "2/2",
        "2 -2",
        "10 - 10",
        "  2+2  ",
    }
    r, _ = regexp.Compile(ArabicRegExpStr)
    for _, s = range mustMatch {
        if !r.MatchString(s) {
            t.Fatal("Didn't matched:", s)
        }
    }
    var mustNotMatch = []string {
        "0+1",
        "11-23",
        "90+60",
        "23",
        "3 + III",
        "1 + 2 jg",
    }
    for _, s = range mustNotMatch {
        if r.MatchString(s) {
            t.Fatal("False match with:", s)
        }
    }
}


func TestRoman (t *testing.T) {
    var s string
    var r *regexp.Regexp
    var mustMatch = []string {
        "I+II",
        "III + IV",
        "V      + VI",
        "VII * VIII",
        "IX/X",
        "I -I",
        "X - X",
        "  II+II  ",
    }
    r, _ = regexp.Compile(RomanRegExpStr)
    for _, s = range mustMatch {
        if !r.MatchString(s) {
            t.Fatal("Didn't matched:", s)
        }
    }
    var mustNotMatch = []string {
        "0+I",
        "XI-II",
        "C+L",
        "X",
        "I + II jg",
    }
    for _, s = range mustNotMatch {
        if r.MatchString(s) {
            t.Fatal("False match with:", s)
        }
    }
}


func TestSingleRomanNumber (t *testing.T) {
    var s string
    var r *regexp.Regexp
    var mustMatch = []string {
        "I", "II", "III", "IV",
        "V", "VI", "VII", "VIII",
        "IX", "X",
    }
    r, _ = regexp.Compile(SingleRomanNumberRegExpStr)
    for _, s = range mustMatch {
        if !r.MatchString(s) {
            t.Fatal("Didn't matched:", s)
        }
    }
    var mustNotMatch = []string {
        "0",
        "IIIII",
        "IIII",
        "IXI",
        "C",
        "L",
        "XI",
        "XII",
        "XIII",
    }
    for _, s = range mustNotMatch {
        if r.MatchString(s) {
            t.Fatal("False match with:", s)
        }
    }
}