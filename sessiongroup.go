package discordgo

import "time"

// SessionGroup is a helper for coordinating state between multiple Sessions.
type SessionGroup struct {
	Sessions []*Session
}

func NewSessionGroup() *SessionGroup {
	return &SessionGroup{}
}

func (g *SessionGroup) Add(s ...*Session) {
	g.Sessions = append(g.Sessions, s...)
}

func (g *SessionGroup) CloseAll() {
	for _, s := range g.Sessions {
		if err := s.Close(); err != nil {
			s.log(LogError, "failed to close shard %d: %s", s.ShardID, err.Error())
		}
	}
}

func (g *SessionGroup) OpenAll() error {
	for _, s := range g.Sessions {
		for i := 0; i < 3; i++ {
			if err := s.Open(); err == nil {
				break
			} else {
				s.log(LogError, "failed to open shard %d (attempt %d): %s", s.ShardID, i+1, err.Error())
				if i == 2 {
					g.CloseAll()
					return err
				}
				time.Sleep(5 * time.Second)
			}
		}
	}
	return nil
}
