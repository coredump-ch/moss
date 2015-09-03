// BOFH excuses.

package plugin_excuse

import (
	"fmt"
	"github.com/thoj/go-ircevent"
	"math/rand"
	"time"

	"github.com/coredump-ch/moss/conf"
	"github.com/coredump-ch/moss/plugin"
)

func printExcuse(args string, e *irc.Event, con *irc.Connection) error {
	rand.Seed(time.Now().Unix())
	i := rand.Intn(len(excuses))
	con.Privmsg(conf.Channel, fmt.Sprintf("The problem: %s.", excuses[i]))
	return nil
}

func init() {
	plugin.RegisterCommand("excuse", "PRIVMSG", printExcuse)
}
