package main

import (
    "bufio"
    "fmt"
    "net"
)

// Connection struct
type Connection struct {
    // Server connection
    server string
    socket net.Conn

    // User information
    nick, user string

    // Read/write channels
    chanRead, chanWrite chan string

    // Exit channels
    chanExit chan bool
}

// Create a new IRC connection object
func IRC(nick, user string) *Connection {
    irc := &Connection {
        nick: nick,
        user: user,
        chanExit: make(chan bool),
    }
    return irc
}

// Method to connect to IRC server
func (irc *Connection) Connect(server string) error {
    var err error
    irc.server = server
    irc.socket, err = net.Dial("tcp", server)
    if err != nil {
        return err
    }
    fmt.Printf("Connected to %s (%s)\n", irc.server, irc.socket.RemoteAddr())

    irc.chanRead = make(chan string, 16)
    irc.chanWrite = make(chan string, 16)

    go irc.readLoop()
    go irc.writeLoop()

    return nil
}

// Write loop
func (irc *Connection) writeLoop() {
    for {
        data := <-irc.chanWrite

        fmt.Printf("--> %s\n", data)
        _, err := irc.socket.Write([]byte(data))
        if err != nil {
            // TODO log errors
            break
        }
    }
}

// Read loop
func (irc *Connection) readLoop() {
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

    irc.chanWrite <- fmt.Sprintf("JOIN %s\r\n", channel)

    <-irc.chanExit
}
