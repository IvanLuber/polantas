package handler

import (
	"math/rand"
	"time"

	"github.com/yanzay/tbot"
)

var (
	allMembers = map[string][]string{
		"BE":  []string{"IvanLuberiski", "nurudianto", "qisthi", "ShandyD", "kennabila"},
		"APP": []string{"wahyuprihantoro", "yogieclinov", "nayanda", "rifq39", "Sumanto01", "ryanbaskara", "ngengs"},
		"FE":  []string{"ferrwan", "fauzana_s"},
		"QA":  []string{"luthfiswees", "ricosc27", "ulvasianturi", "yanisihombing", "Arga91", "RirinZulandra"},
	}
	teams       = []string{"BE", "APP", "FE", "QA"}
	admin       = "nurudianto"
	teamOrder   stack
	memberOrder stack
	name        string
	team        string
	isStarted   bool
)

func HandleStart(m *tbot.Message) {
	if isStarted {
		m.Replyf("Oi stand up nya udah di start @%s !", m.From.UserName)
		return
	}
	isStarted = true
	initStack()
	team, teamOrder = teamOrder.pop()
	initMemberStack(team)

	//get first member
	name, memberOrder = memberOrder.pop()
	t := time.Now()
	m.Reply("Yok dimulai stand up " + t.Format("Rabu _2 January 2006"))
	m.Reply("Acak ya mulai dari tim " + team)
	m.Reply("Dimulai dari @" + name)
}

func initStack() {
	if len(memberOrder) != 0 {
		return
	}
	start := rand.Intn(len(teams))
	for i := range teams {
		index := (start + i) % len(teams)
		teamOrder = teamOrder.push(teams[index])
	}
}

func initMemberStack(team string) {
	members := allMembers[team]
	start := rand.Intn(len(members))

	for i := range members {
		index := (start + i) % len(members)
		memberOrder = memberOrder.push(members[index])
	}
}

//build stack

type stack []string

func (s stack) pop() (string, stack) {
	if s.len() == 0 {
		return "", s
	}
	return s[len(s)-1], s[:len(s)-1]
}

func (s stack) peek() string {
	if s.len() == 0 {
		return ""
	}
	return s[len(s)-1]
}

func (s stack) push(name string) stack {
	return append(s, name)
}

func (s stack) len() int {
	return len(s)
}
