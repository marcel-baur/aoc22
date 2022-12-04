package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
                s1 := strings.Split(line, ",")
                elf1 := strings.Split(s1[0], "-")
                elf2 := strings.Split(s1[1], "-")
                e1 := []int{0,0}
                e2 := []int{0,0}
                for idx, elf := range elf1 {
                        conv, _ := strconv.Atoi(elf)
                        e1[idx] = conv
                }
                for idx, elf := range elf2 {
                        conv, _ := strconv.Atoi(elf)
                        e2[idx] = conv
                }
                // 6/18, 14/19
                if e1[1] >= e2[0] && e1[0] <= e2[1] {
                        iter++
                        continue
                }
                if e2[1] >= e1[0] && e2[0] <= e1[1] {
                        iter++
                        continue
                }
                // fmt.Printf("%v \n", items)
        }
        fmt.Printf("%v\n", iter)
}
