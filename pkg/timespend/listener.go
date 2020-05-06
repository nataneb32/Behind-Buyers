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
			if old, ok := l.s.Read("timespend").(map[string]map[int]int); ok {
				old[page][int(times)]++

				l.s.Store("timespend", old)
			} else {
				l.s.Store("timespend", map[string]map[int]int{
					page: map[int]int{
						int(times): 1,
					},
				})
			}
		}
	}
}

func CreateListener(s storage.Storage) *Listener {
	return &Listener{
		s: s,
	}
}
