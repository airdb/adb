package cmd

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"

	"github.com/spf13/cobra"
)

const (
	passwdDefaultLength = 16
	passwdLowerCharset  = "abcdefghijklmnopqrstuvwxyz"
	passwdUpperCharset  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	passwdDigitCharset  = "0123456789"
	passwdSymbolCharset = "!@#$%^&*()-_=+[]{}:,.?"
)

var (
	passwdLength  int
	passwdComplex bool
	passwdReader  io.Reader = rand.Reader
)

var passwdCommand = &cobra.Command{
	Use:     "passwd",
	Aliases: []string{"password", "gen-passwd"},
	Short:   "Generate a random password",
	Long:    "Generate a random password. By default it returns a 16-character alphanumeric password. Use --complex to require upper, lower, digit, and symbol characters.",
	RunE: func(cmd *cobra.Command, args []string) error {
		password, err := generatePassword(passwdLength, passwdComplex)
		if err != nil {
			return err
		}

		fmt.Println(password)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(passwdCommand)
	passwdCommand.Flags().IntVarP(&passwdLength, "length", "n", passwdDefaultLength, "Password length")
	passwdCommand.Flags().BoolVarP(&passwdComplex, "complex", "c", false, "Require upper, lower, digit, and symbol characters")
}

func generatePassword(length int, complex bool) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("password length must be positive")
	}

	if !complex {
		return randomString(passwdLowerCharset+passwdUpperCharset+passwdDigitCharset, length)
	}

	requiredSets := []string{
		passwdLowerCharset,
		passwdUpperCharset,
		passwdDigitCharset,
	}
	if length < len(requiredSets)+1 {
		return "", fmt.Errorf("complex password length must be at least %d", len(requiredSets)+1)
	}

	nonSymbolChars := passwdLowerCharset + passwdUpperCharset + passwdDigitCharset
	password := make([]byte, 0, length-1)
	for _, charset := range requiredSets {
		ch, err := randomChar(charset)
		if err != nil {
			return "", err
		}
		password = append(password, ch)
	}

	extra, err := randomString(nonSymbolChars, length-len(requiredSets)-1)
	if err != nil {
		return "", err
	}
	password = append(password, extra...)

	if err := shuffleBytes(password); err != nil {
		return "", err
	}

	symbol, err := randomChar(passwdSymbolCharset)
	if err != nil {
		return "", err
	}

	return string(append(password, symbol)), nil
}

func randomString(charset string, length int) (string, error) {
	buf := make([]byte, length)
	for i := range buf {
		ch, err := randomChar(charset)
		if err != nil {
			return "", err
		}
		buf[i] = ch
	}
	return string(buf), nil
}

func randomChar(charset string) (byte, error) {
	n, err := randInt(len(charset))
	if err != nil {
		return 0, err
	}
	return charset[n], nil
}

func shuffleBytes(data []byte) error {
	for i := len(data) - 1; i > 0; i-- {
		j, err := randInt(i + 1)
		if err != nil {
			return err
		}
		data[i], data[j] = data[j], data[i]
	}
	return nil
}

func randInt(max int) (int, error) {
	if max <= 0 {
		return 0, fmt.Errorf("random bound must be positive")
	}

	n, err := rand.Int(passwdReader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}
