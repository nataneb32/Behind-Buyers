package click

import (
	"fmt"

	"../storage"
)

type Listener struct {
	s storage.Storage
}

func (l *Listener) Notify(report map[string]interface{}) {
	if clicks, ok := report["clicks"].(map[string]interface{}); ok {
		fmt.Println(clicks)
		var storedValue map[string]int
		l.s.Read("clicks", &storedValue)

		if storedValue == nil {
			storedValue = make(map[string]int)
		}

		for xpath, inclick := range clicks {
			nclick, _ := inclick.(float64)
			storedValue[xpath] += int(nclick)
		}

		l.s.Store("clicks", storedValue)
	}
}

func CreateListener(s storage.Storage) *Listener {
	return &Listener{
		s: s,
	}
}
