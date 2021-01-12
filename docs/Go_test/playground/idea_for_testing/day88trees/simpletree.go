package simpletree

import (
	"fmt"
)

type Tree struct {
	Name string
	R    *Tree
	L    *Tree
}

func subset(seq1 []int, seq2 []int) bool {
	m := map[int]int{}

	for _, v := range seq2 {
		m[v] = 1
	}
	for _, v := range seq1 {
		if _, ok := m[v]; ok {
			m[v]++
		}
	}

	for k, v := range m {
		if v == 1 {
			fmt.Printf("%d not found\n", k)
			return false
		}
	}

	return true
}

func numbersMissing(seq1 []int, seq2 []int) []int {
	m := map[int]int{}

	for _, v := range seq2 {
		m[v] = 1
	}
	for _, v := range seq1 {
		if _, ok := m[v]; ok {
			m[v]++
		}
	}

	missing := []int{}
	for k, v := range m {
		if v == 1 {
			fmt.Printf("%d not found\n", k)
			missing = append(missing,k)
		}
	}

	return missing

}