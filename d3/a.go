package main

import (
	"bufio"
	"fmt"
	"os"

)

const OFFSET = 96

const LG_OFFSET = 64

func main() {
        input, _ := os.Open("input.txt")
        defer input.Close()
        sc := bufio.NewScanner(input)
        iter := 0
        for sc.Scan() {
                line := sc.Text()
                items := []rune(line)
                fmt.Printf("%v\n", string(getMatchingItem(items)))
                value := getItemValue(getMatchingItem(items))
                iter = int(value) + iter
                fmt.Printf("%v\n", iter)
                // fmt.Printf("%v \n", items)
        }
        fmt.Printf("%v\n", iter)
}

func getItemValue(item rune) rune {
        if item - OFFSET > 0 {
                return item - OFFSET
        }
        return item - LG_OFFSET + 26
}

func getMatchingItem(items []rune) rune {
        first := items[0:len(items)/2]
        second := items[len(items)/2:]
        for _, item := range first {
                if contains(item, second) {
                        return item
                }
        }
        return -1
}


func contains(item rune, items []rune) bool {
        for _, i := range items {
                if i == item {
                        return true
                }
        }
        return false
}

