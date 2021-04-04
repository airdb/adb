package adblib

import (
	"github.com/MakeNowJust/heredoc"
)

var AwkDoc = heredoc.Doc(`
$ awk
awk '{sum[$2]+=1} END {for(k in sum) printf("%s %s %d %s\n", $1, k, sum[k], $NF)}' | sort -n -r -k 3
	`)

var SedDoc = heredoc.Doc(`
$ awk
awk '{sum[$2]+=1} END {for(k in sum) printf("%s %s %d %s\n", $1, k, sum[k], $NF)}' | sort -n -r -k 3
	`)