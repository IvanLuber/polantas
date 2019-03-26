package handler

import (
	"math/rand"
	"time"

	"github.com/yanzay/tbot"
)

var (
	allMembers = map[string][]string{
		"BE":  []string{"IvanLuberiski"},
		"APP": []string{"rifq39", "udafik"},
		"FE":  []string{"ferrwan"},
		"QA":  []string{"luthfiswees", "ricosc27", "ulvasianturi", "yanisihombing", "RirinZulandra"},
		"PM":  []string{"fauridhomahran"},
	}
	admin = map[string]bool{
		"nurudianto":     true,
		"yogieclinov":    true,
		"windawinanti":   true,
		"galihmuji":      true,
		"fauridhomahran": true,
		"IvanLuberiski":  true,
	}
	teams       = []string{"BE", "APP", "FE", "QA"}
	teamOrder   stack
	memberOrder stack
	skipOrder   stack
	name        string
	team        string
	isStarted   bool
	startTime   time.Time
)

func HandleStart(m *tbot.Message) {
	if isStarted {
		m.Replyf("Oi stand up nya udah di start @%s !", m.From.UserName)
		return
	}
	isStarted = true
	startTime = time.Now()
	initStack()
	team, teamOrder = teamOrder.pop()
	initMemberStack(team)

	//get first member
	name, memberOrder = memberOrder.pop()
	t := time.Now()
	m.Reply("Yok dimulai stand up " + t.Format("_2 January 2006"))
	m.Reply("Acak ya mulai dari tim " + team)
	m.Reply("Dimulai dari @" + name)
}

func initStack() {
	if len(memberOrder) != 0 {
		return
	}
	shuffledTeam := shuffle(teams)
	for i := range shuffledTeam {
		teamOrder = teamOrder.push(shuffledTeam[i])
	}
}

func initMemberStack(team string) {
	members := shuffle(allMembers[team])
	for i := range members {
		memberOrder = memberOrder.push(members[i])
	}
}

func shuffle(a []string) []string {
	rand.Seed(int64(time.Now().UnixNano()))
	rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
	return a
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
