package usecase_test

import (
	"context"
	"strings"
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_email_authorization_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

func TestSendLiveHouseStaffEmailAuthorization(t *testing.T) {

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
				t.Errorf("SendLiveHouseStaffEmailAuthorizationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}

			auth := []live_house_staff_email_authorization_domain.LiveHouseStaffEmailAuthorizationIntf{}
			for _, v := range store.LiveHouseStaffEmailAuthorizations {
				if v.EmailAddress().String() == tt.args.emailAddress {
					auth = append(auth, v)
				}
			}

			if len(auth) != 1 {
				t.Fatalf("SendLiveHouseStaffEmailAuthorizationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}

			mails := []SentMail{}
			for _, v := range store.SentMails {
				if v.To == tt.args.emailAddress {
					mails = append(mails, v)
				}
			}

			if len(mails) != 1 {
				t.Fatalf("SendLiveHouseStaffEmailAuthorizationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}

			if mails[0].To != tt.args.emailAddress {
				t.Fatalf("SendLiveHouseStaffEmailAuthorizationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}

			if mails[0].Subject != "認証メール" {
				t.Fatalf("SendLiveHouseStaffEmailAuthorizationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}

			if !strings.Contains(mails[0].Body, auth[0].Token().String()) {
				t.Fatalf("SendLiveHouseStaffEmailAuthorizationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}
		})
	}
}
