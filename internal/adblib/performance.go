package adblib

import "github.com/MakeNowJust/heredoc"

var PerformanceDoc = heredoc.Doc(`
$ Linux

	1. dstat
	2. iostat
		iostat -x 5

	3. iotop
		iotop -p $pid

	4. top, htop
	5. iptraf
	6. iftop

	7. lsof

	8. strace
		strace -p $pid

	9. perf
		perf top  --cpu=0
		perf top  --cpu --pid=407595

	Memory: https://colobu.com/2019/08/28/go-memory-leak-i-dont-think-so/

	Refer:
	https://www.brendangregg.com/linuxperf.html
	https://queue.acm.org/detail.cfm?id=2927301
	https://www.usenix.org/sites/default/files/conference/protected-files/gregg_lisa13_flamegraphs.pdf
	
	https://gnuser.github.io/lpo/io/stack
	https://gnuser.github.io/lpo/index

$ Golang

	1. go tool pprof -alloc_space http://127.0.0.1:8080/debug/pprof/heap
		top

	2. go tool pprof -alloc_space -cum -svg http://127.0.0.1:8080/debug/pprof/heap > heap.svg
		(apt-get  install graphviz)
`)


var WrkDoc = heredoc.Doc(`
$ wrk
config.lua:
request = function()
  wrk.method = "POST"
  wrk.headers["Content-Type"] = "application/json"
  wrk.headers["Authorization"] = "Bearer xx"
  wrk.body = "{}"

  return wrk.format("POST", headers)
end

wrk -t4 -c100 -d30s -T30s --script=config.lua --latency https://airdb.io

`)
