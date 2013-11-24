package main

import (
    "fmt"
    "log"
    "strings"

    "github.com/thoj/go-ircevent"

    "github.com/coredump-ch/moss/rivebot"
)

func main() {
    fmt.Println("Moss is starting...")

    const (
        server  = "irc.freenode.net:6697"
        channel = "#coredump"
        nick    = "mossbot"
        user    = "mossbot"
    )

    con := irc.IRC(nick, user)
    con.UseTLS = true

    // Start rivebot
    rbot := rivebot.NewRivebot()
    rbot.Start()

    // Connect
    err := con.Connect(server)
    if err != nil {
        fmt.Println("Failed connecting.")
        return
    }

    // Join channel
    con.AddCallback("001", func(e *irc.Event) {
        con.Join(channel)
    })

    // Reply to mentions
    con.AddCallback("PRIVMSG", func(e *irc.Event) {
        if strings.HasPrefix(e.Message, nick) {
            msg := strings.TrimPrefix(e.Message, nick)
            msg = strings.TrimLeftFunc(msg, func(char rune) bool {
                return char == ',' || char == ':' || char == '-' || char == ' '
            })
            reply, err := rbot.Ask(msg)
            if err != nil {
                log.Printf("Error: %s", err)
            } else {
                con.Privmsg(channel, reply)
            }
        }
    })

    // Event loop
    con.Loop()
}
