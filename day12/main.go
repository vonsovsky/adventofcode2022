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

func Abs(x uint8) uint8 {
	if x < 0 {
		return -x
	}
	return x
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

func findEnd(grid [][]uint8) pos {
	e := pos{x: -1, y: -1}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 'S' {
				grid[i][j] = 'a'
			}
			if grid[i][j] == 'E' {
				e.x = j
				e.y = i
				grid[i][j] = 'z'
			}
		}
	}

	return e
}

func addNode(grid [][]uint8, path *[][]int, queue *[]pos, fromx int, fromy int, tox int, toy int) {
	h := grid[fromy][fromx]

	if tox < 0 || tox >= len(grid[0]) {
		return
	}
	if toy < 0 || toy >= len(grid) {
		return
	}

	//if Abs(grid[toy][tox]-h) <= 1 && (*path)[toy][tox] == 0 {
	if h <= grid[toy][tox]+1 && (*path)[toy][tox] == 0 {
		(*path)[toy][tox] = (*path)[fromy][fromx] + 1
		*queue = append(*queue, pos{x: tox, y: toy})
	}
}

func bfs(grid [][]uint8, path *[][]int, s pos) int {
	indexPointer := 0
	queue := make([]pos, 1)
	queue[0] = pos{x: s.x, y: s.y}
	for indexPointer < len(queue) {
		node := queue[indexPointer]
		if grid[node.y][node.x] == 'a' {
			return (*path)[node.y][node.x] - 1
		}
		addNode(grid, path, &queue, node.x, node.y, node.x+1, node.y)
		addNode(grid, path, &queue, node.x, node.y, node.x-1, node.y)
		addNode(grid, path, &queue, node.x, node.y, node.x, node.y+1)
		addNode(grid, path, &queue, node.x, node.y, node.x, node.y-1)
		indexPointer += 1
	}
	return -1
}

func main() {
	grid := readInput("input1")
	path := makePathGrid(len(grid), len(grid[0]))
	e := findEnd(grid)
	path[e.y][e.x] = 1
	dist := bfs(grid, &path, e)
	fmt.Println(dist)
}
