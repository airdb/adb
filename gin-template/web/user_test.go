package web

import (
	"net/http"
	"testing"
)

func TestListUser(t *testing.T) {
	uri := "/apis/user/v1/dean"
	resp := APIRequest(uri, "GET", nil)

	if resp.Code != http.StatusOK {
		t.Error(uri, resp.Code)
	}

	t.Log(uri, resp.Code)
	t.Log("resp:", resp.Body)
}
