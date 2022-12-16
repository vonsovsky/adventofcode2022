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

func coverRow(beacons []beacon, sensors []sensor, y int, endX int) (int, int) {
	ResetSegments()
	for _, b := range beacons {
		if b.y == y && b.x >= 0 && b.x <= endX {
			AddSegment(b.x, b.x)
		}
	}

	for _, s := range sensors {
		if s.y == y && s.x >= 0 && s.x <= endX {
			AddSegment(s.x, s.x)
		}

		distY := Abs(s.y - y)
		wingWidth := s.bdist - distY
		if wingWidth < 1 {
			continue
		}
		AddSegment(max(0, s.x-wingWidth), min(s.x+wingWidth, endX))
	}

	if len(GetSegments()) > 1 {
		fmt.Println(GetSegments())
		x := GetSegments()[0].end + 1
		return x, y
	}

	return -1, -1
}

func TestAddSegment() {
	AddSegment(5, 10)
	AddSegment(1, 2)
	AddSegment(9, 16)
	AddSegment(2, 6)
	fmt.Println(GetSegments())
	ResetSegments()

	AddSegment(3, 4)
	AddSegment(1, 2)
	AddSegment(5, 6)
	fmt.Println(GetSegments())
	AddSegment(1, 4)
	fmt.Println(GetSegments())
	ResetSegments()

	AddSegment(3, 4)
	AddSegment(1, 2)
	AddSegment(5, 6)
	fmt.Println(GetSegments())
	AddSegment(1, 6)
	fmt.Println(GetSegments())
	ResetSegments()
}

func main() {
	//TestAddSegment()
	//return

	dim := 4000000
	beacons, sensors := readInput("input2")
	for i := 0; i <= dim; i++ {
		x, y := coverRow(beacons, sensors, i, dim)
		if x >= 0 || y >= 0 {
			fmt.Println(x, y, x*4000000+y)
		}
	}
	return

	/*dim := 4000000
	beacons, sensors := readInput("input2")
	for i := 0; i <= dim; i++ {
		if i%(dim/100) == 0 {
			fmt.Printf("Progress: %d%% complete", i/(dim/100))
		}
		x, y := coverRow(beacons, sensors, i, dim)
		if x >= 0 || y >= 0 {
			fmt.Println(i, x, y)
		}
	}*/
}
