package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/tmickel/advent2021/fileutil"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	// input, err := fileutil.ScanStrings("example.txt")
	input, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	graph := make(map[string][]string, 0) // name -> edges

	for _, edge := range input {
		nodes := strings.Split(edge, "-")
		start := nodes[0]
		end := nodes[1]
		if _, ok := graph[start]; !ok {
			graph[start] = make([]string, 0)
		}
		graph[start] = append(graph[start], end)
		if _, ok := graph[end]; !ok {
			graph[end] = make([]string, 0)
		}
		graph[end] = append(graph[end], start)
	}

	paths := make(map[string]bool, 0)
	searchPartOne(graph, paths, "start", "start,")
	fmt.Println(len(paths))
}

func searchPartOne(graph map[string][]string, paths map[string]bool, v string, path string) {
	for _, edge := range graph[v] {
		if edge == "end" {
			paths[path+"end"] = true
			continue
		}
		visited := strings.Contains(path, ","+edge+",")
		if (!visited || isBig(edge)) && edge != "start" {
			searchPartOne(graph, paths, edge, path+edge+",")
		}
	}
}

func partTwo() {
	// input, err := fileutil.ScanStrings("example.txt")
	input, err := fileutil.ScanStrings("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	graph := make(map[string][]string, 0) // name -> edges

	for _, edge := range input {
		nodes := strings.Split(edge, "-")
		start := nodes[0]
		end := nodes[1]
		if _, ok := graph[start]; !ok {
			graph[start] = make([]string, 0)
		}
		graph[start] = append(graph[start], end)
		if _, ok := graph[end]; !ok {
			graph[end] = make([]string, 0)
		}
		graph[end] = append(graph[end], start)
	}

	paths := make(map[string]bool, 0)
	searchPartTwo(graph, paths, "start", "start,", false)
	fmt.Println(len(paths))
}

func isBig(node string) bool {
	return strings.ToUpper(node) == node
}

func searchPartTwo(graph map[string][]string, paths map[string]bool, v string, path string, usedSmall bool) {
	for _, edge := range graph[v] {
		if edge == "end" {
			paths[path+"end"] = true
			continue
		}
		u := usedSmall
		visited := false
		if !isBig(edge) {
			if usedSmall {
				visited = strings.Contains(path, ","+edge+",")
			} else {
				ct := strings.Count(path, ","+edge+",")
				if ct == 1 {
					u = true
				}
				if ct >= 2 {
					visited = true
				}
			}
		}
		if !visited && edge != "start" {
			searchPartTwo(graph, paths, edge, path+edge+",", u)
		}
	}
}
