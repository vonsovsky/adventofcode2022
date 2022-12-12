package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type pos struct {
	x int
	y int
}

func readInput(fileName string) [][]uint8 {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]uint8

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		gridLine := make([]uint8, len(row))
		for i, ch := range row {
			gridLine[i] = uint8(ch)
		}
		grid = append(grid, gridLine)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return grid
}

func makePathGrid(height int, width int) [][]int {
	path := make([][]int, height)
	for i := range path {
		path[i] = make([]int, width)
	}
	return path
}

func makeCharMap(height int, width int) [][]byte {
	path := make([][]byte, height)
	for i := range path {
		path[i] = make([]byte, width)
		for j := 0; j < width; j++ {
			path[i][j] = '.'
		}
	}
	return path
}

func findStartEnd(grid [][]uint8) (pos, pos) {
	s := pos{x: -1, y: -1}
	e := pos{x: -1, y: -1}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				s.x = j
				s.y = i
				grid[i][j] = 'a'
			}
			if grid[i][j] == 'E' {
				e.x = j
				e.y = i
				grid[i][j] = 'z'
			}
		}
	}

	return s, e
}

func addNode(grid [][]uint8, path *[][]int, charMap *[][]byte, queue *[]pos, fromx int, fromy int, tox int, toy int, char byte) {
	h := grid[fromy][fromx]

	if tox < 0 || tox >= len(grid[0]) {
		return
	}
	if toy < 0 || toy >= len(grid) {
		return
	}

	//if Abs(grid[toy][tox]-h) <= 1 && (*path)[toy][tox] == 0 {
	if grid[toy][tox] <= h+1 && (*path)[toy][tox] == 0 {
		(*path)[toy][tox] = (*path)[fromy][fromx] + 1
		(*charMap)[toy][tox] = char
		*queue = append(*queue, pos{x: tox, y: toy})
	}
}

func bfs(grid [][]uint8, path *[][]int, charMap *[][]byte, s pos, e pos) int {
	indexPointer := 0
	queue := make([]pos, 1)
	queue[0] = pos{x: s.x, y: s.y}
	for indexPointer < len(queue) {
		node := queue[indexPointer]
		if node.x == e.x && node.y == e.y {
			return (*path)[node.y][node.x] - 1
		}
		addNode(grid, path, charMap, &queue, node.x, node.y, node.x+1, node.y, '>')
		addNode(grid, path, charMap, &queue, node.x, node.y, node.x-1, node.y, '<')
		addNode(grid, path, charMap, &queue, node.x, node.y, node.x, node.y+1, 'v')
		addNode(grid, path, charMap, &queue, node.x, node.y, node.x, node.y-1, '^')
		indexPointer += 1
	}
	return -1
}

func main() {
	grid := readInput("input2")
	path := makePathGrid(len(grid), len(grid[0]))
	charMap := makeCharMap(len(grid), len(grid[0]))
	s, e := findStartEnd(grid)
	path[s.y][s.x] = 1
	dist := bfs(grid, &path, &charMap, s, e)
	fmt.Println(dist)

	/*
		for _, chLine := range charMap {
			for _, chr := range chLine {
				fmt.Printf("%s ", string(chr))
			}
			fmt.Println()
		}

		f, _ := os.Create("output.txt")
		defer f.Close()

		for i, chLine := range path {
			for j, chr := range chLine {
				f.WriteString(fmt.Sprintf("%s", string(grid[i][j])))
				if chr < 10 {
					f.WriteString("00")
				} else if chr < 100 {
					f.WriteString("0")
				}
				f.WriteString(fmt.Sprintf("%d ", chr))
			}
			f.WriteString("\n")
		}
	*/
}
