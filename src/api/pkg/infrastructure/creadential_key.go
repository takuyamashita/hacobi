package infrastructure

import (
	"encoding/json"
	"io"

	"github.com/go-webauthn/webauthn/protocol"
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

func (c CredentialKeyImpl) ParseCredentialKey(body io.Reader) (*protocol.ParsedCredentialCreationData, error) {

	var ccr protocol.CredentialCreationResponse

	if err := json.NewDecoder(body).Decode(&ccr); err != nil {
		return nil, err
	}

	return ccr.Parse()
}
