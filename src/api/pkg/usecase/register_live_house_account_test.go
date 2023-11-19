package usecase_test

import (
	"context"
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

func TestRegisterLiveHouseAccount(t *testing.T) {

	store := NewStore()
	container := NewTestContainer(&store)

	var uuidRepo usecase.UuidRepositoryIntf
	container.Make(&uuidRepo)

	store.SetupStaffs(uuidRepo, t)

	uniqueStaffId, _ := uuidRepo.Generate()

	type args struct {
		id string
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
				id: store.Staffs[0].Id().String(),
			},
			want: want{
				hasError: false,
			},
		},
		{
			name: "存在しないスタッフIDで登録",
			args: args{
				id: uniqueStaffId,
			},
			want: want{
				hasError: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			id, err := usecase.RegisterLiveHouseAccount(tt.args.id, context.Background(), container)

			if tt.want.hasError && err == nil {
				t.Fatal("エラーが発生していません")
			}

			if tt.want.hasError {
				return
			}

			if err != nil {
				t.Fatal(err)
			}

			if id == "" {
				t.Fatal("id is empty")
			}

			exists := false
			var account live_house_account_domain.LiveHouseAccountIntf
			for _, a := range store.Accounts {
				if a.Id().String() == id {
					account = a
					exists = true
				}
			}
			if !exists {
				t.Fatal("アカウントが保存されていません")
			}

			if len(account.Staffs()) != 1 {
				t.Fatal("スタッフが登録されていないか、複数登録されています")
			}

			if !account.Staffs()[0].Role().Is(live_house_account_domain.GetRoleMaster()) {
				t.Fatalf("want role master, but set %d", account.Staffs()[0].Role().Number())
			}
		})
	}
}
