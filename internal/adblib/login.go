package adblib

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"

	"github.com/zitadel/oidc/v3/pkg/client/rp"
	httphelper "github.com/zitadel/oidc/v3/pkg/http"
)

func Login() error {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	clientID := ConfigNew.ClientID
	issuer := ConfigNew.AuthIssuer

	if clientID == "" || issuer == "" {
		return errors.New("auth_issuer and client_id must be set in ~/.config/adb/config.json")
	}

	clientSecret := os.Getenv("CLIENT_SECRET")
	keyPath := os.Getenv("KEY_PATH")
	scopes := strings.Split(os.Getenv("SCOPES"), " ")

	cookieKey := make([]byte, 16)
	if _, err := rand.Read(cookieKey); err != nil {
		return err
	}

	cookieHandler := httphelper.NewCookieHandler(cookieKey, cookieKey, httphelper.WithUnsecure())

	var options []rp.Option
	if clientSecret == "" {
		options = append(options, rp.WithPKCE(cookieHandler))
	}

	if keyPath != "" {
		options = append(options, rp.WithJWTProfile(rp.SignerFromKeyPath(keyPath)))
	}

	provider, err := rp.NewRelyingPartyOIDC(ctx, issuer, clientID, clientSecret, "", scopes, options...)
	if err != nil {
		return fmt.Errorf("create oidc provider: %w", err)
	}

	log.Println("starting device authorization flow")

	resp, err := rp.DeviceAuthorization(ctx, scopes, provider, nil)
	if err != nil {
		return err
	}

	fmt.Printf("\nPlease browse to %s and enter code %s\n", resp.VerificationURI, resp.UserCode)

	log.Println("start polling")

	token, err := rp.DeviceAccessToken(ctx, resp.DeviceCode, time.Duration(resp.Interval)*time.Second, provider)
	if err != nil {
		return err
	}

	log.Printf("successfully obtained token, type: %s, expires in: %ds", token.TokenType, token.ExpiresIn)

	return nil
}
