package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func splitByOutermostCommas(input string) []string {
	var split []string
	var buffer []byte

	if len(input) == 0 {
		return []string{}
	}

	nestLevel := 0
	input = input[1 : len(input)-1]
	for i := 0; i < len(input); i++ {
		if input[i] == '[' {
			nestLevel += 1
		}
		if input[i] == ']' {
			nestLevel -= 1
		}
		if nestLevel == 0 && input[i] == ',' {
			split = append(split, string(buffer))
			buffer = []byte{}
		} else {
			buffer = append(buffer, input[i])
		}
	}
	split = append(split, string(buffer))

	return split
}

// [[1],[2,3,4]] vs [[1],4]
func compareStrings(a string, b string) int {
	partsA := splitByOutermostCommas(a)
	partsB := splitByOutermostCommas(b)

	for i := 0; i < min(len(partsA), len(partsB)); i++ {
		if len(partsA[i]) == 0 && len(partsB[i]) == 0 {
			return 1
		}
		if len(partsA[i]) == 0 {
			return 2
		}
		if len(partsB[i]) == 0 {
			return 0
		}

		if partsA[i][0] != '[' && partsB[i][0] != '[' {
			retVal := compareInts(partsA[i], partsB[i])
			if retVal != 1 {
				return retVal
			}
			continue
		}
		if partsA[i][0] == '[' && partsB[i][0] != '[' {
			partsB[i] = fmt.Sprintf("[%s]", partsB[i])
		}
		if partsA[i][0] != '[' && partsB[i][0] == '[' {
			partsA[i] = fmt.Sprintf("[%s]", partsA[i])
		}

		cmpSubResult := compareStrings(partsA[i], partsB[i])
		if cmpSubResult != 1 {
			return cmpSubResult
		}
	}

	if len(partsA) < len(partsB) {
		return 2
	}
	if len(partsA) > len(partsB) {
		return 0
	}
	return 1
}

func compareInts(a string, b string) int {
	x, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	y, err := strconv.Atoi(b)
	if err != nil {
		fmt.Println(err)
		return 0
	}

	if x > y {
		return 0
	}
	if x == y {
		return 1
	}

	return 2
}

func readInput(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	prevStr := ""
	index := 0
	sum := 0
	for scanner.Scan() {
		index += 1
		row := scanner.Text()

		if index%3 == 1 {
			prevStr = row
		}
		if index%3 == 2 {
			result := compareStrings(prevStr, row)
			if result > 0 {
				fmt.Println(result, (index+1)/3)
				sum += (index + 1) / 3
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return sum
}

func main() {
	//split := splitByOutermostCommas("[[1],[2,3,4]]")
	//split := splitByOutermostCommas("[1,[2,[3,[4,[5,6,7]]]],8,9]")
	//fmt.Println(split)
	sum := readInput("input2")
	fmt.Println(sum)
}
