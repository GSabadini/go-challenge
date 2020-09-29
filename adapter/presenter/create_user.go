package presenter

import (
	"github.com/GSabadini/go-challenge/domain/entity"
	"github.com/GSabadini/go-challenge/usecase"
)

type CreateUserPresenter struct{}

func (c CreateUserPresenter) Output(u entity.User) usecase.CreateUserOutput {
	return usecase.CreateUserOutput{
		ID:       u.ID().Value(),
		FullName: u.FullName().Value(),
		Document: usecase.CreateUserDocumentOutput{
			Type:  u.Document().Type().String(),
			Value: u.Document().Value(),
		},
		Email:    u.Email().Value(),
		Password: u.Password().Value(),
		Wallet: usecase.CreateUserWalletOutput{
			Currency: u.Wallet().Money().Currency().String(),
			Amount:   u.Wallet().Money().Amount().Value(),
		},
		Type: u.TypeUser().String(),
	}
}
