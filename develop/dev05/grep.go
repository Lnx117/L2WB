package main

import (
	"fmt"
	"sort"
)

type Node struct {
	Key   int
	Value string
}

type GrepStruct struct {
	Result []Node
}

func NewGrep() *GrepStruct {
	return &GrepStruct{
		Result: []Node{},
	}
}

//Сортировка строк результата по возрастанию
func (g *GrepStruct) SortResultASC() {
	sort.Slice(g.Result, func(i, j int) bool {
		return g.Result[i].Key < g.Result[j].Key
	})
}

func (g *GrepStruct) Print(indexing bool) {
	g.SortResultASC()
	switch indexing {
	//Если нужно печатать номер строки, то этот вариант
	case true:
		for _, v := range g.Result {
			fmt.Printf("%d. %s\n", v.Key, v.Value)
		}
	default:
		for _, v := range g.Result {
			fmt.Printf("%s\n", v.Value)
		}
	}
}
