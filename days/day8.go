package days

import (
	"aoc-2025/util"
	"fmt"
	"math"
	"slices"
	"strings"
)

// --- Day 8: Playground ---

type Point3D struct {
	x, y, z int
}

func (p Point3D) Hash() int64 {
	return int64(p.x << 42 & p.y << 21 & p.z)
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.x, p.y, p.z)
}

func NewPoint3D(x, y, z int) Point3D {
	return Point3D{x, y, z}
}

func (p Point3D) Equal(other *Point3D) bool {
	return p.x == other.x && p.y == other.y && p.z == other.z
}

func (p Point3D) Distance(other *Point3D) float64 {
	xDelta := float64(p.x - other.x)
	yDelta := float64(p.y - other.y)
	zDelta := float64(p.z - other.z)
	return math.Sqrt(xDelta*xDelta + yDelta*yDelta + zDelta*zDelta)
}

func Distance2D(x, y, x1, y1 int) float64 {
	xDelta := float64(x - x1)
	yDelta := float64(y - y1)
	return math.Sqrt(xDelta*xDelta + yDelta*yDelta)
}

type BoxPair struct {
	p1, p2   Point3D
	distance float64
}

func Day8Part1(lines []string) int {
	points := make([]Point3D, 0, len(lines))
	// Max coordinate is approx. 100_000
	for _, line := range lines {
		terms := strings.Split(line, ",")
		x, _ := util.ParseInt(terms[0])
		y, _ := util.ParseInt(terms[1])
		z, _ := util.ParseInt(terms[2])
		point := NewPoint3D(x, y, z)
		points = append(points, point)
	}

	closestPairs := make([]BoxPair, 0)
	for i := 0; i < len(points); i++ {
		box := points[i]
		for j := i + 1; j < len(points); j++ {
			otherBox := points[j]
			distance := box.Distance(&otherBox)
			boxPair := BoxPair{p1: box, p2: otherBox, distance: distance}
			closestPairs = append(closestPairs, boxPair)
		}
	}
	slices.SortFunc(closestPairs, func(a, b BoxPair) int {
		if a.distance < b.distance {
			return -1
		}
		if a.distance > b.distance {
			return 1
		}
		return 0
	})

	circuits := make([][]Point3D, 0)
	for pairIndex, pair := range closestPairs {
		if pairIndex == 1000 {
			break
		}
		var p1Circuit, p2Circuit []Point3D
		p1Index, p2Index := -1, -1
		for circuitIndex, circuit := range circuits {
			for _, point := range circuit {
				if point.Equal(&pair.p1) {
					p1Circuit = circuit
					p1Index = circuitIndex
					continue
				}
				if point.Equal(&pair.p2) {
					p2Circuit = circuit
					p2Index = circuitIndex
				}
			}
		}
		if p1Circuit != nil && p2Circuit != nil {
			// Both points have their own circuits or are in the same circuit, therefore they cannot be connected directly
			continue
		}
		if p1Circuit == nil && p2Circuit == nil {
			// Create new circuit
			var newCircuit []Point3D
			newCircuit = append(newCircuit, pair.p1, pair.p2)
			circuits = append(circuits, newCircuit)
			continue
		}
		// Append to one of the existing circuits
		if p1Circuit != nil {
			circuits[p1Index] = append(p1Circuit, pair.p2)
		} else {
			circuits[p2Index] = append(p2Circuit, pair.p1)
		}
	}
	if len(circuits) < 3 {
		fmt.Println("Not enough circuits. Expected at least 3")
		return -1
	}
	fmt.Printf("Circuit count: %v\n", len(circuits))
	printCircuits(circuits)
	product := 1
	circuitSizes := make([]int, len(circuits))
	for _, circuit := range circuits {
		circuitSizes = append(circuitSizes, len(circuit))
	}
	threeLargest := util.LargestN(circuitSizes, 3)
	for _, large := range threeLargest {
		product *= large
	}
	return product
}

func printCircuits(circuits [][]Point3D) {
	for _, circuit := range circuits {
		for _, point := range circuit {
			fmt.Printf("%v,", point)
		}
		fmt.Println()
	}
}
