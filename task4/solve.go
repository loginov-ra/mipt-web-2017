package main

import (
    "fmt"
    "strings"
    "unicode"
)

func RemoveEven(slice []int) []int {
    var resultSlice []int
    for i := 0; i < len(slice); i++ {
        if slice[i] % 2 != 0 {
            resultSlice = append(resultSlice, slice[i])
        }
    }
    return resultSlice
}

func PowerGenerator(radix int) func() int {
    currDegree := 1
    return func() int {
        currDegree *= radix
        return currDegree
    }
}

func DifferentWordsCount(str string) (int) {
    s := strings.ToLower(str)
    var currWord string
    dict := make(map[string]int)

    for i := 0; i < len(s); i++ {
        if !unicode.IsLetter(int32(s[i])) {
            if len(currWord) > 0 {
                dict[currWord] = 1
            }
            currWord = ""
        } else {
            currWord += string(s[i])
        }
    }
    
    if len(currWord) > 0 {
        dict[currWord] = 1
    }
    
    return len(dict)
}
