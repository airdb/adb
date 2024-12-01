package adblib

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/AlecAivazis/survey/v2"
	"github.com/airdb/sailor/osutil"
	"github.com/airdb/sailor/process"
	"github.com/imroc/req"

	"github.com/zitadel/oidc/v3/pkg/client/rp"
	httphelper "github.com/zitadel/oidc/v3/pkg/http"
)

// The  flag will be written to this struct.
type User struct {
	Name     string `json:"name" survey:"name" mapstructure:"name"`
	Password string `json:"password" survey:"Password" mapstructure:"password"`
}

func LoginWithToken() {
	fmt.Print(TokenRequest)

	var bar process.Bar

	bar.NewOption(0, 100)

	sleepInterval := 100

	for i := 0; i <= 100; i++ {
		time.Sleep(time.Duration(sleepInterval) * time.Millisecond)
		bar.Play(int64(i))
	}

	bar.Finish()
}

// The questions to ask.
var userLogin = []*survey.Question{
	{
		Name:     "name",
		Prompt:   &survey.Input{Message: "Username:"},
		Validate: survey.Required,
	},
	{
		Name:     "password",
		Prompt:   &survey.Password{Message: "Password:"},
		Validate: survey.Required,
	},
}

var (
	key = []byte("test1234test1234")
)

func Login() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT)
	defer stop()

	fmt.Println(ConfigNew.CLIENT_ID, ConfigNew.AuthIssuer)
	clientID := ConfigNew.CLIENT_ID
	clientSecret := os.Getenv("CLIENT_SECRET")
	keyPath := os.Getenv("KEY_PATH")
	issuer := ConfigNew.AuthIssuer
	scopes := strings.Split(os.Getenv("SCOPES"), " ")

	cookieHandler := httphelper.NewCookieHandler(key, key, httphelper.WithUnsecure())

	var options []rp.Option
	if clientSecret == "" {
		options = append(options, rp.WithPKCE(cookieHandler))
	}
	if keyPath != "" {
		options = append(options, rp.WithJWTProfile(rp.SignerFromKeyPath(keyPath)))
	}

	provider, err := rp.NewRelyingPartyOIDC(ctx, issuer, clientID, clientSecret, "", scopes, options...)
	if err != nil {
		log.Fatalf("error creating provider %s", err.Error())
	}

	log.Println("starting device authorization flow")
	resp, err := rp.DeviceAuthorization(ctx, scopes, provider, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("resp", resp)
	fmt.Printf("\nPlease browse to %s and enter code %s\n", resp.VerificationURI, resp.UserCode)

	log.Println("start polling")
	token, err := rp.DeviceAccessToken(ctx, resp.DeviceCode, time.Duration(resp.Interval)*time.Second, provider)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("successfully obtained token: %#v", token)
}

func Logo() {
	args := []string{IconFile()}

	out, err := osutil.ExecCommand("cat", args)
	if err != nil {
		downloadIcon()

		return
	}

	fmt.Println(out)
}

func downloadIcon() {
	mtod := "https://init.airdb.host/mtod/icon"
	r, err := req.Get(mtod)

	if err == nil {
		msg, _ := r.ToString()
		fmt.Print(msg)

		if err = r.ToFile(IconFile()); err == nil {
			return
		}
	}
}
