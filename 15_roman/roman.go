package main

import (
    "strings"
)

func ConvertToRoman(n int) string {
    var result strings.Builder

    for i := 0; i < n; i++ {
        result.WriteString("I")
    }
    return result.String()
}
