package service

import (
	"context"

	"github.com/symptomatichq/kit/pgutil"
)

type Repository interface {
	InsertAccount(context.Context, Account) (Account, error)
	GetAccountByID(context.Context, string) (Account, error)
	SelectAccounts(context.Context, map[string]interface{}) ([]Account, error)
	InsertUser(context.Context, User) (User, error)
	GetUserByID(context.Context, string) (User, error)
	SelectUsers(context.Context, map[string]interface{}) ([]User, error)
}

func NewRepository(dbConfig *pgutil.ConnectionOptions) Repository {
	return &repository{}
}

type repository struct{}

func (r *repository) InsertAccount(ctx context.Context, newAccount Account) (account Account, err error) {
	return
}

func (r *repository) GetAccountByID(ctx context.Context, id string) (account Account, err error) {
	return
}

func (r *repository) SelectAccounts(ctx context.Context, filters map[string]interface{}) (account []Account, err error) {
	return
}

func (r *repository) InsertUser(ctx context.Context, newUser User) (user User, err error) {
	return

}

func (r *repository) GetUserByID(ctx context.Context, id string) (user User, err error) {
	return

}

func (r *repository) SelectUsers(ctx context.Context, filters map[string]interface{}) (users []User, err error) {
	return
}
