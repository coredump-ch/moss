package main

import (
    "fmt"
    "net"
)

// Connection struct
type Connection struct {
    server string
    socket net.Conn

    nick, user string
}

// Create a new IRC connection object
func IRC(nick, user string) *Connection {
    irc := &Connection {
        nick: nick,
        user: user,
    }
    return irc
}

// Method to connect to IRC server
func (irc *Connection) Connect(server string) error {
    var err error
    irc.socket, err = net.Dial("tcp", server)
    if err != nil {
        return err
    }
    return nil
}

// Main loop
func main() {
    fmt.Println("Moss is starting...")

    const server = "irc.freenode.org:6667"
    const channel = "#coredump"

    irc := IRC("mossbot", "mossbot")
    err := irc.Connect(server)
    if err != nil {
        panic(err)
    }
}
