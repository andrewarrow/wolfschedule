package main

import (
	"fmt"
	"sort"
	"strings"
)

type DigitAndIndex struct {
	Digit byte
	Index int
}

var mapByOne = map[string][]DigitAndIndex{}

func allThe(s ...string) string {
	list := []DigitAndIndex{}
	for _, item := range s {
		list = append(list, mapByOne[item]...)
	}
	if len(list) == 0 {
		return ""
	}
	sort.SliceStable(list, func(i, j int) bool {
		return list[i].Index < list[j].Index
	})
	max := list[len(list)-1]
	m := map[int]byte{}
	for _, thing := range list {
		m[thing.Index] = thing.Digit
	}
	buff := []string{}
	for i := 0; i < max.Index+1; i++ {
		if m[i] != 0 {
			buff = append(buff, fmt.Sprintf(" %d", m[i]))
		} else {
			buff = append(buff, "  ")
		}
	}
	return strings.Join(buff, " ")
}
