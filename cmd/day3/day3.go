package day3

import (
	"regexp"
	"strconv"
)

func ClearInput(input *string, conditional bool) []string {
	pattern := `mul\(\d{0,3},\d{0,3}\)`
	if conditional {
		pattern = `(?s)(?:don't\(\).*?do\(\))|(mul\(\d{0,3},\d{0,3}\))`
	}
	r, _ := regexp.Compile(pattern)

	result := []string{}
	for _, matches := range r.FindAllStringSubmatch(*input, -1) {
		if conditional {
			//only capture groups
			for _, captureGroup := range matches[1:] {
				if len(captureGroup) > 0 {
					result = append(result, captureGroup)
				}
			}
		} else {
			//only the matches
			result = append(result, matches[0])
		}

	}
	return result
}

func SplitNumbers(input *[]string) [][2]uint32 {
	r, _ := regexp.Compile(`\((\d{0,3}),(\d{0,3})\)`)

	rows := make([][2]uint32, len(*input))
	for i, v := range *input {
		matches := r.FindStringSubmatch(v)

		cells := [2]uint32{}
		for j := 1; j < len(matches); j += 2 {
			v1, err1 := strconv.ParseInt(matches[j], 10, 32)
			v2, err2 := strconv.ParseInt(matches[j+1], 10, 32)

			if err1 != nil || err2 != nil {
				panic("Error while parsing int")
			}

			cells[0] = uint32(v1)
			cells[1] = uint32(v2)
		}
		rows[i] = cells
	}

	return rows
}

func MultiplyNumbers(input *[][2]uint32) []uint32 {
	results := make([]uint32, len(*input))
	for i, row := range *input {
		results[i] = row[0] * row[1]
	}
	return results
}

func SumNumbers(input *[]uint32) uint32 {
	var sum uint32 = 0
	for _, row := range *input {
		sum += row
	}
	return sum
}
