package transport

import (
	"context"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	customerEndpoint "github.com/symptomatichq/customers/endpoint"
	"github.com/symptomatichq/customers/service"
	pb "github.com/symptomatichq/protos/customers"
)

type grpcServer struct {
	logger log.Logger

	createAccount grpctransport.Handler
	getAccount    grpctransport.Handler
	fetchAccounts grpctransport.Handler

	createUser grpctransport.Handler
	getUser    grpctransport.Handler
	fetchUsers grpctransport.Handler
}

// CreateAccount
func (s *grpcServer) CreateAccount(ctx context.Context, req *pb.CreateAccountRequest) (*pb.CreateAccountResponse, error) {
	_, resp, err := s.createAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.CreateAccountResponse), nil
}

// GetAccount
func (s *grpcServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	_, resp, err := s.getAccount.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GetAccountResponse), nil
}

// FetchAccounts
func (s *grpcServer) FetchAccounts(ctx context.Context, req *pb.FetchAccountsRequest) (*pb.FetchAccountsResponse, error) {
	_, resp, err := s.fetchAccounts.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.FetchAccountsResponse), nil
}

// CreateUser
func (s *grpcServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	_, resp, err := s.createUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.CreateUserResponse), nil
}

// GetUser
func (s *grpcServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	_, resp, err := s.getUser.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.GetUserResponse), nil
}

// FetchUsers
func (s *grpcServer) FetchUsers(ctx context.Context, req *pb.FetchUsersRequest) (*pb.FetchUsersResponse, error) {
	_, resp, err := s.fetchUsers.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}

	return resp.(*pb.FetchUsersResponse), nil
}

// NewGRPCServer create new grpc server
func NewGRPCServer(endpoints customerEndpoint.Endpoints, logger log.Logger) pb.CustomersServer {
	options := []grpctransport.ServerOption{
		grpctransport.ServerErrorLogger(logger),
	}

	return &grpcServer{
		logger: logger,
		createAccount: grpctransport.NewServer(
			endpoints.CreateAccountEndpoint,
			decodeGrpcCreateAccountRequest,
			encodeGrpcCreateAccountResponse,
			options...,
		),
		getAccount: grpctransport.NewServer(
			endpoints.GetAccountEndpoint,
			decodeGrpcGetAccountRequest,
			encodeGrpcAccountResponse,
			options...,
		),
		fetchAccounts: grpctransport.NewServer(
			endpoints.FetchAccountsEndpoint,
			decodeGrpcFetchAccountsRequest,
			encodeGrpcFetchAccountsResponse,
			options...,
		),
		createUser: grpctransport.NewServer(
			endpoints.CreateUserEndpoint,
			decodeGrpcCreateUserRequest,
			encodeGrpcCreateUserResponse,
			options...,
		),
		getUser: grpctransport.NewServer(
			endpoints.GetUserEndpoint,
			decodeGrpcGetUserRequest,
			encodeGrpcUserResponse,
			options...,
		),
		fetchUsers: grpctransport.NewServer(
			endpoints.FetchUsersEndpoint,
			decodeGrpcFetchUsersRequest,
			encodeGrpcFetchUsersResponse,
			options...,
		),
	}
}

// MakeGRPCGetAccountEndpoint creates GetAccount Endpoint for GRPC
func MakeGRPCGetAccountEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.GetAccountRequest)
		account, err := svc.GetAccount(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, errors.Wrap(err, "internal error").Error())
		}

		return account, nil
	}
}

// MakeGRPCCreateAccountEndpoint creates CreateAccount endpoint.Endpoint for GRPC
func MakeGRPCCreateAccountEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.CreateAccountRequest)
		account, err := svc.CreateAccount(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid request")
		}

		return account, nil
	}
}

// MakeGRPCFetchAccountsEndpoint creates FetchAccounts Endpoint
func MakeGRPCFetchAccountsEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.FetchAccountsRequest)
		accounts, err := svc.FetchAccounts(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid request")
		}

		return accounts, nil
	}
}

// MakeGRPCGetUserEndpoint creates GetUser Endpoint for GRPC
func MakeGRPCGetUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.GetUserRequest)
		user, err := svc.GetUser(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.Internal, errors.Wrap(err, "internal error").Error())
		}

		return user, nil
	}
}

// MakeGRPCCreateUserEndpoint creates CreateUser endpoint.Endpoint for GRPC
func MakeGRPCCreateUserEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.CreateUserRequest)
		user, err := svc.CreateUser(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid request")
		}

		return user, nil
	}
}

// MakeGRPCFetchUsersEndpoint creates FetchUsers Endpoint
func MakeGRPCFetchUsersEndpoint(svc service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(service.FetchUsersRequest)
		users, err := svc.FetchUsers(ctx, req)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "invalid request")
		}

		return users, nil
	}
}

// encodeAccount serializes the engine into a valid protobuf response
func encodeAccount(a service.Account) *pb.Account {
	return &pb.Account{
		ID:           a.ID,
		Name:         a.Name,
		ContactEmail: a.ContactEmail,
		Status:       encodeAccountStatus(a.Status),
		UpdatedAt:    a.UpdatedAt,
		CreatedAt:    a.CreatedAt,
	}
}

// decodeAccount serializes the engine into a valid protobuf response
func decodeAccount(a *pb.Account) service.Account {
	return service.Account{
		ID:           a.ID,
		Name:         a.Name,
		ContactEmail: a.ContactEmail,
		Status:       decodeAccountStatus(a.Status),
		UpdatedAt:    a.UpdatedAt,
		CreatedAt:    a.CreatedAt,
	}
}

