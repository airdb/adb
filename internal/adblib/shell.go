package adblib

import (
	"github.com/MakeNowJust/heredoc"
)

var AwkDoc = heredoc.Doc(`
$ awk
awk '{sum[$2]+=1} END {for(k in sum) printf("%s %s %d %s\n", $1, k, sum[k], $NF)}' | sort -n -r -k 3

awk -F '|' '{sum += $1} END {print sum}'
	`)

var SedDoc = heredoc.Doc(`
$ awk
awk '{sum[$2]+=1} END {for(k in sum) printf("%s %s %d %s\n", $1, k, sum[k], $NF)}' | sort -n -r -k 3

awk -F '|' '{sum += $1} END {print sum}'
`)

var NginxDoc = heredoc.Doc(`
$ nginx
./auto/configure --add-module=/srv/nginx-ssl-ja3 --with-http_ssl_module --with-stream_ssl_module --with-debug --with-stream  --with-openssl=/srv/openssl
`)
