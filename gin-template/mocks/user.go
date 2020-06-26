package mocks

import (
	"{{ .GoModulePath }}/model/po"
)

var User1 = &po.User{
	Nickname:   "test",
	Headimgurl: "https://airdb.bio/test",
	Token:      "token_base64",
}
