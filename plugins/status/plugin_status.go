package plugin_status

import (
	"github.com/thoj/go-ircevent"

	"github.com/coredump-ch/moss/conf"
	"github.com/coredump-ch/moss/plugin"
)

func printStatus(args string, e *irc.Event, con *irc.Connection) error {
	con.Privmsg(conf.Channel, "Status unknown.")
	return nil
}

func init() {
	plugin.RegisterCommand("status", "PRIVMSG", printStatus)
}
