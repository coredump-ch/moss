package main

import (
    "fmt"
    "strings"

    "github.com/thoj/go-ircevent"
)

func main() {
    fmt.Println("Moss is starting...")

    const (
        server = "irc.freenode.net:6667"
        channel = "#coredump"
        nick = "mossbot"
        user = "mossbot"
    )

    con := irc.IRC(nick, user)
    //con.UseTLS = true

    // Connect
    err := con.Connect(server)
    if err != nil {
        fmt.Println("Failed connecting.")
        return
    }

    // Join channel
    con.AddCallback("001", func (e *irc.Event) {
        con.Join(channel)
    })

    // Reply to mentions
    con.AddCallback("PRIVMSG", func (e *irc.Event) {
        if strings.Contains(e.Message, nick) {
            reply := fmt.Sprintf("Hi, %s!", e.Nick)
            con.Privmsg(channel, reply)
        }
    })

    // Event loop
    con.Loop()
}
