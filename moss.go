package main

import (
    "fmt"
    "log"
    "strings"

    "github.com/thoj/go-ircevent"

    "github.com/coredump-ch/moss/conf"
    "github.com/coredump-ch/moss/plugin"
    "github.com/coredump-ch/moss/rivebot"

    // Plugins
    _ "github.com/coredump-ch/moss/plugins/status"
)

func main() {
    fmt.Println("Moss is starting...")

    // Create connection
    con := irc.IRC(conf.Nick, conf.User)
    con.UseTLS = true

    // Start rivebot
    rbot := rivebot.NewRivebot()
    rbot.Start()

    // Connect
    err := con.Connect(conf.Server)
    if err != nil {
        fmt.Println("Failed connecting.")
        return
    }

    // Register plugins
    for _, initFunc := range plugin.Plugins {
        initFunc(con)
    }

    // Join channel
    con.AddCallback("001", func(e *irc.Event) {
        con.Join(conf.Channel)
    })

    // Reply to mentions
    con.AddCallback("PRIVMSG", func(e *irc.Event) {
        if strings.HasPrefix(e.Message, conf.Nick) {
            msg := strings.TrimPrefix(e.Message, conf.Nick)
            msg = strings.TrimLeftFunc(msg, func(char rune) bool {
                return char == ',' || char == ':' || char == '-' || char == ' '
            })
            reply, err := rbot.Ask(msg)
            if err != nil {
                log.Printf("Error: %s", err)
            } else {
                con.Privmsg(conf.Channel, reply)
            }
        }
    })

    // Event loop
    con.Loop()
}
