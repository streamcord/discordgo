package discordgo

import (
	"testing"
)

func TestContentWithMoreMentionsReplaced(t *testing.T) {
	s := &Session{StateEnabled: true, State: NewState()}

	user := &User{
		ID:       "user",
		Username: "User Name",
	}

	m := &Message{
		Content:      "<@&role> <@!user> <@user> <#channel>",
		ChannelID:    "channel",
		MentionRoles: []string{"role"},
		Mentions:     []*User{user},
	}
	if result, _ := m.ContentWithMoreMentionsReplaced(s); result != "@Role Name @User Nick @User Name #Channel Name" {
		t.Error(result)
	}
}
func TestGettingEmojisFromMessage(t *testing.T) {
	msg := "test test <:kitty14:811736565172011058> <:kitty4:811736468812595260>"
	m := &Message{
		Content: msg,
	}
	emojis := m.GetCustomEmojis()
	if len(emojis) < 1 {
		t.Error("No emojis found.")
		return
	}

}
