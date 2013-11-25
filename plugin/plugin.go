// Plugin handler. Inspired by gopherbot.
// TODO define type for func(*irc.Event, *irc.Connection)

package plugin

import (
    "fmt"
    "github.com/thoj/go-ircevent"
    "log"
)

// Plugin stores

var commands = make(map[string]func(string, *irc.Event, *irc.Connection) error)

//var plugins = make([]func(*irc.Connection) error, 0, 16)

// Register a new plugin

/*func RegisterGeneric(initFunc func(*irc.Connection) error) error {
    plugins = append(plugins, initFunc)
    return nil
}*/

func RegisterCommand(command string, eventcode string, callback func(string, *irc.Event, *irc.Connection) error) {
    if _, exists := commands[command]; exists {
        log.Printf("*** Could not load command <%s>: already exists.", command)
    } else {
        commands[command] = callback
    }
}

// Initialize all plugins

/*func Initialize(con *irc.Connection) {
    for _, initFunc := range commands {
        initFunc(con)
    }
    for _, initFunc := range plugins {
        initFunc(con)
    }
}*/

// Invoke a command

func InvokeCommand(command string, args string, e *irc.Event, con *irc.Connection) error {
    callback, exists := commands[command]
    if !exists {
        return fmt.Errorf("Command %s not found.", command)
    }
    return callback(args, e, con)
}
