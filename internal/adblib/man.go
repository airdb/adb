package adblib

import (
	"github.com/MakeNowJust/heredoc"
)

var (
	MysqlDoc = heredoc.Doc(`
$ mysql
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
`)

	OpenSSLDoc = heredoc.Doc(`
$ openssl commands

	openssl x509 -text -in ssl.chain.crt
	openssl req  -noout -text -in ssl.csr
	openssl s_client -servername www.airdb.com -connect www.airdb.com:443 </dev/null 2>/dev/null

	cert -f md www.airdb.com
	Refer: https://github.com/genkiroid/cert
`)
)
