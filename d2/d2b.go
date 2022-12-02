package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 1Rock: LOSE, ROCK 2Paper: DRAW, PAPER 3Sciccors WIN, SCISSORS
// Win: 6 Draw: 3 Lose: 0
const (
        ROCK string = "A"
        PAPER string = "B"
        SCISSORS string = "C"
        LOSE string = "X"
        DRAW string = "Y"
        WIN string = "Z"
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
                value := match(opp, mine)
                fmt.Printf("Opponent: %v, Me: %v, value: %v \n", opp, mine, value)
                score += value
        }
        fmt.Printf("%v", score)
}

func signValue(sign string) int {
        switch sign {
                case PAPER: return 2 
                case SCISSORS: return 3 
                case ROCK: return 1 
        }
        fmt.Println("ERROR SIGN VALUE")
        return 0
}

func counterTo(sign string) string {
        switch sign {
                case ROCK: return PAPER
                case PAPER: return SCISSORS 
                case SCISSORS: return ROCK 
        }
        fmt.Println("ERROR COUNTER TO")
        return ""
}

func counterFrom(sign string) string {
        switch sign {
                case ROCK: return SCISSORS
                case PAPER: return ROCK 
                case SCISSORS: return PAPER 
        }
        fmt.Println("ERROR COUNTER FROM")
        return ""
}

func match(opp string, signal string) int {
        if signal == DRAW {
                return signValue(opp) + 3 
        }
        if signal == WIN {
                return signValue(counterTo(opp)) + 6
        }
        if signal == LOSE {
                return signValue(counterFrom(opp))
        }
        fmt.Println("ERROR MATCH")
        return 0
}
