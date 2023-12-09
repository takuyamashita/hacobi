package live_house_staff_account_domain

import (
	"errors"
	"time"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

type ProvisionalRegistrationParam struct {
	Token     string
	CreatedAt time.Time
}

type CredentialChallengeParam struct {
	Challenge string
	CreatedAt time.Time
}

type ProfileParam struct {
	DisplayName string
}

type NewLiveHouseStaffAccountParams struct {
	Id                        string
	Email                     string
	IsProvisional             bool
	ProvisionalRegistration   *ProvisionalRegistrationParam
	CredentialChallengeParams *CredentialChallengeParam
	CredentialKeys            []domain.PublicKeyId
	ProfileParams             *ProfileParam
}

func NewLiveHouseStaffAccount(params NewLiveHouseStaffAccountParams) (account LiveHouseStaffAccountIntf, err error) {

	if !params.IsProvisional && params.ProvisionalRegistration != nil {
		return nil, errors.New("仮登録でない場合は、ProvisionalRegistrationはnilである必要があります。")
	}
	if params.IsProvisional {
		if params.ProvisionalRegistration == nil {
			return nil, errors.New("仮登録の場合は、ProvisionalRegistrationは必須です。")
		}
		if params.ProvisionalRegistration.Token == "" {
			return nil, errors.New("仮登録の場合は、ProvisionalRegistration.Tokenは必須です。")
		}
		if params.ProvisionalRegistration.CreatedAt.IsZero() {
			return nil, errors.New("仮登録の場合は、ProvisionalRegistration.CreatedAtは必須です。")
		}
	}

	id := NewLiveHouseStaffAccountId(params.Id)

	email, err := domain.NewLiveHouseStaffEmailAddress(params.Email)
	if err != nil {
		return nil, err
	}

	var challenge CredentialChallengeIntf

	if params.CredentialChallengeParams != nil {
		challenge, err = NewCredentialChallenge(params.CredentialChallengeParams.Challenge, params.CredentialChallengeParams.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	if params.ProvisionalRegistration == nil {
		return &LiveHouseStaffAccountImpl{
			id:                      id,
			email:                   email,
			isProvisional:           params.IsProvisional,
			credentialChallenge:     challenge,
			provisionalRegistration: nil,
		}, nil
	}

	provisionalRegistration, err := NewLiveHouseStaffAccountProvisionalRegistration(
		params.ProvisionalRegistration.Token,
		params.ProvisionalRegistration.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	var profile LiveHouseStaffProfileIntf
	if params.ProfileParams != nil {
		profile, err = NewLiveHouseStaffProfile(params.ProfileParams.DisplayName)
		if err != nil {
			return nil, err
		}
	}

	return &LiveHouseStaffAccountImpl{
		id:                      id,
		email:                   email,
		isProvisional:           params.IsProvisional,
		provisionalRegistration: provisionalRegistration,
		credentialChallenge:     challenge,
		credentialKeys:          params.CredentialKeys,
		profile:                 profile,
	}, nil
}
