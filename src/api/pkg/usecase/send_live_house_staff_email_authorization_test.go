package usecase_test

import (
	"context"
	"testing"

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

			count := 0
			for _, v := range store.LiveHouseStaffEmailAuthorizations {
				if v.EmailAddress().String() == tt.args.emailAddress {
					count++
				}
			}

			if count != 1 {
				t.Errorf("SendLiveHouseStaffEmailAuthorizationUsecase.Execute() error = %v, wantErr %v", err, tt.want.hasError)
				return
			}
		})
	}
}
