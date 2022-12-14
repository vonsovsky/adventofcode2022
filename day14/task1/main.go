package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type pos struct {
	x int
	y int
}

const (
	MinimumX = 450
	MaximumX = 510
	MinimumY = 0
	MaximumY = 170
)

func min(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

func convertPosition(posStr string) pos {
	posStr = strings.TrimSpace(posStr)
	parts := strings.Split(posStr, ",")
	x, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	y, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}

	return pos{x: x, y: y}
}

func parseLine(grid *[][]uint8, line string) {
	parts := strings.Split(line, "->")
	for k := 0; k < len(parts)-1; k++ {
		posFrom := convertPosition(parts[k])
		posTo := convertPosition(parts[k+1])
		for i := min(posFrom.y, posTo.y); i <= max(posFrom.y, posTo.y); i++ {
			for j := min(posFrom.x, posTo.x); j <= max(posFrom.x, posTo.x); j++ {
				(*grid)[i-MinimumY][j-MinimumX] = '#'
			}
		}
	}
}

func parseLineNoSaving(line string) []pos {
	parts := strings.Split(line, "->")
	var positions []pos
	for k := 0; k < len(parts); k++ {
		positions = append(positions, convertPosition(parts[k]))
	}
	return positions
}

func makeGrid(height int, width int) [][]uint8 {
	path := make([][]uint8, height)
	for i := range path {
		path[i] = make([]uint8, width)
		for j := 0; j < width; j++ {
			path[i][j] = '.'
		}
	}
	return path
}

func visualize(grid [][]uint8, height int, width int) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Printf("%c ", grid[i][j])
		}
		fmt.Println()
	}
}

func isInfiniteFall(grid [][]uint8, obj pos) bool {
	for i := obj.y + 1; i < len(grid); i++ {
		if grid[i][obj.x] != '.' {
			return false
		}
	}
	return true
}

func fallOneStep(grid [][]uint8, obj *pos) bool {
	if grid[(*obj).y+1][(*obj).x] == '.' {
		(*obj).y += 1
		return false
	}
	if grid[(*obj).y+1][(*obj).x-1] == '.' {
		(*obj).x -= 1
		(*obj).y += 1
		return false
	}
	if grid[(*obj).y+1][(*obj).x+1] == '.' {
		(*obj).x += 1
		(*obj).y += 1
		return false
	}
	return true
}

func fallOneSand(grid *[][]uint8, obj *pos) bool {
	done := false
	isInfinite := false

	for !done {
		(*grid)[obj.y][obj.x] = '.'
		done = fallOneStep(*grid, obj)
		(*grid)[obj.y][obj.x] = 'o'
		if isInfiniteFall(*grid, *obj) {
			done = true
			isInfinite = true
		}
	}

	return isInfinite
}

func fallAllTheSand(grid *[][]uint8) int {
	isInf := false
	index := 0
	for !isInf {
		isInf = fallOneSand(grid, &pos{500 - MinimumX, 0 - MinimumY})
		//fmt.Println(isInf)
		//visualize(*grid, MaximumY-MinimumY, MaximumX-MinimumX)
		index += 1
	}
	visualize(*grid, MaximumY-MinimumY, MaximumX-MinimumX)
	return index - 1
}

func readInput(fileName string) [][]uint8 {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := makeGrid(MaximumY-MinimumY+1, MaximumX-MinimumX+1)

	scanner := bufio.NewScanner(file)
	/*
		minX := 900
		maxX := -1
		minY := 900
		maxY := -1
		for scanner.Scan() {
			row := scanner.Text()
			positions := parseLineNoSaving(row)
			for _, ppos := range positions {
				if ppos.x < minX {
					minX = ppos.x
				}
				if ppos.x > maxX {
					maxX = ppos.x
				}
				if ppos.y < minY {
					minY = ppos.y
				}
				if ppos.y > maxY {
					maxY = ppos.y
				}
			}
		}

		fmt.Println(minX)
		fmt.Println(maxX)
		fmt.Println(minY)
		fmt.Println(maxY)
		return [][]uint8{}

		file.Seek(0, 0)
	*/

	for scanner.Scan() {
		row := scanner.Text()
		parseLine(&grid, row)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}

func main() {
	grid := readInput("input2")
	visualize(grid, MaximumY-MinimumY, MaximumX-MinimumX)
	steps := fallAllTheSand(&grid)
	fmt.Println(steps)
}
