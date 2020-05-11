package access

import (
	"fmt"
	"sort"

	"../storage"
)

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}
type Chart = []Point

type Plotter struct {
	s storage.Storage
}

func removeFromIndex(array []int, rindex []int) []int {
	sort.Ints(rindex)
	fmt.Println(rindex)

	for weight, index := range rindex {
		if index-weight < 0 {

		} else if index-weight < len(array) {
			array = append(array[:index-weight], array[index-weight+1:]...)
		} else {
			array = array[:index-weight]
		}
	}
	rindex = []int{}
	return array
}

func (a *Plotter) PlotChartAccessOfPage(page string, from int, to int, steps int) Chart { // Isso deveria sair daqui. Criar um PlotterClass
	accesses, ok := a.s.Read("access").(map[string]Accesses)

	if !ok {
		return Chart{}
	}

	accessOfPage := make([]int, len(accesses[page]))
	copy(accessOfPage, accesses[page])
	chart := Chart{}
	from = (from / steps) * steps
	to = (to/steps)*steps + steps

	rindex := make([]int, 0)

	for index, timeOfAccess := range accessOfPage {
		if timeOfAccess < from {
			rindex = append(rindex, index)
		}
	}
	accessOfPage = removeFromIndex(accessOfPage, rindex)
	rindex = []int{}

	for i := from; i < to; i += steps {
		point := Point{X: i, Y: 0}
		for index, timeOfAccess := range accessOfPage {
			if timeOfAccess <= i+steps {
				rindex = append(rindex, index)
				point.Y++
				fmt.Println(timeOfAccess)
			}
		}
		accessOfPage = removeFromIndex(accessOfPage, rindex)
		rindex = []int{}
		chart = append(chart, point)
	}

	return chart
}

func NewPlotter(s storage.Storage) *Plotter {
	return &Plotter{
		s: s,
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
