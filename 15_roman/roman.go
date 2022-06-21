package main

import (
    "strings"
)

func ConvertToRoman(n int) string {
    var result strings.Builder

    for n > 0 {
        switch {
        case n > 4:
            result.WriteString("V")
            n -= 5
        case n > 3:
            result.WriteString("IV")
            n -= 4
        default:
            result.WriteString("I")
            n--
        }
    }
    return result.String()
}
