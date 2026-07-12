package adblib

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const maxKeysPerUser = 3

var githubClient = &http.Client{Timeout: 10 * time.Second}

func GetGithubKeysByID(id string) error {
	keysURL := fmt.Sprintf("https://github.com/%s.keys", id)

	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, keysURL, nil)
	if err != nil {
		return err
	}

	resp, err := githubClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("fetch %s: unexpected status %s", keysURL, resp.Status)
	}

	count := 0
	scanner := bufio.NewScanner(resp.Body)

	for scanner.Scan() && count < maxKeysPerUser {
		line := scanner.Text()
		if !strings.HasPrefix(line, "ssh-ed25519") {
			continue
		}

		fmt.Printf("%s %s\n", line, keysURL)
		count++
	}

	return scanner.Err()
}

func GetGithubKeys(userID []string) error {
	var errs []error

	for _, id := range userID {
		if err := GetGithubKeysByID(strings.TrimSpace(id)); err != nil {
			errs = append(errs, err)
		}
	}

	return errors.Join(errs...)
}
