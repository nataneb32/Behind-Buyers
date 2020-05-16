package access

import (
	"sort"
	"time"

	"../storage"
)

type Accesses = []int

type Listener struct {
	data storage.Storage
}

func (a *Listener) GetData() map[string]Accesses {
	var accesses map[string]Accesses
	a.data.Read("access", &accesses)
	return accesses
}

func (a *Listener) store(page string, timeInMilliseconds int) {
	var old map[string]Accesses
	a.data.Read("access", &old)
	if old == nil {
		old = make(map[string][]int)
	}
	new := old
	new[page] = append(old[page], timeInMilliseconds)
	sort.Ints(new[page])
	a.data.Store("access", new)
}

func (a *Listener) Notify(report map[string]interface{}) {
	if page, ok := report["page"].(string); ok {
		now := time.Now()
		a.store(page, int(now.UnixNano()/1e6))
	}
}

func CreateListener(s storage.Storage) *Listener {
	return &Listener{
		data: s,
	}
}
