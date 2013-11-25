package main

import (
    "fmt"
    "log"
    "strings"

    // 3rd party libraries
    "github.com/thoj/go-ircevent"

    // Internal packages
    "github.com/coredump-ch/moss/conf"
    "github.com/coredump-ch/moss/plugin"
    "github.com/coredump-ch/moss/rivebot"

    // Plugins
    _ "github.com/coredump-ch/moss/plugins/status"
)

// Clean a message, remove leading username
func cleanMessage(rawMessage string) string {
    msg := strings.TrimPrefix(rawMessage, conf.Nick)
    msg = strings.TrimLeftFunc(msg, func(char rune) bool {
        return char == ',' || char == ':' || char == '-' || char == ' '
    })
    return msg
}

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
    //plugin.Initialize(con)

    // Join channel
    con.AddCallback("001", func(e *irc.Event) {
        con.Join(conf.Channel)
    })

    // Reply to mentions
    con.AddCallback("PRIVMSG", func(e *irc.Event) {
        if strings.HasPrefix(e.Message, conf.Nick) {
            msg := cleanMessage(e.Message)
            // Commands start with a bang (!)
            if strings.HasPrefix(msg, "!") {
                parts := strings.SplitN(strings.TrimPrefix(msg, "!"), " ", 2)
                var args, command string
                command = parts[0]
                if len(parts) > 1 {
                    args = parts[1]
                } else {
                    args = ""
                }
                log.Printf("Invoking command: %s", command)
                err := plugin.InvokeCommand(command, args, e, con)
                if err != nil {
                    log.Printf("*** %s", err)
                }
            } else {
                reply, err := rbot.Ask(msg)
                if err != nil {
                    log.Printf("Error: %s", err)
                } else {
                    con.Privmsg(conf.Channel, reply)
                }
            }
        }
    })

    // Event loop
    con.Loop()
}
