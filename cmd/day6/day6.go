package day6

import (
	"fmt"
	"strconv"
)

func CountSteps(_map [][]rune) int {
	initialPosition := [2]int{0, 0}
	for ri, row := range _map {
		for ci, column := range row {
			if column == '^' {
				initialPosition[0] = ri
				initialPosition[1] = ci
			}
		}
	}

	currentPosition := initialPosition
	for {
		front := getInFront(&_map, currentPosition)
		if isOutOfBounds(&_map, front) {
			_map[currentPosition[0]][currentPosition[1]] = 'X'
			break
		}
		fmt.Printf("%v -> %v (%v) \n", strconv.QuoteRune(_map[currentPosition[0]][currentPosition[1]]), strconv.QuoteRune(_map[front[0]][front[1]]), front)
		if _map[front[0]][front[1]] == '#' {
			turnRight(&_map, currentPosition)
			continue
		} else {
			currentPosition = moveFoward(&_map, currentPosition)
			continue
		}
	}

	countX := 0
	for _, row := range _map {
		for _, c := range row {
			fmt.Print(strconv.QuoteRune(c))
			if c == 'X' {
				countX++
			}
		}
		fmt.Print("\n")
	}
	return countX
}

func getInFront(_map *[][]rune, currentPosition [2]int) [2]int {
	switch (*_map)[currentPosition[0]][currentPosition[1]] {
	case '^':
		return [2]int{currentPosition[0] - 1, currentPosition[1]}
	case 'v':
		return [2]int{currentPosition[0] + 1, currentPosition[1]}
	case '>':
		return [2]int{currentPosition[0], currentPosition[1] + 1}
	case '<':
		return [2]int{currentPosition[0], currentPosition[1] - 1}
	}
	panic("Invalid current position")
}

func isOutOfBounds(_map *[][]rune, position [2]int) bool {
	return (position[0] >= len(*_map) || position[0] >= len((*_map)[0])) || position[0] < 0 || position[1] < 0
}

func turnRight(_map *[][]rune, currentPosition [2]int) {
	switch (*_map)[currentPosition[0]][currentPosition[1]] {
	case '^':
		(*_map)[currentPosition[0]][currentPosition[1]] = '>'
	case 'v':
		(*_map)[currentPosition[0]][currentPosition[1]] = '<'
	case '>':
		(*_map)[currentPosition[0]][currentPosition[1]] = 'v'
	case '<':
		(*_map)[currentPosition[0]][currentPosition[1]] = '^'
	}
}

func moveFoward(_map *[][]rune, currentPosition [2]int) [2]int {
	var newPosition [2]int
	switch (*_map)[currentPosition[0]][currentPosition[1]] {
	case '^':
		newPosition = [2]int{currentPosition[0] - 1, currentPosition[1]}
	case 'v':
		newPosition = [2]int{currentPosition[0] + 1, currentPosition[1]}
	case '>':
		newPosition = [2]int{currentPosition[0], currentPosition[1] + 1}
	case '<':
		newPosition = [2]int{currentPosition[0], currentPosition[1] - 1}
	}

	(*_map)[newPosition[0]][newPosition[1]] = (*_map)[currentPosition[0]][currentPosition[1]]
	(*_map)[currentPosition[0]][currentPosition[1]] = 'X'

	return newPosition
}
