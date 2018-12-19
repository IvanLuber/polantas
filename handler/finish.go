package handler

import (
	"github.com/yanzay/tbot"
)

func HandleFinish(m *tbot.Message) {
	if name != m.From.UserName && m.From.UserName != admin {
		m.Replyf("Oi sabar napa @%s, ini lagi giliran @%s", m.From.UserName, name)
		return
	}

	m.Replyf("Oke, Makasih reportnya @%s", name)
	if teamOrder.len() == 0 && memberOrder.len() == 0 {
		isStarted = false
		m.Reply("Oke udah semuanya nih, terima kasih..")
		return
	}
	if memberOrder.len() == 0 {
		team, teamOrder = teamOrder.pop()
		initMemberStack(team)

		name, memberOrder = memberOrder.pop()
		m.Reply("Sekarang lanjut ke tim " + team)
		m.Reply("Yuk mulai dari @" + name)
		return
	}

	name, memberOrder = memberOrder.pop()
	m.Reply("next @" + name + "?")
}
