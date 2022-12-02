package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1Rock: X, A 2Paper: Y, B 3Sciccors Z, C
// Win: 6 Draw: 3 Lose: 1
const (
        A string = "A"
        B string = "B"
        C string = "C"
        X string = "X"
        Y string = "Y"
        Z string = "Z"
)

func main() {
        input, _ := os.Open("input.txt")
        defer input.Close()
        sc := bufio.NewScanner(input)
        score := 0
        for sc.Scan() {
                line := sc.Text()
                arr := strings.Split(line, " ")
                opp := arr[0]
                mine := arr[1]
                value := comp(opp, mine)
                score += value
        }
        fmt.Printf("%v", score)
}

func comp(opp string, mine string) int {
        val := 0
        if mine == X {
                val += 1
                if opp == A {
                        val += 3 
                        return val
                }
                if opp == B {
                        val += 0
                        return val
                }
                if opp == C {
                        val += 6
                        return val
                }
        }
        if mine == Y {
                val += 2
                if opp == A {
                        val += 6 
                        return val
                }
                if opp == B {
                        val += 3
                        return val
                }
                if opp == C {
                        val += 0
                        return val
                }
        }
        if mine == Z {
                val += 3
                if opp == A {
                        val += 0
                        return val
                }
                if opp == B {
                        val += 6
                        return val
                }
                if opp == C {
                        val += 3
                        return val
                }
        }
        return 0

}
