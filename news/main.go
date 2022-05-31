package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func main() {
	files := sortedFiles()
	m := map[string]map[string]bool{}
	for _, file := range files {
		m[file] = map[string]bool{}
		asBytes, _ := ioutil.ReadFile("data/" + file)
		lines := strings.Split(string(asBytes), "\n")

		for _, line := range lines {
			m[file][line] = true
		}

	}

	for i := 0; i < len(files)-1; i++ {
		for k, _ := range m[files[i]] {
			if m[files[i+1]][k] {
				fmt.Println(files[i], k)
			}
		}
	}

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
