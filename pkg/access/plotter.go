package access

import (
	"fmt"
	"sort"

	"../storage"
)

type Chart struct {
	Label []int
	Data  []int
}

type Plotter struct {
	s storage.Storage
}

func removeFromIndex(array []int, rindex []int) []int {
	sort.Ints(array)
	for weight, index := range rindex {
		if index-weight < len(array) {
			array = append(array[:index-weight], array[index-weight+1:]...)
		} else {
			array = array[:index-weight]
		}
	}
	return array
}

func (a *Plotter) PlotChartAccessOfPage(page string, from int, to int, steps int) Chart { // Isso deveria sair daqui. Criar um PlotterClass
	accesses, ok := a.s.Read("access").(map[string]Accesses)

	if !ok {
		return Chart{Label: []int{}, Data: []int{}}
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
		chart.Label = append(chart.Label, i)
		chart.Data = append(chart.Data, 0)
		for index, timeOfAccess := range accessOfPage {
			if timeOfAccess == 0 {
				continue
			}
			fmt.Println(timeOfAccess < i)
			fmt.Printf("From: %i \n", from)
			fmt.Printf("From/Steps: %i \n", from/steps)

			if timeOfAccess <= i+steps {
				rindex = append(rindex, index)
				chart.Data[len(chart.Data)-1]++
				fmt.Println(timeOfAccess)
			}
		}
		accessOfPage = removeFromIndex(accessOfPage, rindex)
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
