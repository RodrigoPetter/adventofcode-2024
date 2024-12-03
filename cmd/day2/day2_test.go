package day2

import (
	"testing"

	"github.com/RodrigoPetter/adventofcode-2024/internal/csv"
)

func TestSafetyExample(t *testing.T) {
	//given
	report := [][]uint32{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	expected := []bool{true, false, false, false, false, true}
	//when
	result := ReportSafety(report, 0)
	//then
	for i, row := range result {
		if row != expected[i] {
			t.Errorf(`Expected %v but got %v`, expected, result)
			break
		}
	}

	if CountSafeReports(result) != 2 {
		t.Errorf(`Expected %v but got %v`, 2, result)
	}
}

func TestSafetyQuestion(t *testing.T) {
	//given
	report := csv.ReadCsvFileInput("../../cmd/day2/input.csv")
	expected := 598
	//when
	result := CountSafeReports(ReportSafety(report, 0))
	//then
	if result != expected {
		t.Errorf(`Expected %v but got %v`, expected, result)
	}
}

func TestProblemDampenerExample(t *testing.T) {
	//given
	report := [][]uint32{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}
	expected := []bool{true, false, false, true, true, true}
	//when
	result := ReportSafety(report, 1)
	//then
	for i, row := range result {
		if row != expected[i] {
			t.Errorf(`Expected %v but got %v`, expected, result)
			break
		}
	}

	if CountSafeReports(result) != 4 {
		t.Errorf(`Expected %v but got %v`, 4, result)
	}
}

func TestProblemDampenerQuestion(t *testing.T) {
	//given
	report := csv.ReadCsvFileInput("../../cmd/day2/input.csv")
	expected := 634
	//when
	result := CountSafeReports(ReportSafety(report, 1))
	//then
	if result != expected {
		t.Errorf(`Expected %v but got %v`, expected, result)
	}
}
