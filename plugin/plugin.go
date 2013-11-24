// Plugin handler. Inspired by gopherbot.

package plugin

import (
    "github.com/thoj/go-ircevent"
)

// List of plugins
var Plugins = make([]func(*irc.Connection) error, 0, 16)

// Register a new plugin
func Register(initFunc func(*irc.Connection) error) error {
    Plugins = append(Plugins, initFunc)
    return nil
}
