package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getCommand(row string) (string, string) {
	splits := strings.Split(row, " ")
	sub := ""
	if len(splits) > 1 {
		sub = splits[1]
	}
	return splits[0], sub
}

func processCommand(row string, addAt *int, addValue *int) {
	command, sub := getCommand(row)
	if command == "noop" {
		*addAt += 1
		*addValue = 0
	}
	if command == "addx" {
		add, _ := strconv.Atoi(sub)
		*addValue = add
		*addAt += 2
	}
}

func readInput(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	cycle := 0
	register := 1
	addAt := 0
	addValue := 0
	reportAt := 40 // 20 for 1st task
	signalSum := 0
	for scanner.Scan() {
		row := scanner.Text()
		processCommand(row, &addAt, &addValue)
		for i := cycle; i < addAt; i++ {
			if Abs(i%40-register) <= 1 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
			if i == reportAt-1 {
				fmt.Println()
				signalSum += reportAt * register
				reportAt += 40
			}
		}
		register += addValue
		cycle = addAt
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return signalSum
}

func main() {
	signal := readInput("input2")
	fmt.Println(signal)
}
