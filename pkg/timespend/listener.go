package timespend

import (
	"../storage"
)

type Listener struct {
	s storage.Storage
}

func (l *Listener) Notify(report map[string]interface{}) {
	if times, ok := report["timespend"].(float64); ok {
		if page, ok := report["page"].(string); ok {
			var old map[string]map[int]int
			l.s.Read("timespend", &old)

			if old == nil {
				l.s.Store("timespend", map[string]map[int]int{
					page: map[int]int{
						int(times): 1,
					},
				})
				return
			}

			old[page][int(times)]++
			l.s.Store("timespend", old)

		}
	}
}

func CreateListener(s storage.Storage) *Listener {
	return &Listener{
		s: s,
	}
}
