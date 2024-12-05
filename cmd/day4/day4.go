package day4

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func CountXmas(matrix *[][]rune) uint32 {
	var counter uint32 = 0
	for ir, row := range *matrix {
		for ic, col := range row {

			if col == 'X' {
				connections := getMatrixGraphConnections(ir, ic, len(*matrix), len(row))
				counter += searchMatrixConnections(connections, nil, []rune{'M', 'A', 'S'}, matrix)
			}
		}
	}
	return counter
}

func CountX_mas(matrix *[][]rune) uint32 {
	var counter uint32 = 0
	for ir, row := range *matrix {
		for ic, col := range row {

			is_X_mas := 0
			if col == 'A' {
				connections := getMatrixGraphConnections(ir, ic, len(*matrix), len(row))
				// fmt.Printf("%v\n", connections)
				for _, conn := range connections {
					if conn.direction == "←" || conn.direction == "↑" || conn.direction == "→" || conn.direction == "↓" {
						continue
					}
					value := (*matrix)[conn.coordinates[0]][conn.coordinates[1]]
					// fmt.Printf("%v\n", strconv.QuoteRune(value))
					if value != 'M' {
						continue
					}

					found := slices.IndexFunc(connections, func(c Connection) bool { return c.direction == getInverseDirection(conn.direction) })
					// fmt.Printf("%v\n", found)
					if found < 0 {
						continue
					}
					valorContrario := (*matrix)[connections[found].coordinates[0]][connections[found].coordinates[1]]
					// fmt.Printf("v:%v contrario: %v\n", strconv.QuoteRune(value), strconv.QuoteRune(valorContrario))
					if valorContrario == 'S' {
						is_X_mas++
					}
				}
			}
			if is_X_mas == 2 {
				counter++
			}
		}
	}
	return counter
}

type Connection struct {
	direction   string
	coordinates [2]uint32
}

func getInverseDirection(direction string) string {
	directions := map[string]string{
		"←": "→",
		"↖": "↘",
		"↑": "↓",
		"↗": "↙",
		"→": "←",
		"↘": "↖",
		"↓": "↑",
		"↙": "↗",
	}
	return directions[direction]
}

func getMatrixGraphConnections(positionX int, positionY int, totalRows int, totalColumns int) (connections []Connection) {
	// "← ↖ ↑ ↗ → ↘ ↓ ↙"
	var result []Connection

	for row := int32(positionX) - 1; row <= int32(positionX)+1; row++ {
		// fmt.Printf("row %d\n", row)
		for column := int32(positionY) - 1; column <= int32(positionY)+1; column++ {
			// fmt.Printf("column %d\n", column)
			//check if outbounds
			if row > -1 && column > -1 && row < int32(totalRows) && column < int32(totalColumns) {

				direction := ""
				switch {
				case row > int32(positionX):
					switch {
					case column < int32(positionY):
						direction = "↙"
					case column == int32(positionY):
						direction = "↓"
					case column > int32(positionY):
						direction = "↘"
					}
				case row == int32(positionX):
					switch {
					case column < int32(positionY):
						direction = "←"
					case column == int32(positionY):
						direction = "???"
					case column > int32(positionY):
						direction = "→"
					}
				case row < int32(positionX):
					switch {
					case column < int32(positionY):
						direction = "↖"
					case column == int32(positionY):
						direction = "↑"
					case column > int32(positionY):
						direction = "↗"
					}
				}

				//check if its not itself
				if row != int32(positionX) || column != int32(positionY) {
					// fmt.Printf("Vai fazer append [%v,%v]\n", row, column)
					result = append(result, Connection{direction: direction, coordinates: [2]uint32{uint32(row), uint32(column)}})
				}
			}
		}
	}
	return result
}

func searchMatrixConnections(connections []Connection, direction *string, searchTerm []rune, matrix *[][]rune) uint32 {
	fmt.Printf("searchMatrixConnections: %v \n", string(searchTerm))
	if len(searchTerm) == 0 {
		fmt.Println("JACKPOT!")
		return 1
	}
	var matches uint32 = 0

	for _, conn := range connections {

		if direction == nil || conn.direction == *direction {
			// fmt.Printf("%v %v", conn.direction, *direction)
			connValue := (*matrix)[conn.coordinates[0]][conn.coordinates[1]]
			if connValue == searchTerm[0] {
				fmt.Printf("Found [%v] direction %v %v\n", strconv.QuoteRune(searchTerm[0]), conn.direction, searchTerm)
				childConnections := getMatrixGraphConnections(int(conn.coordinates[0]), int(conn.coordinates[1]), len(*matrix), len((*matrix)[0]))
				matches += searchMatrixConnections(childConnections, &conn.direction, searchTerm[1:], matrix)
			}
		}
	}
	return matches
}

func BuildMatrixGraph(text string) [][]rune {
	input := [][]rune{}
	for _, line := range strings.Split(text, "\n") {
		runesLine := make([]rune, len(line))
		for _, char := range line {
			runesLine = append(runesLine, char)
		}
		input = append(input, runesLine)
	}
	return input
}
