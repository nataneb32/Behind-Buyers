package access

import (
	"fmt"
	"testing"
	"time"

	"../storage"
	"github.com/stretchr/testify/assert"
)

func TestStore(t *testing.T) {
	assert := assert.New(t)
	s := storage.NewInMemory()
	a := CreateListener(s)
	a.store("test", 1)
	a.store("test", 2)
	a.store("test", 1)

	data := a.GetData()

	assert.Equal(map[string]Accesses{"test": Accesses{1, 1, 2}}, data)
}

func TestNotify(t *testing.T) {
	s := storage.NewInMemory()
	a := CreateListener(s)
	a.Notify("access", "test")
	d, _ := time.ParseDuration("1s")
	time.Sleep(d)
	a.Notify("access", "test")

	fmt.Println(a.data)
}
