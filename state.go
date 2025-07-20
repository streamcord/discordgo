// Discordgo - Discord bindings for Go
// Available at https://github.com/bwmarrin/discordgo

// Copyright 2015-2016 Bruce Marriner <bruce@sqls.net>.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This file contains code related to state tracking.  If enabled, state
// tracking will capture the initial READY packet and many other websocket
// events and maintain an in-memory state of guilds, channels, users, and
// so forth.  This information can be accessed through the Session.State struct.

package discordgo

import (
	"errors"
)

// ErrNilState is returned when the state is nil.
var ErrNilState = errors.New("state not instantiated, please use discordgo.New() or assign Session.State")

// ErrStateNotFound is returned when the state cache
// requested is not found
var ErrStateNotFound = errors.New("state cache not found")

// ErrMessageIncompletePermissions is returned when the message
// requested for permissions does not contain enough data to
// generate the permissions.
var ErrMessageIncompletePermissions = errors.New("message incomplete, unable to determine permissions")

// StateI is an interface that can be used to create custom state engines.
// Check discordgo/state.go for an example implementation of this interface.
//
// WARNING: Custom implementations of this interface must be thread-safe!
type StateI interface {
	Channel(channelID string) (*Channel, error)
	Guild(guildID string) (*Guild, error)
	Member(guildID, userID string) (*Member, error)
	Role(guildID, roleID string) (*Role, error)
	UserChannelPermissions(userID, channelID string) (int64, error)

	// OnInterface is the receiver for all raw state events from the gateway.
	// It must, at minimum, handle READY events to set the value behind SelfUser().
	OnInterface(s *Session, i interface{}) error

	// SelfUser is the bot user returned in the READY event payload.
	SelfUser() *User
}
