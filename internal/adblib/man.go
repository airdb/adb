package adblib

import (
	"github.com/MakeNowJust/heredoc"
)

var (
	MysqlDoc = heredoc.Doc(`
$ mysql
Backup table // https://stackoverflow.com/questions/40724046/mysql-gtid-consistency-violation
> CREATE TABLE new_table LIKE old_table; 
> INSERT new_table SELECT * FROM old_table;

> ALTER TABLE xx_tab change status process int(10)
> ALTER TABLE xx_tab modify status int(10) after id
> ALTER TABLE xx_tab modify status varchar(600) NOT NULL

SELECT DISTINCT column1, column2 FROM table_name;
	`)

	ToolsDoc = heredoc.Doc(`
$ tools
	asciinema play https://asciinema.org/a/349488	
	docker run --rm -v $PWD:/data asciinema/asciicast2gif -s 2 -t solarized-dark demo.json demo.gif

	docker run --rm -v $PWD:/data asciinema/asciicast2gif -s 2 -t solarized-dark demo.json demo.gif
`)

	HelmDoc = heredoc.Doc(`
$ helm commands

	sudo apt update
	sudo apt install snapd
	sudo snap install helm --classic

	helm plugin install https://github.com/airdb/helm-kube
	
	helm repo add airdb https://www.airdb.com/helm/"
	helm repo update
	helm search repo helm/mychart
	
	helm show readme airdb/mychat
	
	helm install chart airdb/mychat
	helm install chart airdb/mychat --dry-run --debug
	
		helm get notes mychat
`)

	TerraformDoc = heredoc.Doc(`
$ terraform commands

	terraform -install-autocomplete
	terraform init -upgrade
	
	terraform validate
	terraform plan
	terraform apply
	#terraform destroy
	
	Refer: https://github.com/airdb/init/tree/master/terraform
`)

	BrewDoc = heredoc.Docf(`
$ Brew Common Command:

	brew outdated

	brew cask outdated
	brew outdated adb --verbose --debug
	brew install github/gh/gh
	brew install aliyun-cli
	
	brew tap aidb/taps
	brew install airdb/taps/adb
	brew install adb

`)

	GithubDoc = heredoc.Docf(`
$ Github or Git Command:

	1. Maintain a repo without permission
		git remote add upstream https://github.com/bfenetworks/bfe.git
		git fetch upstream
		git checkout develop

		git merge upstream/develop

		git rebase upstream/develop

		Refer: https://docs.github.com/en/github/collaborating-with-issues-and-pull-requests/syncing-a-fork


	2. git config setting
		git config --global core.hooksPath .github/hooks
		git config --global core.excludefile .gitignore_global
		git config --global url.'ssh://git@github.com'.insteadOf https://github.com

	3. Github Commit. For Close Github Issue, commit message should as follow:
		close #x
		closes #x
		closed #x
		fix #x
		fixes #x
		fixed #x
		resolve #x
		resolves #x
		resolved #x
		add new quick sort algorithm, fixes #4, resolve #6, closed #12

	4. Github Command Line
		brew install github/gh/gh
		gh --repo bbhj/lbs issue status
		gh --repo bbhj/lbs issue view 1

	5. Delete branch or tag
		git branch -D dev
		git push origin --delete dev

		git tag -d v1.0
		git push --delete origin v1.0

`)

	OpenSSLDoc = heredoc.Doc(`
$ openssl commands

	openssl x509 -text -in ssl.chain.crt
	openssl req  -noout -text -in ssl.csr
	openssl s_client -servername www.airdb.com -connect www.airdb.com:443 </dev/null 2>/dev/null

	Check Keypair 1
	openssl rsa -pubout -in privkey.pem
	openssl x509 -pubkey -noout -in fullchain.pem

	Check keypair 2
	diff -eq <(openssl x509 -pubkey -noout -in fullchain.pem) <(openssl rsa -pubout -in privkey.pem)

	cert -f md www.airdb.com
	Refer: https://github.com/genkiroid/cert
`)

	TcpdumpDoc = heredoc.Doc(`
$ tcpdump commands

	sudo tcpdump -i bond0.1000  -nnAAAA  | grep -A 20 -B 3  http-gateway.spex.test.shopee.sg
`)
	S3Doc = heredoc.Doc(`
# Minio
	wget https://dl.min.io/client/mc/release/linux-amd64/mc
	mc config host add <bucketname> https://s3.airdb.io <accessKey> <secretKey>
`)

	PerformanceDoc = heredoc.Doc(`
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

	Refer: https://www.brendangregg.com/linuxperf.html

$ Golang

	1. go tool pprof -alloc_space http://127.0.0.1:8080/debug/pprof/heap
		top

	2. go tool pprof -alloc_space -cum -svg http://127.0.0.1:8080/debug/pprof/heap > heap.svg
		(apt-get  install graphviz)
`)
)
