package session

import (
	"fmt"
	"testing"
)

func TestGenSession(t *testing.T) {
	g := Generator{
		sessions: []Session{},
	}
	token := g.generateToken()
	if !g.validate(token) {
		t.Errorf("A token validation isn't work.")
	}
	fmt.Println(g.sessions[0].created_at)
}
