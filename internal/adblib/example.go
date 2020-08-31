package adblib

import (
	"github.com/MakeNowJust/heredoc"
)

var (
	SqlDoc = heredoc.Doc(`
adb mysql mina-api -uroot -ppasswd --db test
	`)
)
