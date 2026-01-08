package days

import (
	"aoc-2025/util"
	"fmt"
	"math"
	"slices"
	"strings"
)

// --- Day 8: Playground ---
// Although not explicitly stated the number of shortest connections to be made depends on the input type:
// - demo = 10
// - full = 1000

type Point3D struct {
	X, Y, Z int
}

// Hash - max coordinate is approximately 100_000
func (p Point3D) Hash() int64 {
	return int64(p.X<<42 | p.Y<<21 | p.Z)
}

const MASK = int64((1 << 21) - 1)

func UnHash(hash int64) Point3D {
	x := (hash >> 42) & MASK
	y := (hash >> 21) & MASK
	z := hash & MASK
	return Point3D{int(x), int(y), int(z)}
}

func (p Point3D) String() string {
	return fmt.Sprintf("(%d,%d,%d)", p.X, p.Y, p.Z)
}

func NewPoint3D(x, y, z int) Point3D {
	return Point3D{x, y, z}
}

func (p Point3D) Equal(other *Point3D) bool {
	return p.X == other.X && p.Y == other.Y && p.Z == other.Z
}

func (p Point3D) Distance(other *Point3D) float64 {
	xDelta := float64(p.X - other.X)
	yDelta := float64(p.Y - other.Y)
	zDelta := float64(p.Z - other.Z)
	return math.Sqrt(xDelta*xDelta + yDelta*yDelta + zDelta*zDelta)
}

type BoxPair struct {
	p1, p2   Point3D
	distance float64
}

func Day8Part1(lines []string, limitTo int) int {
	points := make([]Point3D, 0, len(lines))
	for _, line := range lines {
		terms := strings.Split(line, ",")
		x, _ := util.ParseInt(terms[0])
		y, _ := util.ParseInt(terms[1])
		z, _ := util.ParseInt(terms[2])
		point := NewPoint3D(x, y, z)
		points = append(points, point)
	}

	closestPairs := make([]BoxPair, 0)
	for i := 0; i < len(points)-1; i++ {
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

	connectionsMade := 0
	circuits := make([]util.Set[int64], 0)
	for _, pair := range closestPairs {
		p1Index, p2Index := -1, -1
		for index, circuit := range circuits {
			if p1Index == -1 && circuit.Contains(pair.p1.Hash()) {
				p1Index = index
			}
			if p2Index == -1 && circuit.Contains(pair.p2.Hash()) {
				p2Index = index
			}
		}

		if p1Index == p2Index {
			if p1Index != -1 {
				// It's the same circuit, skip
				continue
			}
			// Neither has a circuit
			set := util.NewSet[int64](2)
			set.Add(pair.p1.Hash())
			set.Add(pair.p2.Hash())
			circuits = append(circuits, *set)
			connectionsMade++
			if connectionsMade >= limitTo {
				break
			}
			continue
		}

		mergeCircuits := p1Index != -1 && p2Index != -1
		if mergeCircuits {
			circuits[p1Index].AddAll(&circuits[p2Index])
			circuits = util.DeleteAt(circuits, p2Index)
		} else if p1Index == -1 {
			// P1 has no circuit, add it to P2's circuit
			circuits[p2Index].Add(pair.p1.Hash())
		} else {
			// P2 has no circuit, add it to P1's circuit
			circuits[p1Index].Add(pair.p2.Hash())
		}
		connectionsMade++
		if connectionsMade >= limitTo {
			break
		}
	}
	if len(circuits) < 3 {
		fmt.Println("Not enough circuits. Expected at least 3")
		return -1
	}
	fmt.Printf("Circuit count: %v\n", len(circuits))
	printCircuits(circuits, points, true)
	product := 1
	circuitSizes := make([]int, len(circuits))
	for _, circuit := range circuits {
		circuitSizes = append(circuitSizes, circuit.Size())
	}
	threeLargest := util.LargestN(circuitSizes, 3)
	for _, large := range threeLargest {
		product *= large
	}
	return product
}

func printCircuits(circuits []util.Set[int64], points []Point3D, byIndex bool) {
	for _, circuit := range circuits {
		for k := range circuit.Set {
			point := UnHash(k)
			if byIndex {
				index := inputIndexOf(point, points)
				fmt.Printf("%v,", index)
			} else {
				fmt.Printf("%v,", point)
			}
		}
		fmt.Println()
	}
}

func inputIndexOf(p Point3D, points []Point3D) int {
	for i, point := range points {
		if point.Equal(&p) {
			return i
		}
	}
	return -1
}
