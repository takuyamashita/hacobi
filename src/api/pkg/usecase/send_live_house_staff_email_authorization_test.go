package usecase_test

import (
	"context"
	"strings"
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

func TestSendLiveHouseStaffAccountProvisionalRegistration(t *testing.T) {

	store := NewStore()
	container := NewTestContainer(&store)

	type args struct {
		emailAddress string
	}
	type want struct {
		hasError bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "正常に送信",
			args: args{
				emailAddress: "test@test.com",
			},
			want: want{
				hasError: false,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			ctx := context.Background()

			err := usecase.SendLiveHouseStaffEmailAuthorization(tt.args.emailAddress, ctx, container)
			if (err != nil) != tt.want.hasError {
				t.Errorf("SendLiveHouseStaffAccountProvisionalRegistrationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}

			account := []live_house_staff_account_domain.LiveHouseStaffAccountIntf{}
			for _, v := range store.LiveHouseStaffAccounts {
				if v.EmailAddress().String() == tt.args.emailAddress {
					account = append(account, v)
				}
			}

			if len(account) != 1 {
				t.Fatalf("SendLiveHouseStaffAccountProvisionalRegistrationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}

			mails := []SentMail{}
			for _, v := range store.SentMails {
				if v.To == tt.args.emailAddress {
					mails = append(mails, v)
				}
			}

			if len(mails) != 1 {
				t.Fatalf("SendLiveHouseStaffAccountProvisionalRegistrationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}

			if mails[0].To != tt.args.emailAddress {
				t.Fatalf("SendLiveHouseStaffAccountProvisionalRegistrationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}

			if mails[0].Subject != "認証メール" {
				t.Fatalf("SendLiveHouseStaffAccountProvisionalRegistrationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}

			if !strings.Contains(mails[0].Body, account[0].ProvisionalToken().String()) {
				t.Fatalf("SendLiveHouseStaffAccountProvisionalRegistrationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}
		})
	}
}
