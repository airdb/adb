package adblib

import (
	"github.com/MakeNowJust/heredoc"
)

var (
	DNSSrvDoc = heredoc.Doc(`
$ adb srv add _sip._tcp "10 60 443 bigbox.airdb.red"
$ adb srv add _sip._tcp "10 20 443 smallbox1.airdb.blue"
$ adb srv add _sip._tcp "10 20 443 smallbox2.airdb.green"
$ adb srv add _sip._tcp "20 0  443 backup.airdb.black"

> Refer: https://zh.wikipedia.org/wiki/SRV记录
	`)
)
