package infrastructure

import (
	"encoding/json"
	"io"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

type CredentialKeyIntf interface {
	usecase.CredentialKeyIntf
}

type CredentialKeyImpl struct {
}

func NewCredentialKey() CredentialKeyIntf {
	return &CredentialKeyImpl{}
}

func (c CredentialKeyImpl) CreateChallenge() (protocol.URLEncodedBase64, error) {

	return protocol.CreateChallenge()
}

func (c CredentialKeyImpl) CreateCredentialCreationOptions(challenge protocol.URLEncodedBase64, rpId string) protocol.PublicKeyCredentialCreationOptions {

	return protocol.PublicKeyCredentialCreationOptions{
		Challenge: challenge,
		RelyingParty: protocol.RelyingPartyEntity{
			CredentialEntity: protocol.CredentialEntity{
				Name: "localhost",
				Icon: "https://localhost/favicon.ico",
			},
		},
		User: protocol.UserEntity{
			ID:          []byte("1234567890"),
			DisplayName: "test-user",
		},
		CredentialExcludeList: []protocol.CredentialDescriptor{},
		Parameters: []protocol.CredentialParameter{
			{
				Type:      protocol.PublicKeyCredentialType,
				Algorithm: webauthncose.AlgES256,
			},
			{
				Type:      protocol.PublicKeyCredentialType,
				Algorithm: webauthncose.AlgRS256,
			},
		},
		Timeout: int((5 * time.Minute).Milliseconds()),
	}

}

func (c CredentialKeyImpl) ParseCredentialKey(body io.Reader) (*protocol.ParsedCredentialCreationData, error) {

	var ccr protocol.CredentialCreationResponse

	if err := json.NewDecoder(body).Decode(&ccr); err != nil {
		return nil, err
	}

	return ccr.Parse()
}
