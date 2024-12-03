package day2

import "fmt"

type Direction int

const (
	NONE Direction = iota
	INCREASING
	DECREASING
)

func ReportSafety(report [][]uint32, problemDampener int16) []bool {
	result := make([]bool, len(report))

	for i, row := range report {
		fmt.Printf("ROW %v \n", i)
		result[i] = checkRowSafety(row, problemDampener)
	}

	return result
}

func CountSafeReports(reports []bool) int {
	counter := 0
	for _, r := range reports {
		if r == true {
			counter++
		}
	}
	return counter
}

func checkRowSafety(row []uint32, problemDampener int16) bool {
	fmt.Printf("checkRowSafety (%d) ->  %d \n", problemDampener, row)
	direction := getDirection(row[0], row[1])
	for ri, v := range row {
		if ri > 0 {
			if !checkSafety(row[ri-1], v, direction) {

				if problemDampener == 0 {
					return false
				}

				for i2, _ := range row {
					if checkRowSafety(removeElement(row, i2), 0) {
						return true
					}
				}
				return false
			}
		}
	}
	return true
}

func checkSafety(v1 uint32, v2 uint32, currentDirection Direction) bool {
	fmt.Println("checkSafety -> v1:", v1, "v2:", v2, "currentDirection", currentDirection)
	levelDiff := max(v1, v2) - min(v1, v2)
	if levelDiff == 0 || levelDiff > 3 {
		fmt.Println("UNSAFE DIFF", levelDiff)
		return false
	}

	newDirection := getDirection(v1, v2)

	if newDirection != currentDirection {
		fmt.Println("UNSAFE DIRECTION", newDirection)
		return false
	}

	return true
}

func getDirection(v1 uint32, v2 uint32) Direction {
	if v1 > v2 {
		return INCREASING
	}
	if v1 < v2 {
		return DECREASING
	}
	return NONE
}

func removeElement(slice []uint32, index int) []uint32 {
	tmp := make([]uint32, len(slice))
	copy(tmp, slice)
	return append(tmp[:index], tmp[index+1:]...)
}
