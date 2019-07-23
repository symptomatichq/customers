package service

import (
	"context"
	"time"

	"github.com/go-kit/kit/log"

	"github.com/symptomatichq/kit/logutil"
)

type AccountStatus string

const (
	AccountActive    AccountStatus = "active"
	AccountSuspended AccountStatus = "suspended"
	AccountInactive  AccountStatus = "inactive"
)

type Account struct {
	ID           string        `db:"id"`
	Name         string        `db:"name"`
	ContactEmail string        `db:"contact_email"`
	Status       AccountStatus `db:"status"`
	UpdatedAt    time.Time     `db:"updated_at"`
	CreatedAt    time.Time     `db:"created_at"`
}

type UserStatus string

const (
	UserActive    UserStatus = "active"
	UserSuspended UserStatus = "suspended"
	UserInactive  UserStatus = "inactive"
)

type User struct {
	ID        string     `db:"id"`
	AccountID string     `db:"account_id"`
	Status    UserStatus `db:"status"`
	Email     string     `db:"email"`
	Name      string     `db:"name"`
	UpdatedAt time.Time  `db:"updated_at"`
	CreatedAt time.Time  `db:"created_at"`
	LastLogin *time.Time `db:"last_login"`
}

type CreateAccountRequest struct {
	Name         string `db:"name"`
	ContactEmail string `db:"contact_email"`
}

type GetAccountRequest struct {
	ID string
}

type FetchAccountsRequest struct {
	ID string
}

type CreateUserRequest struct {
	AccountID string `db:"account_id"`
	Name      string `db:"name"`
	Email     string `db:"email"`
}

type GetUserRequest struct {
	ID string
}

type FetchUsersRequest struct {
	ID string
}

type Service interface {
	CreateAccount(context.Context, CreateAccountRequest) (Account, error)
	GetAccount(context.Context, GetAccountRequest) (Account, error)
	FetchAccounts(context.Context, FetchAccountsRequest) ([]Account, error)
	CreateUser(context.Context, CreateUserRequest) (User, error)
	GetUser(context.Context, GetUserRequest) (User, error)
	FetchUsers(context.Context, FetchUsersRequest) ([]User, error)
}

// NewService ...
func NewService(repo Repository) Service {
	return &customersService{
		logger: logutil.NewServerLogger(false, "customers"),
		repo:   repo,
	}
}

type customersService struct {
	logger log.Logger
	repo   Repository
}

func (svc *customersService) CreateAccount(ctx context.Context, req CreateAccountRequest) (account Account, err error) {
	account, err = svc.repo.InsertAccount(ctx, Account{ContactEmail: req.ContactEmail, Name: req.Name})
	if err != nil {
		svc.logger.Log("level", "error", "message", "error", err.Error(), "message", "failed to insert account")
	}

	return
}

func (svc *customersService) GetAccount(ctx context.Context, req GetAccountRequest) (account Account, err error) {
	account, err = svc.repo.GetAccountByID(ctx, req.ID)
	if err != nil {
		svc.logger.Log("level", "error", "message", "error", err.Error(), "message", "failed to retrieve account")
	}

	return
}

func (svc *customersService) FetchAccounts(ctx context.Context, req FetchAccountsRequest) (accounts []Account, err error) {
	accounts, err = svc.repo.SelectAccounts(ctx, map[string]interface{}{})
	if err != nil {
		svc.logger.Log("level", "error", "message", "error", err.Error(), "message", "failed to retrieve accounts")
	}
	return
}

func (svc *customersService) CreateUser(ctx context.Context, req CreateUserRequest) (user User, err error) {
	user, err = svc.repo.InsertUser(ctx, User{AccountID: req.AccountID, Email: req.Email, Name: req.Name})
	if err != nil {
		svc.logger.Log("level", "error", "message", "error", err.Error(), "message", "failed to insert user")
	}

	return
}

func (svc *customersService) GetUser(ctx context.Context, req GetUserRequest) (user User, err error) {
	user, err = svc.repo.GetUserByID(ctx, req.ID)
	if err != nil {
		svc.logger.Log("level", "error", "message", "error", err.Error(), "message", "failed to retrieve user")
	}

	return
}

func (svc *customersService) FetchUsers(ctx context.Context, req FetchUsersRequest) (users []User, err error) {
	users, err = svc.repo.SelectUsers(ctx, map[string]interface{}{})
	if err != nil {
		svc.logger.Log("level", "error", "message", "error", err.Error(), "message", "failed to retrieve users")
	}
	return
}
