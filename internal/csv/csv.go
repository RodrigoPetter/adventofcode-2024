package csv

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

func ReadCsvFileInput(path string) (cells [][]uint32) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("Unable to read input file ", err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	csvReader.Comma = ' '
	csvReader.FieldsPerRecord = -1
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV", err)
	}

	var numberCells [][]uint32
	for rowNumber, row := range records {
		cells := make([]uint32, len(row))
		for columnNumber, cellValue := range row {
			i, _ := strconv.ParseUint(cellValue, 10, 32)
			if err != nil {
				log.Fatalf(`Invalid uint32 at csv row %v and collumn %v.`, rowNumber, columnNumber)
			}
			cells[columnNumber] = uint32(i)
		}
		numberCells = append(numberCells, cells)
	}

	return numberCells
}
