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
			in:             "eyJhbGciOiJIUzI1NiIsImtpZCI6IkNaMThKWEhCUTBMSzk5ZnZGeEV1eUhQVmdMMUw3dHJTIiwidHlwIjoiSldUIn0.eyJleHAiOjE2MjYwNzkwNDB9.GrC5Rv-P1y5L3b_UAeP4pLUWEnROEaLE8s4JGhr4fVA",
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
