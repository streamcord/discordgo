package discordgo

// SessionGroup is a helper for coordinating state between multiple Sessions.
type SessionGroup struct {
	sessions []*Session
}

func (g *SessionGroup) Add(s ...*Session) {
	g.sessions = append(g.sessions, s...)
}

func (g *SessionGroup) AddFromFactory(shardIDs []int, shardCount int, factory func(shardID, shardCount int) *Session) {
	for _, shardID := range shardIDs {
		g.Add(
			factory(shardID, shardCount),
		)
	}
}

func (g *SessionGroup) CloseAll() {
	for _, s := range g.sessions {
		if err := s.Close(); err != nil {
			s.log(LogError, "failed to close shard %d: %s", s.ShardID, err.Error())
		}
	}
}

func (g *SessionGroup) OpenAll() {
	for _, s := range g.sessions {
		if err := s.Open(); err != nil {
			s.log(LogError, "failed to open shard %d: %s", s.ShardID, err.Error())
		}
	}
}
