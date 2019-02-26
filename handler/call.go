package handler

import (
	"strings"

	"github.com/yanzay/tbot"
)

func HandleCall(m *tbot.Message) {
	team := strings.ToUpper(m.Vars["team"])
	members := allMembers[team]
	if len(members) == 0 {
		m.Reply("Yang bener dong mention timnya !\nPick one: BE, APP, FE dan QA")
		return
	}
	var mentions string

	for _, v := range members {
		mentions += "@" + v + " "
	}
	m.Reply("#" + team + " " + mentions)
}
