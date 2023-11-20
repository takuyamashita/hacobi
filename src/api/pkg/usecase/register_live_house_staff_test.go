package usecase_test

import (
	"context"
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

func TestRegisterLiveHouseStaff(t *testing.T) {

	store := NewStore()
	container := NewTestContainer(&store)

	type args struct {
		displayName  string
		emailAddress string
		password     string
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
			name: "正常に保存",
			args: args{
				displayName:  "山田　太郎",
				emailAddress: "test@test.com",
				password:     "123password^{}$s",
			},
			want: want{
				hasError: false,
			},
		},
		{
			name: "名前が空文字",
			args: args{
				displayName:  "",
				emailAddress: "test@test.com",
				password:     "123password^{}$s",
			},
			want: want{
				hasError: true,
			},
		},
		{
			name: "メールアドレスが空文字",
			args: args{
				displayName:  "山田　太郎",
				emailAddress: "",
				password:     "123password^{}$s",
			},
			want: want{
				hasError: true,
			},
		},
		{
			name: "メールアドレスが重複",
			args: args{
				displayName:  "山田　太郎",
				emailAddress: "duplicate@test.com",
				password:     "123password^{}$s",
			},
			want: want{
				hasError: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			var uuidRepo usecase.UuidRepositoryIntf
			container.Make(&uuidRepo)

			store.SetupStaffs(uuidRepo, t)
			_, err := usecase.RegisterLiveHouseStaff(tt.args.displayName, tt.args.emailAddress, tt.args.password, context.Background(), container)
			if tt.want.hasError && err == nil {
				t.Errorf("エラーが発生していません")
			}

			if tt.want.hasError {
				return
			}

			// xxx: passwordのhash化ができているかテストする

			// 存在チェック
			staffExists := false
			for _, staff := range store.Staffs {
				if staff.EmailAddress().String() != tt.args.emailAddress {
					continue
				}
				if staff.DisplayName().String() != tt.args.displayName {
					continue
				}
				staffExists = true
				break
			}

			if !staffExists {
				t.Error("スタッフが保存されていません")
			}
		})
		store.Clear()
	}

}
