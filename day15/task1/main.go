package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type beacon struct {
	x int
	y int
}

type sensor struct {
	x     int
	y     int
	bdist int
}

const (
	MinimumX = 0
	MaximumX = 6000000
	MinimumY = 0
	MaximumY = 4000000
	xMod     = 3000000
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

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func distance(x1 int, y1 int, x2 int, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

func parseLine(line string, beacons *[]beacon, sensors *[]sensor) {
	rex := regexp.MustCompile(`Sensor at x=([\d\-]+), y=([\d\-]+): closest beacon is at x=([\d\-]+), y=([\d\-]+)`)
	mt := rex.FindAllStringSubmatch(line, -1)
	sx, _ := strconv.Atoi(mt[0][1])
	sy, _ := strconv.Atoi(mt[0][2])
	bx, _ := strconv.Atoi(mt[0][3])
	by, _ := strconv.Atoi(mt[0][4])
	*beacons = append(*beacons, beacon{x: bx, y: by})
	*sensors = append(*sensors, sensor{x: sx, y: sy, bdist: distance(sx, sy, bx, by)})
}

func makeGrid(height int, width int) [][]uint8 {
	path := make([][]uint8, height)
	for i := range path {
		path[i] = make([]uint8, width)
		for j := 0; j < width; j++ {
			path[i][j] = '#'
		}
	}
	return path
}

func visualize(grid [][]uint8, height int, width int) {
	for i := 0; i < height; i++ {
		if i < 10 {
			fmt.Printf("0")
		}
		fmt.Printf("%d ", i)
		for j := 0; j < width; j++ {
			fmt.Printf("%c ", grid[i][j])
		}
		fmt.Println()
	}
}

func readInput(fileName string) ([]beacon, []sensor) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var beacons []beacon
	var sensors []sensor

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		parseLine(row, &beacons, &sensors)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return beacons, sensors
}

func isWithinSensorRange(beacons []beacon, sensors []sensor, x int, y int) bool {
	for _, s := range sensors {
		if s.x == x && s.y == y {
			return false
		}
	}
	for _, b := range beacons {
		if b.x == x && b.y == y {
			return false
		}
	}

	for _, s := range sensors {
		dist := distance(s.x, s.y, x, y)
		if dist <= s.bdist {
			return true
		}
	}
	return false
}

func main() {
	beacons, sensors := readInput("input2")
	countInRanges := 0
	for i := -5000000; i < 5000000; i++ {
		inRange := isWithinSensorRange(beacons, sensors, i, 2000000)
		if inRange {
			countInRanges += 1
		}
		//fmt.Print(i)
		//fmt.Println(inRange)
	}
	fmt.Println(countInRanges)
}
