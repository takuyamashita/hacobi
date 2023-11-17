package live_house_staff_usecase_test

import (
	"context"
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/repository/uuid_repository"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_staff_usecase"
)

func TestRegisterAccountSuccess(t *testing.T) {

	uuidRepository := uuid_repository.NewUuidRepository()
	store := NewStore(uuidRepository)
	liveHouseStaffRepositoryMock := LiveHouseStaffRepositoryMock{
		Store: &store,
	}

	liveHouseStaffEmailAddressChecker := live_house_staff_domain.NewLiveHouseStaffEmailAddressChecker(liveHouseStaffRepositoryMock)
	usecase := live_house_staff_usecase.NewLiveHouseStaffUsecase(uuidRepository, liveHouseStaffRepositoryMock, liveHouseStaffEmailAddressChecker)

	type args struct {
		name         string
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
				name:         "山田　太郎",
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
				name:         "",
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
				name:         "山田　太郎",
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
				name:         "山田　太郎",
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

			store.SetupStore(uuidRepository, t)

			_, err := usecase.RegisterAccount(tt.args.name, tt.args.emailAddress, tt.args.password, context.Background())
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
				if staff.Name().String() != tt.args.name {
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
