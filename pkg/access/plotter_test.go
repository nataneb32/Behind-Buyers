package access

import (
	"fmt"
	"testing"

	"../storage"
	"github.com/stretchr/testify/assert"
)

func TestPlotChart(t *testing.T) {
	assert := assert.New(t)
	s := storage.NewInMemory()

	plotter := NewPlotter(s)

	plotter.s.Store("access", map[string][]int{
		"test": []int{
			0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 11,
		},
	},
	)

	data := plotter.PlotChartAccessOfPage("test", 1, 10, 5)
	fmt.Println(data)
	assert.Equal(Chart{Point{0, 6}, Point{5, 4}, Point{10, 1}}, data)
}
