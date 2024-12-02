package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ReadCsvFileInput() (column1 []uint32, column2 []uint32) {
	f, err := os.Open("../../cmd/day1/input.csv")
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ' '
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV", err)
	}

	var firstColumn, secondColumn []uint32
	for rowNumber, row := range records {
		i1, err1 := strconv.ParseUint(row[0], 10, 32)
		i2, err2 := strconv.ParseUint(row[1], 10, 32)
		if err1 != nil || err2 != nil {
			log.Fatalf(`Invalid uint32 at csv row %v`, rowNumber)
		}
		firstColumn = append(firstColumn, uint32(i1))
		secondColumn = append(secondColumn, uint32(i2))
	}

	return firstColumn, secondColumn
}
