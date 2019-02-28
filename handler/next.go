package handler

import (
	"strconv"
	"time"

	"github.com/yanzay/tbot"
)

func HandleNext(m *tbot.Message) {
	from := m.From.UserName
	if name != from && !admin[from] {
		m.Replyf("Oi sabar napa @%s, ini lagi giliran @%s", from, name)
		return
	}
	m.Replyf("Oke, Makasih reportnya @%s", name)
	logic(m)
}

func HandleSkip(m *tbot.Message) {
	from := m.From.UserName
	if name != from && !admin[from] {
		m.Replyf("Oi sabar napa @%s, ini lagi giliran @%s", m.From.UserName, name)
		return
	}
	skipOrder = skipOrder.push(name)
	m.Replyf("Oke, di akhir nanti akan panggil lagi @%s", name)
	logic(m)
}

func HandleCuti(m *tbot.Message) {
	from := m.From.UserName
	if name != from && !admin[from] {
		m.Replyf("Oi sabar napa @%s, ini lagi giliran @%s", from, name)
		return
	}
	m.Reply("Bagus2, cutilah karena cuti sudah tidak bisa diuangkan.")
	logic(m)
}

func logic(m *tbot.Message) {
	if teamOrder.len() == 0 && memberOrder.len() == 0 {
		if skipOrder.len() != 0 {
			m.Reply("Sekarang giliran member yang tadi skip")
			memberOrder = skipOrder
			skipOrder = stack{}
		} else {
			isStarted = false
			duration := time.Since(startTime).Minutes()
			m.Reply("Oke udah semuanya nih, terima kasih..")
			m.Reply("Durasi stand up: " + strconv.Itoa(int(duration)) + " Menit")
			if int(duration) > 15 {
				m.Reply("Stand up selanjutnya harus lebih fokus!")
			} else {
				m.Reply("Mantaab....")
			}
			return
		}
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
	m.Reply("selanjutnya @" + name + " ?")
}
