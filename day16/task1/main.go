package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type valve struct {
	flow int
	valves []string
}

type action struct {
	pos1 string
	pos2 string
	action1 string
	action2 string
	opened map[string]bool
	recentOpen map[string]bool
	score int
	step int
}

func getRoundScore(valves map[string]valve, openValves []string) int {
	score := 0
	for _, valveName := range openValves {
		score += valves[valveName].flow
	}

	return score
}

func mapKeys(mp any) []string {
	keys := reflect.ValueOf(mp).MapKeys()
	strkeys := make([]string, len(keys))
	for i := 0; i < len(keys); i++ {
		strkeys[i] = keys[i].String()
	}
	return strkeys
}

func copyMap(mapToCopy map[string]bool) map[string]bool {
	newMap := make(map[string]bool)
	for k,v := range mapToCopy {
		newMap[k] = v
	}
	return newMap
}

func getGoPositions(valves map[string]valve, state action) []string {
	var actions []string
	for _, path := range valves[state.pos1].valves {
		actions = append(actions, path)
		actions = append(actions, "go")
	}

	return actions
}

func isProspective(expScore int, curScore int) bool {
	allowedMinimumScore := float64(expScore) * 0.9
	if curScore < int(allowedMinimumScore) {
		return false
	}
	return true
}

func goDFS(valves map[string]valve) int {
	var stack []action
	stack = append(stack, action{
		pos1: "AA",
		pos2: "AA",
		action1: "go",
		action2: "go",
		opened: map[string]bool{"AA": true},
		recentOpen: map[string]bool{},
		score: 0,
		step: 0,
	})

	highestScores := make([]int, 32)

	highestScore := 0
	for len(stack) > 0 {
		l := len(stack)
		state := stack[l-1]
		stack = stack[:l-1]

		if state.score > highestScores[state.step] {
			highestScores[state.step] = state.score
		}

		if state.step >= 31 {
			if state.score > highestScore {
				highestScore = state.score
			}
			continue
		}

		newScore := getRoundScore(valves, mapKeys(state.opened))

		opened := copyMap(state.opened)
		if len(state.recentOpen) > 0 {
			for k, v := range state.recentOpen {
				opened[k] = v
			}
			state.recentOpen = map[string]bool{}
		}

		var actions1 []string
		if state.action1 == "go" {
			actions1 = append(actions1, getGoPositions(valves, state)...)
			if !state.opened[state.pos1] {
				actions1 = append(actions1, state.pos1)
				actions1 = append(actions1, "open")
			}
		}

		for i := 0; i < len(actions1); i += 2 {
			if isProspective(highestScores[state.step+1], state.score + newScore) {
				stack = append(stack, action{
					pos1: actions1[i],
					action1: actions1[i+1],
					recentOpen: state.recentOpen,
					opened:     opened,
					score:      state.score + newScore,
					step:       state.step + 1,
				})
			}
		}

		if state.action1 == "open" || state.action2 == "open" {
			recent := copyMap(state.recentOpen)
			if state.action1 == "open" {
				recent[state.pos1] = true
			}
			stack = append(stack, action{
				pos1: state.pos1,
				pos2: state.pos2,
				action1: "go",
				action2: "go",
				opened: state.opened,
				recentOpen: recent,
				score: state.score,
				step: state.step,
			})
		}
	}

	return highestScore
}


func parseLine(line string, valves *map[string]valve) {
	rex := regexp.MustCompile(`Valve ([A-Z]+) has flow rate=(\d+); tunnels? leads? to valves? ([A-Z\s,]+)`)
	mt := rex.FindAllStringSubmatch(line, -1)
	valveName := mt[0][1]
	flow, _ := strconv.Atoi(mt[0][2])
	otherValves := strings.Split(strings.ReplaceAll(mt[0][3]," ","") , ",")

	(*valves)[valveName] = valve{flow: flow, valves: otherValves}
}

func readInput(fileName string) map[string]valve {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	valves := map[string]valve{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		row := scanner.Text()
		parseLine(row, &valves)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return valves
}

func main() {
	valves := readInput("input1")

	start := time.Now()
	highestScore := goDFS(valves)
	elapsed := time.Since(start) // running time ~90s

	fmt.Println(highestScore)
	fmt.Println("Time elapsed:", elapsed)
}
