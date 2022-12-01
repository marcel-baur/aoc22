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

        mostCals := 0
        curr := 0

        for sc.Scan() {
                fmt.Println(sc.Text())
                if sc.Text() == "" {
                        if curr > mostCals {
                                mostCals = curr
                        }
                        curr = 0
                }
                number, _ := strconv.Atoi(sc.Text())
                curr += number
        }
        fmt.Printf("%d", mostCals)
}
