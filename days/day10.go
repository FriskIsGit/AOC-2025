package days

import (
	"aoc-2025/util"
	"strings"
)

func Day10Part1(lines []string) int {
	machines := parseMachines(lines)
	presses := 0
	for _, m := range machines {
		_ = m
	}
	return presses
}

func parseMachines(lines []string) []Machine {
	machines := make([]Machine, 0, len(lines))
	for _, line := range lines {
		lineSplit := strings.Split(line, " ")
		rawRequirements := util.TrimLeftAndRight(lineSplit[len(lineSplit)-1], "{", "}")
		requirements := make([]int, 0)
		for _, req := range strings.Split(rawRequirements, ",") {
			val, _ := util.ParseInt(req)
			requirements = append(requirements, val)
		}
		buttons := make([][]int, 0)
		for i := 1; i < len(lineSplit)-1; i++ {
			button := make([]int, 0)
			rawButton := util.TrimLeftAndRight(lineSplit[i], "(", ")")
			for _, rawElement := range strings.Split(rawButton, ",") {
				val, _ := util.ParseInt(rawElement)
				button = append(button, val)
			}
			buttons = append(buttons, button)
		}
		rawLights := util.TrimLeftAndRight(lineSplit[0], "[", "]")
		lights := make([]bool, 0, len(rawLights))
		for _, light := range rawLights {
			if light == '#' {
				lights = append(lights, true)
			} else {
				lights = append(lights, false)
			}
		}
		machine := Machine{
			targetLights: lights,
			buttons:      buttons,
			requirements: requirements,
		}
		machines = append(machines, machine)
	}
	return machines
}

type Machine struct {
	targetLights []bool
	buttons      [][]int
	requirements []int
}

func Press(lights []bool, button []int) {
	for _, lightIndex := range button {
		lights[lightIndex] = !lights[lightIndex]
	}
}

func Compile(button1 []int, button2 []int) []int {
	compiled := make([]int, 0)
	for _, b1 := range button1 {
		unique := true
		for _, b2 := range button2 {
			if b1 == b2 {
				unique = false
				break
			}
		}
		if unique {
			compiled = append(compiled, b1)
		}
	}
	return compiled
}
