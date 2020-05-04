package access

import "../storage"

type Chart struct {
	Label []int
	Data  []int
}

type Plotter struct {
	s storage.Storage
}

func (a *Plotter) PlotChartAccessOfPage(page string, from int, to int, steps int32) Chart { // Isso deveria sair daqui. Criar um PlotterClass
	accesses, ok := a.s.Read("access").(map[string]Accesses)

	if !ok {
		return Chart{Label: []int{}, Data: []int{}}
	}

	accessOfPage := accesses[page]
	chart := Chart{}

	for i := 0; i < len(accessOfPage); i++ {
		if accessOfPage[i] < from {
			continue
		}
		if accessOfPage[i] > to {
			continue
		}
		if len(chart.Label) == 0 {
			chart.Label = append(chart.Label, accessOfPage[i])
			chart.Data = append(chart.Data, 1)
		} else if abs(chart.Label[len(chart.Label)-1]-accessOfPage[i]) >= int(steps) {
			chart.Label = append(chart.Label, accessOfPage[i])
			chart.Data = append(chart.Data, 1)
		} else {
			chart.Data[len(chart.Data)-1]++
		}
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
