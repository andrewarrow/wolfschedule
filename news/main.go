package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	files := sortedFiles()
	m := map[string]map[string]int{}
	for _, file := range files {
		m[file] = map[string]int{}
		asBytes, _ := ioutil.ReadFile("data/" + file)
		lines := strings.Split(string(asBytes), "\n")

		for _, line := range lines {
			m[file][line] = 1
		}

	}

	// take the freshest list
	//   for each item, compute how many past files it appears in

	freshest := files[0]
	rest := files[1:]

	for k, _ := range m[freshest] {
		for _, history := range rest {
			if m[history][k] == 1 {
				m[freshest][k]++
			}
		}
	}

	items := []Item{}
	for k, v := range m[freshest] {
		item := Item{v, k}
		items = append(items, item)
	}
	sort.Slice(items, func(a, b int) bool {
		return items[a].Count > items[b].Count
	})

	for _, item := range items {
		fmt.Printf("%02d. %s\n", item.Count, item.Title)
	}
}

type Item struct {
	Count int
	Title string
}

func sortedFiles() []string {
	files, _ := ioutil.ReadDir("data")
	list := []string{}
	for _, file := range files {
		list = append(list, file.Name())
	}
	sort.Slice(list, func(a, b int) bool {
		return list[a] > list[b]
	})

	return list

}
