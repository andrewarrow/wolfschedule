package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/andrewarrow/wolfschedule/redis"
)

func main() {
	files := sortedFiles()
	for _, file := range files {
		asBytes, _ := ioutil.ReadFile("data/" + file)
		lines := strings.Split(string(asBytes), "\n")

		for _, line := range lines {
			tokens := strings.Split(file, ".")
			ts, _ := strconv.ParseInt(tokens[0], 10, 64)
			redis.InsertItem(ts, line)
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
