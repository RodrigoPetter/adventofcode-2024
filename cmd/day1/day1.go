package day1

import (
	"slices"
)

func Distances(list1 []uint32, list2 []uint32) uint32 {
	slices.Sort(list1)
	slices.Sort(list2)

	var distance uint32 = 0
	for i := range len(list1) {
		distance += max(list1[i], list2[i]) - min(list1[i], list2[i])
	}
	return distance
}

func Similarity(list1 []uint32, list2 []uint32) uint32 {
	map2 := make(map[uint32]uint32)

	for i := range len(list1) {
		map2[list2[i]] += 1
	}

	var similarity uint32 = 0
	for _, k := range list1 {
		//Em GO, se o map não contem o elemento vai retornar 0
		similarity += k * map2[k]
	}
	return similarity
}
