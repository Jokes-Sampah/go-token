package gotoken

import (
	"testing"
	"time"
)

const YourSecretKey = "SecretKey"

func TestHandlers_Handler(t *testing.T) {
	tests := []struct {
		name           string
		in             string
		secretKey      string
	}{
		// Test Case List
		{
			name:           "Test #1",
			in:             "YourToken",
			secretKey:		YourSecretKey,
		},
	}

	for _, test := range tests {
		test := test
		token, err := ExtractToken(test.in, test.secretKey)
		t.Run(test.name, func(t *testing.T) {
			if err != nil {
				t.Logf("expected: Success\ngot: %s\n", err)
				t.Fail()
			}
			t.Logf("Is expired: %v", token.ExpiresAt < time.Now().Unix())
			t.Logf("✔ '%s' PASSED\n", test.name)
		})
	}
}
