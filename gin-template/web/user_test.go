package web_test

import (
	"net/http"
	"testing"

	"github.com/airdb-template/gin-api/web"
)

func TestListUser(t *testing.T) {
	uri := "/apis/user/v1/dean"
	resp := web.APIRequest(uri, "GET", nil)

	if resp.Code != http.StatusOK {
		t.Error(uri, resp.Code)
	}

	t.Log(uri, resp.Code)
	t.Log("resp:", resp.Body)
}
