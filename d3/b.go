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
        idx := 0
        var group [3][]rune
        for sc.Scan() {
                line := sc.Text()
                items := []rune(line)
                group[idx] = items
                if idx == 2 {
                        idx = 0
                        iter += int(getItemValue(getMatchingItem(group)))
                        continue
                }
                idx++
        }
        fmt.Printf("%v\n", iter)
}

func getItemValue(item rune) rune {
        if item - OFFSET > 0 {
                return item - OFFSET
        }
        return item - LG_OFFSET + 26
}

func getMatchingItem(items [3][]rune) rune {
        r := make(map[rune][]bool)
        // primes := []int{2, 3, 5}
        for idx, its := range items {

                for _, item := range its {
                        if r[item] == nil {
                                r[item] = []bool{false,false,false}
                        }
                        r[item][idx] = true
                }
        }
        for val, slc := range r {
                if slc[0] && slc[1] && slc[2] {
                        return val
                }
        }
        fmt.Println("ERROR")
        return items[0][0]
}


func contains(item rune, items []rune) bool {
        for _, i := range items {
                if i == item {
                        return true
                }
        }
        return false
}

