package handler

import "github.com/yanzay/tbot"

func HandleReset(m *tbot.Message) {
	isStarted = false
	memberOrder = nil
	teamOrder = nil
	m.Replyf("Reset by @%s", m.From.UserName)
}
