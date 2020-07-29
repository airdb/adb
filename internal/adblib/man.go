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
)
