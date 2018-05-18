// This is an example program showing the usage of hellabot
package main

import (
    "flag"
    "fmt"
    //"time"

    "github.com/valeness/hellabot"
    log "gopkg.in/inconshreveable/log15.v2"
)

var serv = flag.String("server", "irc.freenode.net:6667", "hostname and port for irc server to connect to")
var nick = flag.String("nick", "wotbot", "nickname for the bot")

func main() {
    flag.Parse()

    hijackSession := func(bot *hbot.Bot) {
        bot.HijackSession = true
    }
    channels := func(bot *hbot.Bot) {
        bot.Channels = []string{"#bottest123"}
    }
    irc, err := hbot.NewBot(*serv, *nick, hijackSession, channels)
    if err != nil {
        panic(err)
    }

    irc.AddTrigger(Trigger)
    irc.Logger.SetHandler(log.StdoutHandler)
    //logHandler := log.LvlFilterHandler(log.LvlInfo, log.StdoutHandler)
    // or
    //irc.Logger.SetHandler(logHandler)
    // or
    // irc.Logger.SetHandler(log.StreamHandler(os.Stdout, log.JsonFormat()))

    // Start up bot (this blocks until we disconnect)
    irc.Run()
    fmt.Println("Bot shutting down.")
}

var Trigger = hbot.Trigger{
    func (b *hbot.Bot, m *hbot.Message) bool {
        return m.Content == "-info"
    },

    func(b *hbot.Bot, m *hbot.Message) bool {
        b.Reply(m, "Hello World")
        return false;
    },
}

