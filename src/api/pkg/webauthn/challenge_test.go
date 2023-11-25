package webauthn_test

import (
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/webauthn"
)

func TestChallenge(t *testing.T) {

	challenge, err := webauthn.NewChallenge()

	if err != nil {
		t.Fatal(err)
	}

	if len(challenge) != webauthn.ChallengeLength {
		t.Fatalf("invalid challenge length: %d", len(challenge))
	}

	if challenge.URLSafeBase64() == "" {
		t.Fatal("invalid challenge")
	}

}
