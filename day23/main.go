package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	y int
	x int
}

type proposed struct {
	elves []pos
	count int
}

func visualize(elves map[pos]bool, startX int, endX int, startY int, endY int) {
	for y := startY; y <= endY; y++ {
		for x := startX; x <= endX; x++ {
			if elves[pos{x: x, y: y}] {
				fmt.Print("# ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func freeFields(elves map[pos]bool, startX int, endX int, startY int, endY int) bool {
	free := true
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			if elves[pos{x: x, y: y}] {
				return false
			}
		}
	}
	return free
}

func gatherProposedMoves(elves map[pos]bool, index int) map[pos]proposed {
	moves := make(map[pos]proposed)
	order := []int{0, 1, 2, 3}

	for elf, _ := range elves {
		var addMove pos
		added := false
		if freeFields(elves, elf.x-1, elf.x+1, elf.y-1, elf.y-1) &&
			freeFields(elves, elf.x-1, elf.x+1, elf.y+1, elf.y+1) &&
			freeFields(elves, elf.x-1, elf.x-1, elf.y, elf.y) &&
			freeFields(elves, elf.x+1, elf.x+1, elf.y, elf.y) {
			continue
		}

		for i := 0; i < 4; i++ {
			pointer := order[(i+index)%4]
			if pointer == 0 && freeFields(elves, elf.x-1, elf.x+1, elf.y-1, elf.y-1) {
				addMove = pos{x: elf.x, y: elf.y - 1}
				added = true
				break
			}
			if pointer == 1 && freeFields(elves, elf.x-1, elf.x+1, elf.y+1, elf.y+1) {
				addMove = pos{x: elf.x, y: elf.y + 1}
				added = true
				break
			}
			if pointer == 2 && freeFields(elves, elf.x-1, elf.x-1, elf.y-1, elf.y+1) {
				addMove = pos{x: elf.x - 1, y: elf.y}
				added = true
				break
			}
			if pointer == 3 && freeFields(elves, elf.x+1, elf.x+1, elf.y-1, elf.y+1) {
				addMove = pos{x: elf.x + 1, y: elf.y}
				added = true
				break
			}
		}

		if !added {
			continue
		}

		propElves := moves[addMove]
		propElves.elves = append(propElves.elves, elf)
		propElves.count += 1
		moves[addMove] = propElves
	}
	return moves
}

func moveElves(elves *map[pos]bool, proposedMoves map[pos]proposed) {
	for moveTo, propElves := range proposedMoves {
		if propElves.count == 1 {
			for _, elf := range propElves.elves {
				delete(*elves, elf)
				(*elves)[moveTo] = true
				//fmt.Printf("Elf %d, %d, will move to %d, %d\n",
				//	elf.y+1, elf.x+1, moveTo.y+1, moveTo.x+1)
			}
		}
	}
}

func parseLine(elves *map[pos]bool, row string, y int) {
	for index, chr := range row {
		if chr == '#' {
			elf := pos{x: index, y: y}
			(*elves)[elf] = true
		}
	}
}

func calculateScore(elves map[pos]bool) int {
	minX := 999
	maxX := 0
	minY := 999
	maxY := 0
	for elf, _ := range elves {
		if elf.x < minX {
			minX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}

	return (maxX-minX+1)*(maxY-minY+1) - len(elves)
}

func readInput(fileName string) map[pos]bool {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	elves := make(map[pos]bool)

	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		row := scanner.Text()
		parseLine(&elves, row, index)
		index += 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elves
}

func main() {
	elves := readInput("input3")

	for i := 0; i < 3000; i++ {
		proposedMoves := gatherProposedMoves(elves, i)
		if len(proposedMoves) == 0 {
			print(i + 1)
			break
		}
		moveElves(&elves, proposedMoves)
		//visualize(elves, 0, 14, 0, 14)
		//fmt.Println()
	}

	//score := calculateScore(elves)
	//fmt.Println(score)
}
