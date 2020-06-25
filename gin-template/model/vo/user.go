package vo

import (
	"github.com/airdb-template/gin-api/model/po"
)

type UserReq struct {
	Code  string `form:"code"`
	State string `form:"state"`
}

type UserResp struct {
	Nickname   string `json:"nickname"`
	Headimgurl string `json:"headimgurl"`
	Token      string `json:"token"`
}

func FromPoUser(poUser *po.User) *UserResp {
	return &UserResp{
		Nickname:   "",
		Headimgurl: "",
		Token:      "",
	}
}

func List(voUser string) *UserResp {
	user := po.List(voUser)
	return FromPoUser(user)
}
