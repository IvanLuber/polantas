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
	if admin[from] {
		m.Reply("Ya maap boss :'( @" + from)
		return
	}
	m.Replyf("Kurang ajar @%s, berani ngelawan polantas !?", from)
}

func HandleHelp(m *tbot.Message) {
	m.Reply(`Polantas adalah bot untuk menertibkan antrian reporting stand up SHG
	Command bot polantas:
	1. /start untuk mau mulai stand up. Urutan tim dan member tim akan dipilih secara acak.
	2. /next untuk lanjut ke member berikutnya apabila member telah selesai melakukan reporting. Command hanya bisa dijalankan oleh member yg mendapat giliran
	3. /reset untuk menghapus antrian dan memulai stand up dari awal.
	4. /skip jika member ingin dimasukan ke antrian belakang (emergency only).
	5. /cuti apabila member sedang melakukan cuti.`)
}
