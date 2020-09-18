// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 136.

// The toposort program prints the nodes of a DAG in topological order.
package main

import (
	"fmt"
	"strings"
)

//!+table
// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string]map[string]bool{
	"algorithms":     {"data structures": true},
	"calculus":       {"linear algebra": true},
	"linear algebra": {"calculus": true},
	"mathematics C":  {"mathematics B": true, "mathematics A": true},
	"mathematics B":  {"mathematics A": true},
	"mathematics A":  {"mathematics C": true},

	"compilers": {
		"data structures":       true,
		"formal languages":      true,
		"computer organization": true,
	},

	"data structures":       {"discrete math": true},
	"databases":             {"data structures": true},
	"discrete math":         {"intro to programming": true},
	"formal languages":      {"discrete math": true},
	"networks":              {"operating systems": true},
	"operating systems":     {"data structures": true, "computer organization": true},
	"programming languages": {"data structures": true, "computer organization": true},
}

//!-table

//!+main
func main() {
	order, cycles := topoSort(prereqs)
	for _, cycle := range cycles {
		fmt.Printf("Cycle detected: %s\n", strings.Join(cycle, " => "))
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string]map[string]bool) (order []string, cycles [][]string) {
	const (
		UNSEEN int8 = iota
		SEEN
		FINISHED
	)

	visited := make(map[string]int8)
	var visitAll func(items map[string]bool, history []string)

	// 循環検出法(Cycle detection)
	// https://qiita.com/drken/items/a803d4fc4a727e02f7ba#4-6-%E3%82%B5%E3%82%A4%E3%82%AF%E3%83%AB%E6%A4%9C%E5%87%BA
	// https://en.wikipedia.org/wiki/Cycle_(graph_theory)#Cycle_detection
	visitAll = func(items map[string]bool, history []string) {
		for item := range items {
			if visited[item] == UNSEEN {
				visited[item] = SEEN
				visitAll(m[item], append(history, item))
				visited[item] = FINISHED
				order = append(order, item)
			} else if visited[item] != FINISHED {
				cycles = append(cycles, history)
			}
		}
	}

	keys := map[string]bool{}
	for key := range m {
		keys[key] = true
	}

	visitAll(keys, nil)
	return order, cycles
}

//!-main
