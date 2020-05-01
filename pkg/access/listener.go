package access

import (
	"fmt"
	"sort"
	"time"

	"../storage"
)

type Chart struct {
	Label []int
	Data  []int
}

type Accesses = []int

type Listener struct {
	data storage.Storage
}

func (a *Listener) GetData() map[string]Accesses {
	fmt.Println(a.data.Read("access"))
	if accesses, ok := a.data.Read("access").(map[string]Accesses); ok {
		return accesses
	}
	return nil
}

func (a *Listener) store(page string, timeInMilliseconds int) {
	old, ok := a.data.Read("access").(map[string][]int)
	if !ok {
		old = make(map[string][]int)
	}
	new := old
	new[page] = append(old[page], timeInMilliseconds)
	sort.Ints(new[page])
	a.data.Store("access", new)
}

func (a *Listener) Notify(fieldType string, value interface{}) {
	if fieldType == "access" {
		if page, ok := value.(string); ok {
			now := time.Now()
			a.store(page, int(now.UnixNano()/1e6))
		}
	}
}

func CreateListener(s storage.Storage) *Listener {
	return &Listener{
		data: s,
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
