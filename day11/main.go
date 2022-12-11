package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type monkey struct {
	id        int
	items     []int
	op1       string
	op2       string
	opType    string
	opValue   int
	divisible int
	divTrue   int
	divFalse  int
}

func readInput(fileName string) ([]monkey, int) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reg_monkey_id := regexp.MustCompile(`Monkey\s?(\d+):`)
	reg_starting_items := regexp.MustCompile(`\s*Starting items:\s([\d,\s]+)`)
	reg_operation := regexp.MustCompile(`\s*Operation:\s(\w+)\s=\s(\w+)\s([*+\-/])\s(\w+)`)
	reg_test := regexp.MustCompile(`\s*Test:\sdivisible by (\d+)`)
	reg_throw := regexp.MustCompile(`throw to monkey (\d+)`)

	var monkeys []monkey

	scanner := bufio.NewScanner(file)
	index := 0

	int_monkey_id := 0
	var starting_items []int
	op1 := ""
	op2 := ""
	op_type := ""
	op_value := 0
	divisible := 0
	int_throw_true := 0
	int_throw_false := 0
	maxWorriness := 1

	for scanner.Scan() {
		index += 1
		row := scanner.Text()
		if index%7 == 1 {
			monkey_id := reg_monkey_id.FindAllStringSubmatch(row, -1)
			int_monkey_id, _ = strconv.Atoi(monkey_id[0][1])
		}

		if index%7 == 2 {
			str_starting_items := reg_starting_items.FindAllStringSubmatch(row, -1)
			split_starting_items := strings.Fields(str_starting_items[0][1])
			for _, str_item := range split_starting_items {
				if str_item[len(str_item)-1] == ',' {
					str_item = str_item[:len(str_item)-1]
				}
				item, _ := strconv.Atoi(str_item)
				starting_items = append(starting_items, item)
			}
		}

		if index%7 == 3 {
			operation := reg_operation.FindAllStringSubmatch(row, -1)
			op1 = operation[0][2]
			op_type = operation[0][3]
			if operation[0][4] == "new" || operation[0][4] == "old" {
				op2 = operation[0][4]
			} else {
				op_value, _ = strconv.Atoi(operation[0][4])
			}
		}

		if index%7 == 4 {
			test := reg_test.FindAllStringSubmatch(row, -1)
			divisible, _ = strconv.Atoi(test[0][1])
		}

		if index%7 == 5 {
			throwTrue := reg_throw.FindAllStringSubmatch(row, -1)
			int_throw_true, _ = strconv.Atoi(throwTrue[0][1])
		}

		if index%7 == 6 {
			throwFalse := reg_throw.FindAllStringSubmatch(row, -1)
			int_throw_false, _ = strconv.Atoi(throwFalse[0][1])
		}

		if index%7 == 0 {
			monkeys = append(monkeys, monkey{
				id:        int_monkey_id,
				items:     starting_items,
				op1:       op1,
				op2:       op2,
				opType:    op_type,
				opValue:   op_value,
				divisible: divisible,
				divTrue:   int_throw_true,
				divFalse:  int_throw_false,
			})

			maxWorriness *= divisible
			starting_items = []int{}
			op2 = ""
			op_value = 0
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return monkeys, maxWorriness
}

func calcWorry(input int, op1 string, op2 string, opType string, opValue int, maxWorriness int) int {
	opLeft := 0
	if op1 == "old" {
		opLeft = input
	}
	if op2 == "old" {
		opValue = input
	}

	result := -1
	if opType == "+" {
		result = opLeft + opValue
	}
	if opType == "*" {
		result = opLeft * opValue
	}
	if opType == "-" {
		result = opLeft - opValue
	}
	if opType == "/" {
		result = opLeft / opValue
	}

	return result % maxWorriness
}

func shenanigans(monkeys *[]monkey, maxWorriness int) {
	inspects := make([]int, len(*monkeys))
	for round := 1; round <= 10000; round++ {
		for i := 0; i < len(*monkeys); i++ {
			m := (*monkeys)[i]
			for _, item := range m.items {
				inspects[i] += 1
				worry := calcWorry(item, m.op1, m.op2, m.opType, m.opValue, maxWorriness)
				mNewIndex := -1
				//worry /= 3
				if worry%m.divisible == 0 {
					mNewIndex = m.divTrue
				} else {
					mNewIndex = m.divFalse
				}
				(*monkeys)[mNewIndex].items = append((*monkeys)[mNewIndex].items, worry)
			}
			(*monkeys)[i].items = []int{}
		}
	}

	for i := 0; i < len(*monkeys); i++ {
		m := (*monkeys)[i]
		fmt.Printf("%d,%d ", i, inspects[i])
		fmt.Println(m.items)
	}
}

func main() {
	monkeys, maxWorriness := readInput("input2")
	//maxWorriness = gcd(maxWorriness)
	shenanigans(&monkeys, maxWorriness)
	fmt.Println(monkeys)
}
