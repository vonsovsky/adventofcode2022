package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var __visited map[int]bool

func __Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func __tailToHeadOneStep(hx int, hy int, tx int, ty int) (int, int) {
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

func __visualize(hx int, hy int, tx int, ty int, height int, width int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if hx == j && hy == i {
				fmt.Printf("H ")
			} else if tx == j && ty == i {
				fmt.Printf("T ")
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
}

func __move(hx *int, hy *int, tx *int, ty *int, direction string, steps int) {
	for i := 0; i < steps; i++ {
		if direction == "R" {
			*hx += 1
		}
		if direction == "L" {
			*hx -= 1
		}
		if direction == "U" {
			*hy -= 1
		}
		if direction == "D" {
			*hy += 1
		}
		*tx, *ty = tailToHeadOneStep(*hx, *hy, *tx, *ty)
		visited[(*tx)*1000+(*ty)] = true
		fmt.Println(*hx, *hy, *tx, *ty)
		//visualize(*hx, *hy, *tx, *ty, 6, 6)
		//fmt.Println()
	}
}

func __readInput(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var hx int = 500
	var hy int = 500
	var tx int = 500
	var ty int = 500

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		splits := strings.Split(row, " ")
		steps, err := strconv.Atoi(splits[1])
		if err != nil {
			log.Fatal(err)
		}
		move(&hx, &hy, &tx, &ty, splits[0], steps)
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
