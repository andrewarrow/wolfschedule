package main

import (
	"sort"
	"strings"
)

var mapByOne = map[string][]int{}

func allThe(s ...string) string {
	list := []int{}
	for _, item := range s {
		list = append(list, mapByOne[item]...)
	}
	if len(list) == 0 {
		return ""
	}
	sort.Ints(list)
	max := list[len(list)-1]
	m := map[int]bool{}
	for _, thing := range list {
		m[thing] = true
	}
	buff := []string{}
	for i := 0; i < max; i++ {
		if m[i] {
			buff = append(buff, "**")
		} else {
			buff = append(buff, "  ")
		}
	}
	return strings.Join(buff, " ")
}
