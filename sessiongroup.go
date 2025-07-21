package discordgo

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
		if err := s.Open(); err != nil {
			s.log(LogError, "failed to open shard %d: %s", s.ShardID, err.Error())
			g.CloseAll()
			return err
		}
	}
	return nil
}
