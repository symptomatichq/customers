package endpoint

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"github.com/pkg/errors"

	"github.com/symptomatichq/customers/service"
	"github.com/symptomatichq/kit/middleware"
)

// Endpoints defines the rpc interface for interacting wwith the accounts service
type Endpoints struct {
	CreateAccountEndpoint endpoint.Endpoint
	GetAccountEndpoint    endpoint.Endpoint
	FetchAccountsEndpoint endpoint.Endpoint
	CreateUserEndpoint    endpoint.Endpoint
	GetUserEndpoint       endpoint.Endpoint
	FetchUsersEndpoint    endpoint.Endpoint
}

// CreateAccount ...
func (e Endpoints) CreateAccount(ctx context.Context, newAccount service.Account) (account service.Account, err error) {
	resp, err := e.CreateAccountEndpoint(ctx, newAccount)
	if err != nil {
		return
	}

	account = resp.(service.Account)

	return
}

// GetAccount ...
func (e Endpoints) GetAccount(ctx context.Context, id *string) (accounts []service.Account, err error) {
	resp, err := e.GetAccountEndpoint(ctx, id)
	if err != nil {
		return
	}

	accounts = resp.([]service.Account)

	return
}

// FetchAccounts ...
func (e Endpoints) FetchAccounts(ctx context.Context, req *service.FetchAccountsRequest) (accounts []service.Account, err error) {
	resp, err := e.FetchAccountsEndpoint(ctx, req)
	if err != nil {
		return
	}

	accounts = resp.([]service.Account)

	return
}

// MakeCreateAccountEndpoint creates CreateAccount Endpoint
func MakeCreateAccountEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.CreateAccountRequest)
		account, err := svc.CreateAccount(ctx, req)
		if err != nil {
			return nil, err
		}

		return account, nil
	}

}

// MakeGetAccountEndpoint creates GetAccount Endpoint
func MakeGetAccountEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.GetAccountRequest)
		account, err := svc.GetAccount(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "internal error")
		}

		return account, nil
	}

}

// MakeFetchAccountsEndpoint creates FetchAccounts Endpoint
func MakeFetchAccountsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.FetchAccountsRequest)
		accounts, err := svc.FetchAccounts(ctx, req)
		if err != nil {
			return nil, err
		}

		return accounts, nil
	}

}

// MakeCreateUserEndpoint creates CreateUser Endpoint
func MakeCreateUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.CreateUserRequest)
		user, err := svc.CreateUser(ctx, req)
		if err != nil {
			return nil, err
		}

		return user, nil
	}

}

// MakeGetUserEndpoint creates GetUser Endpoint
func MakeGetUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.GetUserRequest)
		user, err := svc.GetUser(ctx, req)
		if err != nil {
			return nil, errors.Wrap(err, "internal error")
		}

		return user, nil
	}

}

// MakeFetchUsersEndpoint creates FetchUsers Endpoint
func MakeFetchUsersEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.FetchUsersRequest)
		users, err := svc.FetchUsers(ctx, req)
		if err != nil {
			return nil, err
		}

		return users, nil
	}

}

// New returns a set of endpoints that wrap the service provider
func New(svc service.Service, logger log.Logger) Endpoints {
	createAccountEndpoint := middleware.LoggingMiddleware(
		log.With(logger, "method", "CreateAccount"),
	)(MakeCreateAccountEndpoint(svc))

	getAccountEndpoint := middleware.LoggingMiddleware(
		log.With(logger, "method", "GetAccount"),
	)(MakeGetAccountEndpoint(svc))

	fetchAccountsEndpoint := middleware.LoggingMiddleware(
		log.With(logger, "method", "FetchAccounts"),
	)(MakeFetchAccountsEndpoint(svc))

	createUserEndpoint := middleware.LoggingMiddleware(
		log.With(logger, "method", "CreateUser"),
	)(MakeCreateUserEndpoint(svc))

	getUserEndpoint := middleware.LoggingMiddleware(
		log.With(logger, "method", "GetUser"),
	)(MakeGetUserEndpoint(svc))

	fetchUsersEndpoint := middleware.LoggingMiddleware(
		log.With(logger, "method", "FetchUsers"),
	)(MakeFetchUsersEndpoint(svc))

	return Endpoints{
		CreateAccountEndpoint: createAccountEndpoint,
		GetAccountEndpoint:    getAccountEndpoint,
		FetchAccountsEndpoint: fetchAccountsEndpoint,
		CreateUserEndpoint:    createUserEndpoint,
		GetUserEndpoint:       getUserEndpoint,
		FetchUsersEndpoint:    fetchUsersEndpoint,
	}
}
