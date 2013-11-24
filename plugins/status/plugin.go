package plugin_status

import (
    "github.com/thoj/go-ircevent"

    "github.com/coredump-ch/moss/conf"
    "github.com/coredump-ch/moss/plugin"
)

func InitPlugin(con *irc.Connection) error {
    con.AddCallback("PRIVMSG", func(e *irc.Event) {
        con.Privmsg(conf.Channel, "Booya. The status plugin works.")
    })
    return nil
}

func init() {
    // Register the plugin
    plugin.Register(InitPlugin)
}
