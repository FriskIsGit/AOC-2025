package days

import (
	"aoc-2025/util"
	"slices"
	"strings"
)

func Day10Part1(lines []string) int {
	machines := parseMachines(lines)
	presses := 0
	for _, m := range machines {
		targetButton := m.getTargetButton()
		minPresses := len(m.buttons)

		// Generate combinations
	combinations:
		for p := 1; p < len(m.buttons); p++ {
			combInstance := util.NewCombinations(len(m.buttons), p)
			for combInstance.NextCombination() {
				buttonIndexes := combInstance.Indexes
				result := m.buttons[buttonIndexes[0]]
				for i := 1; i < p; i++ {
					result = ButtonResult(result, m.buttons[buttonIndexes[i]])
				}
				if slices.Equal(result, targetButton) {
					if p < minPresses {
						minPresses = p
						break combinations
					}
				}
			}
		}
		presses += minPresses
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

func (m Machine) getTargetButton() []int {
	targetButton := make([]int, 0, len(m.buttons))
	for i, on := range m.targetLights {
		if on {
			targetButton = append(targetButton, i)
		}
	}
	return targetButton
}

// Size of this buffer is decided by max light count
var lightBuffer = make([]bool, 16)

func ButtonResult(button1 []int, button2 []int) []int {
	util.MemZeroBoolArray(lightBuffer)
	result := make([]int, 0)
	for _, lightIndex := range button1 {
		lightBuffer[lightIndex] = !lightBuffer[lightIndex]
	}
	for _, lightIndex := range button2 {
		lightBuffer[lightIndex] = !lightBuffer[lightIndex]
	}
	for i, on := range lightBuffer {
		if on {
			result = append(result, i)
		}
	}
	return result
}

var FACTORIALS = []int64{
	1, 1, 2, 6, 24, 120, 720, 5040, 40320, 362880, 3628800, 39916800, 479001600, 6227020800, 87178291200, 1307674368000}
