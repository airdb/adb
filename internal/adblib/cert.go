package adblib

import (
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	"time"

	"github.com/airdb/sailor/fileutil"
	"github.com/pkg/errors"
)

type CertX509 struct {
	KeyFile    string
	ChainFile  string
	Issuer     string
	CommonName string
	Sans       []string
	NotBefore  time.Time
	NotAfter   time.Time
}

var CaddyTmpl = `
{{ .Sans }} {
    #proxy / gogs:3000 {
    #    header_upstream Host {host}
    #    header_upstream X-Real-IP {remote}
    #    header_upstream X-Forwarded-For {remote}
    #    header_upstream X-Forwarded-Proto {scheme}
    #}

    #log /tmp/caddy.log
    root * /tmp/
    file_server browse

    tls {{ .ChainFile }} {{ .KeyFile }}
}
`

func HandlerCert(keyfile, chainFile string) {
	chainStr, _ := fileutil.ReadFile(chainFile)

	x, _ := ParseCertChain(chainStr)

	x.KeyFile = keyfile
	x.ChainFile = chainFile

	ret, _ := fileutil.TemplateGenerateString(CaddyTmpl, x)
	fmt.Println("xxx", ret)

}

func ParseCertChain(chain []byte) (*CertX509, error) {
	pemBlock, _ := pem.Decode(chain)
	if pemBlock == nil {
		err := errors.Errorf("cert pem decode err")
		log.Println("cert pem decode fail, err:", err)

		return nil, err
	}

	cert, err := x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		log.Println("cert pem parse fail, err:", err)

		return nil, err
	}

	certInfo := CertX509{
		Issuer:     cert.Issuer.CommonName,
		CommonName: cert.Subject.CommonName,
		Sans:       cert.DNSNames,
		NotBefore:  cert.NotBefore,
		NotAfter:   cert.NotAfter,
	}

	return &certInfo, nil
}
