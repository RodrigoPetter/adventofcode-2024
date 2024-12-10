package day5

import (
	"testing"

	"github.com/RodrigoPetter/adventofcode-2024/internal/csv"
)

func TestShouldCountTwoElements(t *testing.T) {
	result, _ := GetValids(
		[][]uint32{
			[]uint32{9, 5},
			[]uint32{5, 9},
		},
		[][]uint32{
			[]uint32{9, 5},
		},
	)
	expected := 1

	if len(result) != expected {
		t.Errorf(`Expected
		%v
		but got
		%v`, expected, result)
	}
}

func TestExamplePartOne(t *testing.T) {
	orders := csv.ReadCsvFileInput("../../cmd/day5/input_example_order.csv", '|')
	updates := csv.ReadCsvFileInput("../../cmd/day5/input_example_updates.csv", ',')
	expected := 3

	result, _ := GetValids(updates, orders)

	if len(result) != expected {
		t.Errorf(`Expected
		%v
		but got
		%v`, expected, result)
	}
}

func TestExamplePartOneSumMidle(t *testing.T) {
	orders := csv.ReadCsvFileInput("../../cmd/day5/input_example_order.csv", '|')
	updates := csv.ReadCsvFileInput("../../cmd/day5/input_example_updates.csv", ',')
	expected := 143

	result, _ := GetValids(updates, orders)
	sum := SumMidle(result)

	if sum != expected {
		t.Errorf(`Expected
		%v
		but got
		%v`, expected, sum)
	}
}

func TestQuestionPartOne(t *testing.T) {
	orders := csv.ReadCsvFileInput("../../cmd/day5/input_question_order.csv", '|')
	updates := csv.ReadCsvFileInput("../../cmd/day5/input_question_updates.csv", ',')
	expected := 6041

	result, _ := GetValids(updates, orders)
	sum := SumMidle(result)

	if sum != expected {
		t.Errorf(`Expected
		%v
		but got
		%v`, expected, sum)
	}
}

func TestExamplePartTwoSumMidle(t *testing.T) {
	orders := csv.ReadCsvFileInput("../../cmd/day5/input_example_order.csv", '|')
	updates := csv.ReadCsvFileInput("../../cmd/day5/input_example_updates.csv", ',')
	expected := 123

	_, invalids := GetValids(updates, orders)
	result := Sort(invalids, orders)
	sum := SumMidle(result)

	if sum != expected {
		t.Errorf(`Expected
		%v
		but got
		%v`, expected, sum)
	}
}

func TestQuestionPartTwo(t *testing.T) {
	orders := csv.ReadCsvFileInput("../../cmd/day5/input_question_order.csv", '|')
	updates := csv.ReadCsvFileInput("../../cmd/day5/input_question_updates.csv", ',')
	expected := 4884

	_, invalids := GetValids(updates, orders)
	result := Sort(invalids, orders)
	sum := SumMidle(result)

	if sum != expected {
		t.Errorf(`Expected
		%v
		but got
		%v`, expected, sum)
	}
}
