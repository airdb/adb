package cmd

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

var jwtCommand = &cobra.Command{
	Use:   "jwt <jwt-token>",
	Short: "Decode a JWT payload without verifying its signature",
	Long:  "Decode a JWT payload and print it as formatted JSON without verifying its signature. Accepts an optional Bearer prefix.",
	Example: "  adb jwt \"eyJhbGciOiJIUzI1NiJ9.eyJzdWIiOiIxMjM0NTYiLCJuYW1lIjoiRGVhbiIsImFkbWluIjp0cnVlfQ.signature\"\n" +
		"  adb jwt 'Bearer header.payload.signature'",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) == 0 {
			return cmd.Help()
		}

		claims, err := DecodeJWT(args[0])
		if err != nil {
			return err
		}

		prettyJSON, err := json.MarshalIndent(claims, "", "  ")
		if err != nil {
			return err
		}

		_, err = fmt.Fprintln(cmd.OutOrStdout(), string(prettyJSON))
		return err
	},
}

func init() {
	rootCmd.AddCommand(jwtCommand)
}

// DecodeJWT decodes the JWT payload without verifying its signature.
func DecodeJWT(token string) (map[string]interface{}, error) {
	token = strings.TrimSpace(token)
	token = strings.TrimPrefix(token, "Bearer ")

	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid JWT: expected 3 parts, got %d", len(parts))
	}

	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, fmt.Errorf("decode payload: %w", err)
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return nil, fmt.Errorf("parse JSON: %w", err)
	}

	return claims, nil
}
