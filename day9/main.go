package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var visited map[int]bool

type pos struct {
	x int
	y int
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func tailToHeadOneStep(hx int, hy int, tx int, ty int) (int, int) {
	if Abs(hx-tx) <= 1 && Abs(hy-ty) <= 1 {
		return tx, ty
	}
	if Abs(hx-tx) > 0 {
		if hx > tx {
			tx += 1
		} else {
			tx -= 1
		}
	}
	if Abs(hy-ty) > 0 {
		if hy > ty {
			ty += 1
		} else {
			ty -= 1
		}
	}
	return tx, ty
}

func visualize(rope []pos, height int, width int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if rope[0].x == j && rope[0].y == i {
				fmt.Printf("H ")
				continue
			}
			found := false
			for k := 1; k < 10; k++ {
				if found {
					break
				}
				for rope[k].x == j && rope[k].y == i {
					fmt.Printf("%d ", k)
					found = true
					break
				}
			}

			if found {
				continue
			}
			fmt.Printf(". ")
		}
		fmt.Println()
	}
}

func move(rope *[]pos, direction string, steps int) {
	for i := 0; i < steps; i++ {
		if direction == "R" {
			(*rope)[0].x += 1
		}
		if direction == "L" {
			(*rope)[0].x -= 1
		}
		if direction == "U" {
			(*rope)[0].y -= 1
		}
		if direction == "D" {
			(*rope)[0].y += 1
		}
		for k := 1; k < 10; k++ {
			(*rope)[k].x, (*rope)[k].y = tailToHeadOneStep((*rope)[k-1].x, (*rope)[k-1].y, (*rope)[k].x, (*rope)[k].y)
		}
		visited[(*rope)[9].x*1000+(*rope)[9].y] = true
		//fmt.Println(*hx, *hy, *tx, *ty)
		//visualize(*rope, 21, 26)
		//fmt.Println()
	}
}

func readInput(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rope := make([]pos, 10)
	for i := 0; i < 10; i++ {
		rope[i].x = 11
		rope[i].y = 15
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		splits := strings.Split(row, " ")
		steps, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatal(err)
		}
		move(&rope, splits[0], steps)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	visited = map[int]bool{}
	readInput("input2")
	fmt.Println(len(visited))
	//tx, ty := tailToHeadOneStep(1, 2, 3, 1)
	//fmt.Println(tx, ty)
	//visualize(2, 1, tx, ty, 5, 5)
}
