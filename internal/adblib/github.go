package adblib

import (
	"bufio"
	"fmt"
	"net/http"
)

func GetGithubKeysByID(id string) {
	keysURL := fmt.Sprintf("https://github.com/%s.keys", id)
	resp, err := http.Get(keysURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 3; i++ {
		fmt.Printf("%s https://github.com/%s\n", scanner.Text(), id)
	}
}

func GetGithubKeys(userID []string) {
	for _, id := range userID {
		GetGithubKeysByID(id)
	}
}
