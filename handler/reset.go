package handler

import "github.com/yanzay/tbot"

func HandleReset(m *tbot.Message) {
	isStarted = false
	memberOrder = stack{}
	teamOrder = stack{}
	m.Replyf("Reset by @%s", m.From.UserName)
}

func HandleBodoh(m *tbot.Message) {
	from := m.From.UserName
	if from == admin {
		m.Reply("Ya maap boss :'( @" + from)
		return
	}
	m.Replyf("Kurang ajar @%s, berani ngelawan polantas !?", from)
}
