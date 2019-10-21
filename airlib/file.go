package airlib

import (
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/imroc/req"
)

func AppendToFile(filename, key string) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	defer f.Close()
	cbyte, _ := ioutil.ReadFile(filename)
	content := string(cbyte)

	if !strings.Contains(content, key) {
		if _, err := f.WriteString(key); err != nil {
			log.Println(err)
		}
	}


}

func SetupSSHPublicKey() {
	url := "https://init.airdb.host/osinit/authorized_keys"
	r, _ := req.Get(url)

	filename := os.Getenv("HOME") + "/.ssh/authorized_keys"
	AppendToFile(filename, r.String())
}