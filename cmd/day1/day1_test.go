package day1

import (
	"testing"

	"github.com/RodrigoPetter/adventofcode-2024/internal/csv"
)

func TestDistancesExample(t *testing.T) {
	//given
	list1 := []uint32{3, 4, 2, 1, 3, 3}
	list2 := []uint32{4, 3, 5, 3, 9, 3}
	const expected uint32 = 11
	//when
	result := Distances(list1, list2)
	//then
	if expected != result {
		t.Errorf(`Expected %v but got %v`, expected, result)
	}
}

func TestDistancesShouldHandleNegativeDistances(t *testing.T) {
	//given
	list1 := []uint32{7, 8, 9}
	list2 := []uint32{1, 2, 3}
	const expected uint32 = 18
	//when
	result := Distances(list1, list2)
	//then
	if expected != result {
		t.Errorf(`Expected %v but got %v`, expected, result)
	}
}

func TestDistancesQuestion(t *testing.T) {
	//given
	list1, list2 := getColumns(csv.ReadCsvFileInput("../../cmd/day1/input.csv", ' '))
	var expected uint32 = 1873376
	//when
	result := Distances(list1, list2)
	//then
	if expected != result {
		t.Errorf(`Expected %v but got %v`, expected, result)
	}
}

func TestSimilarityExample(t *testing.T) {
	// given
	list1 := []uint32{3, 4, 2, 1, 3, 3}
	list2 := []uint32{4, 3, 5, 3, 9, 3}
	const expected uint32 = 31
	// when
	result := Similarity(list1, list2)
	// then
	if expected != result {
		t.Errorf(`Expected %v but got %v`, expected, result)
	}
}

func TestSimilarityQuestion(t *testing.T) {
	//given
	list1, list2 := getColumns(csv.ReadCsvFileInput("../../cmd/day1/input.csv", ' '))
	var expected uint32 = 18997088
	//when
	result := Similarity(list1, list2)
	//then
	if expected != result {
		t.Errorf(`Expected %v but got %v`, expected, result)
	}
}

func getColumns(data [][]uint32) (column1 []uint32, column2 []uint32) {
	firstColumn := make([]uint32, len(data))
	secondColumn := make([]uint32, len(data))

	// Extract columns
	for i, row := range data {
		if len(row) > 0 {
			firstColumn[i] = row[0]
		}
		if len(row) > 1 {
			secondColumn[i] = row[1]
		}
	}

	return firstColumn, secondColumn
}
