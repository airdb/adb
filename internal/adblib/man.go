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
	https://airdb.wiki/os/mac/mac/
`)

	GithubDoc = heredoc.Docf(`
	https://airdb.wiki/os/linux/shell/git.html
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

Check Client Hello:
	sudo ssldump -i  lo
	curl -k https://127.0.0.1:8443 | hexdump -C

Lego:
	https://github.com/go-acme/lego

	lego --email xxx@gmail.com --dns tencentcloud --domains=*.domain1.com --domains=*.domain2.com \
		renew --days=30 --renew-hook="nginx -s reload"
`)

	TcpdumpDoc = heredoc.Doc(`
	https://airdb.wiki/os/linux/shell/tcpdump.html
`)
	S3Doc = heredoc.Doc(`
# Minio
	wget https://dl.min.io/client/mc/release/linux-amd64/archive/mc.RELEASE.2019-10-02T19-41-02Z
	wget https://dl.min.io/client/mc/release/linux-amd64/mc
	mc config host add <bucketname> https://s3.github.com <accessKey> <secretKey>

# Tecent cos
	https://cloud.tencent.com/developer/article/1982033
	./mc-for-cos alias set s3 http://cos.ap-shanghai.myqcloud.com <TENCENTCLOUD_SECRET_ID> <TENCENTCLOUD_SECRET_KEY>
`)
)

var DockerDoc = heredoc.Doc(`
      $docker exec -it -e COLUMNS=$(tput cols) -e LINES=$(tput lines) airdb/go bash

      $ docker save myimage:latest | gzip > myimage_latest.tar.gz
      $ docker save -o fedora-all.tar fedora

      $ docker import /path/to/exampleimage.tgz
      $	sudo tar -c . | docker import --change "ENV DEBUG=true" - exampleimagedir

      podman

      brew install simnalamburt/x/podman-apple-silicon
      podman machine init --cpus=2 --disk-size=20 --memory 1000

      Refer: https://edofic.com/posts/2021-09-12-podman-m1-amd64/


      Error: short-name resolution enforced but cannot prompt without a TTYr
      podman machine ssh

      sed -i 's/short-name-mode="enforcing"/short-name-mode="permissive"/g' /etc/containers/registries.conf
      unqualified-search-registries = ["docker.io", "registry.fedoraproject.org", "registry.access.redhat.com", "quay.io"]
`)

var WebserverDoc = heredoc.Doc(`
webserver:
	$ python -m SimpleHTTPServer
	$ python3 -m http.server

	$ caddy run
`)