// encodeUser serializes the engine into a valid protobuf response
func encodeUser(a service.User) *pb.User {
	return &pb.User{
		ID:        a.ID,
		Name:      a.Name,
		Email:     a.Email,
		Status:    encodeUserStatus(a.Status),
		LastLogin: a.LastLogin,
		UpdatedAt: a.UpdatedAt,
		CreatedAt: a.CreatedAt,
	}
}

// decodeUser serializes the engine into a valid protobuf response
func decodeUser(a *pb.User) service.User {
	return service.User{
		ID:        a.ID,
		Name:      a.Name,
		Email:     a.Email,
		Status:    decodeUserStatus(a.Status),
		UpdatedAt: a.UpdatedAt,
		CreatedAt: a.CreatedAt,
	}
}

// encodeGrpcAccountResponse encodes pb Account responses
func encodeGrpcAccountResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(service.Account)
	return encodeAccount(resp), nil
}

// decodeGrpcCreateAccountRequest decodes Account requests
func decodeGrpcCreateAccountRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateAccountRequest)

	return service.CreateAccountRequest{
		Name:         req.Name,
		ContactEmail: req.ContactEmail,
	}, nil
}

// decodeGrpcGetAccountRequest encodes FetchAccounts responses
func decodeGrpcGetAccountRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetAccountRequest)
	return req.ID, nil
}

// decodeGrpcFetchAccountsRequest encodes FetchAccounts responses
func decodeGrpcFetchAccountsRequest(ctx context.Context, r interface{}) (interface{}, error) {
	_ = r.(*pb.FetchAccountsRequest)
	return service.FetchAccountsRequest{}, nil
}

// encodeGrpcCreateAccountResponse encodes CreateAccountResponse responses
func encodeGrpcCreateAccountResponse(ctx context.Context, r interface{}) (interface{}, error) {
	account := r.(service.Account)
	encoded := encodeAccount(account)
	return &pb.CreateAccountResponse{
		Account: *encoded,
	}, nil
}

// encodeGrpcFetchAccountsResponse encodes FetchAccounts responses
func encodeGrpcFetchAccountsResponse(ctx context.Context, r interface{}) (interface{}, error) {
	accounts := []pb.Account{}
	for _, account := range r.([]service.Account) {
		encoded := encodeAccount(account)
		accounts = append(accounts, *encoded)
	}
	return &pb.FetchAccountsResponse{
		Accounts: accounts,
	}, nil
}

func encodeAccountStatus(status service.AccountStatus) pb.Account_Status {
	switch status {
	case service.AccountActive:
		return pb.Account_ACTIVE
	case service.AccountSuspended:
		return pb.Account_SUSPENDED
	}

	return pb.Account_INACTIVE
}

func decodeAccountStatus(status pb.Account_Status) service.AccountStatus {
	switch status {
	case pb.Account_ACTIVE:
		return service.AccountActive
	case pb.Account_SUSPENDED:
		return service.AccountSuspended
	}

	return service.AccountInactive
}

// encodeGrpcUserResponse encodes pb User responses
func encodeGrpcUserResponse(_ context.Context, r interface{}) (interface{}, error) {
	resp := r.(service.User)
	return encodeUser(resp), nil
}

// decodeGrpcCreateUserRequest decodes User requests
func decodeGrpcCreateUserRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.CreateUserRequest)

	return service.CreateUserRequest{
		AccountID: req.AccountID,
		Email:     req.Email,
		Name:      req.Name,
	}, nil
}

// decodeGrpcGetUserRequest encodes FetchUsers responses
func decodeGrpcGetUserRequest(ctx context.Context, r interface{}) (interface{}, error) {
	req := r.(*pb.GetUserRequest)
	return req.ID, nil
}

// decodeGrpcFetchUsersRequest encodes FetchUsers responses
func decodeGrpcFetchUsersRequest(ctx context.Context, r interface{}) (interface{}, error) {
	_ = r.(*pb.FetchUsersRequest)
	return service.FetchUsersRequest{}, nil
}

// encodeGrpcCreateUserResponse encodes CreateUserResponse responses
func encodeGrpcCreateUserResponse(ctx context.Context, r interface{}) (interface{}, error) {
	account := r.(service.User)
	encoded := encodeUser(account)
	return &pb.CreateUserResponse{
		User: *encoded,
	}, nil
}

// encodeGrpcFetchUsersResponse encodes FetchUsers responses
func encodeGrpcFetchUsersResponse(ctx context.Context, r interface{}) (interface{}, error) {
	accounts := []pb.User{}
	for _, account := range r.([]service.User) {
		encoded := encodeUser(account)
		accounts = append(accounts, *encoded)
	}
	return &pb.FetchUsersResponse{
		Users: accounts,
	}, nil
}

func encodeUserStatus(status service.UserStatus) pb.User_Status {
	switch status {
	case service.UserActive:
		return pb.User_ACTIVE
	case service.UserSuspended:
		return pb.User_SUSPENDED
	}

	return pb.User_INACTIVE
}

func decodeUserStatus(status pb.User_Status) service.UserStatus {
	switch status {
	case pb.User_ACTIVE:
		return service.UserActive
	case pb.User_SUSPENDED:
		return service.UserSuspended
	}

	return service.UserInactive
}
