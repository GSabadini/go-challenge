package usecase

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/domain/vo"
)

type createUserRepoStub struct {
	result entity.User
	err    error
}

func (c createUserRepoStub) Create(context.Context, entity.User) (entity.User, error) {
	return c.result, c.err
}

type createUserPresenterStub struct {
	result CreateUserOutput
}

func (c createUserPresenterStub) Output(_ entity.User) CreateUserOutput {
	return c.result
}

func TestCreateUserInteractor_Execute(t *testing.T) {
	type fields struct {
		repo entity.CreateUserRepository
		pre  CreateUserPresenter
	}

	type args struct {
		ctx context.Context
		i   CreateUserInput
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    CreateUserOutput
		wantErr bool
	}{
		{
			name: "Create custom user success",
			fields: fields{
				repo: createUserRepoStub{
					result: entity.NewCustomUser(
						vo.NewUuidStaticTest(),
						vo.NewFullName("Test testing"),
						vo.Email{},
						vo.NewPassword("passw"),
						vo.NewDocumentTest(vo.CNPJ, "98.521.079/0001-09"),
						nil,
						time.Now(),
					),
					err: nil,
				},
				pre: createUserPresenterStub{
					result: CreateUserOutput{
						ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
						FullName: "Test testing",
						Document: CreateUserDocumentOutput{
							Type:  vo.CNPJ.String(),
							Value: "34018708000191",
						},
						Email:    "",
						Password: "passw",
						Wallet:   CreateUserWalletOutput{},
						Type:     vo.CUSTOM.String(),
					},
				},
			},
			args: args{
				ctx: nil,
				i: CreateUserInput{
					FullName: vo.NewFullName("Test testing"),
					Document: vo.NewDocumentTest(vo.CNPJ, "98.521.079/0001-09"),
					Email:    vo.Email{},
					Password: vo.NewPassword("passw"),
					Wallet:   nil,
					Type:     "CUSTOM",
				},
			},
			want: CreateUserOutput{
				ID:       "0db298eb-c8e7-4829-84b7-c1036b4f0791",
				FullName: "Test testing",
				Document: CreateUserDocumentOutput{
					Type:  vo.CNPJ.String(),
					Value: "34018708000191",
				},
				Email:    "",
				Password: "passw",
				Wallet:   CreateUserWalletOutput{},
				Type:     vo.CUSTOM.String(),
			},
			wantErr: false,
		},
		{
			name: "Create custom user error",
			fields: fields{
				repo: createUserRepoStub{
					result: entity.NewCustomUser(
						vo.NewUuidStaticTest(),
						vo.NewFullName("Test testing"),
						vo.Email{},
						vo.NewPassword("passw"),
						vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
						nil,
						time.Now(),
					),
					err: errors.New("fail"),
				},
				pre: createUserPresenterStub{
					result: CreateUserOutput{},
				},
			},
			args: args{
				ctx: nil,
				i: CreateUserInput{
					FullName: vo.NewFullName("Test testing"),
					Document: vo.NewDocumentTest(vo.CNPJ, "20.770.438/0001-66"),
					Email:    vo.Email{},
					Password: vo.NewPassword("passw"),
					Wallet:   nil,
					Type:     vo.CUSTOM,
				},
			},
			want:    CreateUserOutput{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := NewCreateUserInteractor(
				tt.fields.repo,
				tt.fields.pre,
			)

			got, err := c.Execute(tt.args.ctx, tt.args.i)
			if (err != nil) != tt.wantErr {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}
