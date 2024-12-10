package day5

import (
	"math"
	"slices"
	"sort"
)

type Node struct {
	data int
	next *Node
}

func GetValids(updates [][]uint32, orders [][]uint32) (valids [][]uint32, invalids [][]uint32) {
	order := getOrders(orders)
	valids = [][]uint32{}
	for _, row := range updates {
		for i, v1 := range row {
			if i+1 == len(row) {
				valids = append(valids, row)
				break
			}
			if i < len(row) && !isValid(v1, row[i+1], order) {
				invalids = append(invalids, row)
				break
			}
		}
	}

	return valids, invalids
}

func Sort(updates [][]uint32, orders [][]uint32) [][]uint32 {
	order := getOrders(orders)

	result := [][]uint32{}
	for _, row := range updates {

		sort.Slice(row, func(i, j int) bool {
			return isValid(row[i], row[j], order)
		})
		result = append(result, row)
	}
	return result
}

func SumMidle(updates [][]uint32) int {
	sum := 0
	for _, row := range updates {
		middle := float64(len(row)) / 2.0
		sum += int(row[int(math.Floor(middle))])
	}
	return sum
}

func getOrders(orders [][]uint32) map[uint32][]uint32 {
	order := make(map[uint32][]uint32)
	for _, v := range orders {
		_, prs := order[v[0]]
		if prs {
			order[v[0]] = append(order[v[0]], v[1])
		} else {
			order[v[0]] = []uint32{v[1]}
		}
	}
	return order
}

func isValid(v1 uint32, v2 uint32, order map[uint32][]uint32) bool {
	orders, prs := order[v1]
	if prs && slices.Contains(orders, v2) {
		return true
	}
	return false
}
