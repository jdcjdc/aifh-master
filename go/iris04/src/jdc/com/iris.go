package main

import  "fmt"
// import "../../fn"
import (
	"../../inout"
	"../../fn"
	"../../tools"
	"strings"
)

func main() {
	// todo hierhier: test whether it all works
	fmt.Printf("Hello Iris\n")

	// var nbrs = []float64 { 0.2, 0.4, 0.6, 0.8 }
	// fn.Evaluate(nbrs)
	file := "/Users/jandecock/wss_sources/ws_intellij/AI/go/iris03/src/iris.csv"

	//inout.ReadFileLines(file)

	var _, pairs = inout.ReadFileLines(file)  // pair: line that contains input data + the specy (the expected outcome, not normalized)
	// todo remove nbrOfSpecies, species := calcSpecies(file, 5)
	nbrOfSpecies, species := calcSpecies(pairs, 5)

	fmt.Printf("nbr of species: %d\n", nbrOfSpecies)

	for _, specy := range species {
		fmt.Printf("%v\n", specy)
	}

	// ToDo take first 4 pars iris.csv and normalize them
	// xs[0] :=

	wi := make([]float64, 4)
	for i := 0; i < 4; i++ {
		wi[i] = tools.RandomFloat64(-10, 10)
	}
	// fn.Evaluate(xs)
}

/**
 @param file: file being process
 @param colnr: the columnr in which the species is. Colnrstarts at 1
 */
func calcSpecies(pairs []string, colnr int16) (int, []string) {
	// todo skip first line, panic when column out of range
	species := make(map[string]int)

	// todo remove var _, lines = inout.ReadFileLines(file)
	// var species []string
	firstLine := true
	for _, line := range pairs {
		// fmt.Printf("%s\n", line)
		if !firstLine {
			var i = 0
			for _, token := range strings.Split(line, ",") {
				//fmt.Printf("%s\n", token)
				i++
				if i == 5 {
					species[token]++
				}

			}
		} else {
			firstLine = false
		}
	}
	retSpecies := make([]string, len(species))
	i := 0
	for specy := range species {
		// jdc not necessary to remove quotes strings.Replace(specy, "\"\"", "", -1)
		retSpecies[i] = specy
		i++
	}

	//fmt.Printf("%d\n", nbr)

	return i, retSpecies;
}
