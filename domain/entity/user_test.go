package entity

import (
	"github.com/GSabadini/golang-clean-architecture/domain/vo"
	"reflect"
	"testing"
	"time"
)

func TestNewUser(t *testing.T) {
	type args struct {
		ID        vo.Uuid
		fullName  vo.FullName
		email     vo.Email
		password  vo.Password
		document  vo.Document
		wallet    *vo.Wallet
		typeUser  vo.TypeUser
		createdAt time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr error
	}{
		{
			name: "Test create common user",
			args: args{
				ID:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.NewDocumentTest(vo.CPF, "07010965836"),
				wallet:    vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
				typeUser:  "COMMON",
				createdAt: time.Time{},
			},
			want: User{
				id:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.NewDocumentTest(vo.CPF, "07010965836"),
				roles:     vo.Roles{CanTransfer: true},
				wallet:    vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
				typeUser:  vo.COMMON,
				createdAt: time.Time{},
			},
		},
		{
			name: "Test create merchant user",
			args: args{
				ID:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.NewDocumentTest(vo.CNPJ, "90.691.635/0001-75"),
				wallet:    vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
				typeUser:  "MERCHANT",
				createdAt: time.Time{},
			},
			want: User{
				id:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.NewDocumentTest(vo.CNPJ, "90.691.635/0001-75"),
				roles:     vo.Roles{CanTransfer: false},
				wallet:    vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
				typeUser:  vo.MERCHANT,
				createdAt: time.Time{},
			},
		},
		{
			name: "Test create invalid user",
			args: args{
				ID:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.NewDocumentTest(vo.CNPJ, "07010965836"),
				wallet:    vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
				typeUser:  "INVALID",
				createdAt: time.Time{},
			},
			want:    User{},
			wantErr: vo.ErrInvalidTypeUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(
				tt.args.ID,
				tt.args.fullName,
				tt.args.email,
				tt.args.password,
				tt.args.document,
				tt.args.wallet,
				tt.args.typeUser,
				tt.args.createdAt,
			)
			if (err != nil) && (tt.wantErr != err) {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestUser_CanTransfer(t *testing.T) {
	type args struct {
		id        vo.Uuid
		fullName  vo.FullName
		email     vo.Email
		password  vo.Password
		document  vo.Document
		wallet    *vo.Wallet
		typeUser  vo.TypeUser
		roles     vo.Roles
		createdAt time.Time
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Test whether common type user can transfer",
			args: args{
				id:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.Document{},
				wallet:    nil,
				typeUser:  vo.COMMON,
				roles:     vo.Roles{},
				createdAt: time.Time{},
			},
			want: nil,
		},
		{
			name: "Test whether merchant type user can transfer",
			args: args{
				id:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.Document{},
				wallet:    nil,
				typeUser:  vo.MERCHANT,
				roles:     vo.Roles{},
				createdAt: time.Time{},
			},
			want: vo.ErrNotAllowedTypeUser,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(
				tt.args.id,
				tt.args.fullName,
				tt.args.email,
				tt.args.password,
				tt.args.document,
				tt.args.wallet,
				tt.args.typeUser,
				tt.args.createdAt,
			)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v", tt.name, err)
				return
			}

			if err := got.CanTransfer(); err != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.want)
			}
		})
	}
}

func TestUser_Deposit(t *testing.T) {
	type argsUser struct {
		id        vo.Uuid
		fullName  vo.FullName
		email     vo.Email
		password  vo.Password
		document  vo.Document
		wallet    *vo.Wallet
		typeUser  vo.TypeUser
		roles     vo.Roles
		createdAt time.Time
	}
	type args struct {
		money vo.Money
	}
	tests := []struct {
		name     string
		argsUser argsUser
		args     args
		want     int64
	}{
		{
			name: "Test deposit 100",
			argsUser: argsUser{
				id:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.Document{},
				wallet:    vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
				typeUser:  vo.COMMON,
				roles:     vo.Roles{},
				createdAt: time.Time{},
			},
			args: args{
				money: vo.NewMoneyBRL(vo.NewAmountTest(100)),
			},
			want: 200,
		},
		{
			name: "Test deposit 1000",
			argsUser: argsUser{
				id:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.Document{},
				wallet:    vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
				typeUser:  vo.COMMON,
				roles:     vo.Roles{},
				createdAt: time.Time{},
			},
			args: args{
				money: vo.NewMoneyBRL(vo.NewAmountTest(1000)),
			},
			want: 1100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(
				tt.argsUser.id,
				tt.argsUser.fullName,
				tt.argsUser.email,
				tt.argsUser.password,
				tt.argsUser.document,
				tt.argsUser.wallet,
				tt.argsUser.typeUser,
				tt.argsUser.createdAt,
			)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v", tt.name, err)
				return
			}

			got.Deposit(tt.args.money)

			if got.Wallet().Money().Amount().Value() != tt.want {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got.Wallet().Money().Amount(), tt.want)
			}
		})
	}
}

func TestUser_Withdraw(t *testing.T) {
	type argsUser struct {
		id        vo.Uuid
		fullName  vo.FullName
		email     vo.Email
		password  vo.Password
		document  vo.Document
		wallet    *vo.Wallet
		typeUser  vo.TypeUser
		roles     vo.Roles
		createdAt time.Time
	}
	type args struct {
		money vo.Money
	}
	tests := []struct {
		name     string
		argsUser argsUser
		args     args
		want     int64
		wantErr  error
	}{
		{
			name: "Test withdraw 100",
			argsUser: argsUser{
				id:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.Document{},
				wallet:    vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
				typeUser:  vo.COMMON,
				roles:     vo.Roles{},
				createdAt: time.Time{},
			},
			args: args{
				money: vo.NewMoneyBRL(vo.NewAmountTest(100)),
			},
			want: 0,
		},
		{
			name: "Test withdraw 50",
			argsUser: argsUser{
				id:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.Document{},
				wallet:    vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
				typeUser:  vo.COMMON,
				roles:     vo.Roles{},
				createdAt: time.Time{},
			},
			args: args{
				money: vo.NewMoneyBRL(vo.NewAmountTest(50)),
			},
			want: 50,
		},
		{
			name: "Test withdraw insufficient balance",
			argsUser: argsUser{
				id:        vo.NewUuidStaticTest(),
				fullName:  vo.NewFullName("Test testing"),
				email:     vo.Email{},
				password:  vo.NewPassword("123"),
				document:  vo.Document{},
				wallet:    vo.NewWallet(vo.NewMoneyBRL(vo.NewAmountTest(100))),
				typeUser:  vo.COMMON,
				roles:     vo.Roles{},
				createdAt: time.Time{},
			},
			args: args{
				money: vo.NewMoneyBRL(vo.NewAmountTest(1000)),
			},
			wantErr: ErrUserInsufficientBalance,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewUser(
				tt.argsUser.id,
				tt.argsUser.fullName,
				tt.argsUser.email,
				tt.argsUser.password,
				tt.argsUser.document,
				tt.argsUser.wallet,
				tt.argsUser.typeUser,
				tt.argsUser.createdAt,
			)
			if err != nil {
				t.Errorf("[TestCase '%s'] Err: '%v", tt.name, err)
				return
			}

			err = got.Withdraw(tt.args.money)
			if (err != nil) && (tt.wantErr != err) {
				t.Errorf("[TestCase '%s'] Err: '%v' | WantErr: '%v'", tt.name, err, tt.wantErr)
				return
			}

			if (err == nil) && (got.Wallet().Money().Amount().Value() != tt.want) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got.Wallet().Money().Amount(), tt.want)
			}
		})
	}
}
