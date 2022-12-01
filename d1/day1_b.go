package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
        input, _ := os.Open("input.txt")
        defer input.Close()
        sc := bufio.NewScanner(input)
        sc.Scan()

        var elfArray [3]int

        curr := 0

        for sc.Scan() {
                fmt.Println(sc.Text())
                if sc.Text() == "" {
                        elfArray = evaluateElfs(elfArray, curr)
                        curr = 0
                }
                number, _ := strconv.Atoi(sc.Text())
                curr += number
        }
        elfArray = evaluateElfs(elfArray, curr)
        fmt.Printf("%d", elfArray[0] + elfArray[1] + elfArray[2])
}

func evaluateElfs(elfArray [3]int, curr int) [3]int {
        if curr > elfArray[0] {
                elfArray[2] = elfArray[1]
                elfArray[1] = elfArray[0]
                elfArray[0] = curr
        } else if curr > elfArray[1] {
                elfArray[2] = elfArray[1]
                elfArray[1] = curr
        } else if curr > elfArray[2] {
                elfArray[2] = curr
        }
        return elfArray
}
