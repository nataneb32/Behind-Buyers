package session

import (
	"fmt"
	"math/rand"
	"time"
)

type Session struct {
	token      string
	created_at time.Time
}

type Generator struct {
	sessions []Session
}

func (s *Generator) generateToken() string {
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	session := Session{
		token:      fmt.Sprint(seededRand.Int()),
		created_at: time.Now(),
	}

	s.sessions = append(s.sessions, session)
	return session.token
}

func (g *Generator) validate(token string) bool {
	for _, session := range g.sessions {
		if token == session.token {
			return true
		}
	}
	return false
}
