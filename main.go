package main

import (
	"flag"
	"fmt"
	"net"
	"os"

	"google.golang.org/grpc"

	"github.com/symptomatichq/customers/endpoint"
	"github.com/symptomatichq/customers/service"
	_ "github.com/symptomatichq/customers/telemetry" // load telemetry
	"github.com/symptomatichq/customers/transport"
	"github.com/symptomatichq/kit/env"
	"github.com/symptomatichq/kit/graceful"
	"github.com/symptomatichq/kit/health"
	"github.com/symptomatichq/kit/logutil"
	"github.com/symptomatichq/kit/pgutil"
	pb "github.com/symptomatichq/protos/customers"
)

var (
	debug *bool
	port  *int
)

func main() {
	port = flag.Int("port", env.Int("PORT", 8080), "GRPC server port")
	debug = flag.Bool("debug", env.Bool("DEBUG", false), "run the server in debug mode")

	logger := logutil.NewServerLogger(*debug, "customers")

	dbCfg := pgutil.ConfigFromEnv()
	// err := dbutil.Migrate(dbCfg, "./migrations")
	// if err != nil {
	// 	logger.Log("level", "error", "message", "unable to execute database migrations", "error", err.Error())
	// 	os.Exit(1)
	// } else {
	// 	logger.Log("level", "info", "message", "database migated to latest revision")
	// }

	repo := service.NewRepository(dbCfg)
	svc := service.NewService(repo)

	endpoints := endpoint.Endpoints{
		CreateAccountEndpoint: transport.MakeGRPCCreateAccountEndpoint(svc),
		GetAccountEndpoint:    transport.MakeGRPCGetAccountEndpoint(svc),
		FetchAccountsEndpoint: transport.MakeGRPCFetchAccountsEndpoint(svc),
		CreateUserEndpoint:    transport.MakeGRPCCreateUserEndpoint(svc),
		GetUserEndpoint:       transport.MakeGRPCGetUserEndpoint(svc),
		FetchUsersEndpoint:    transport.MakeGRPCFetchUsersEndpoint(svc),
	}

	addr, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logger.Log(
			"level", "error",
			"message", "failed to listen on network interface",
			"port", port,
			"error", err.Error(),
		)
		os.Exit(1)
	}

	handler := transport.NewGRPCServer(endpoints, logger)
	gRPCServer := grpc.NewServer()

	probe := health.NewServer(9090, logger, map[string]health.Checker{"noop": health.Nop()})
	probe.Start()

	graceful.Handle(func(signal os.Signal) {
		logger.Log("message", "shutting down server", "signal", signal.String())
		probe.Stop()
		gRPCServer.GracefulStop()
	})

	pb.RegisterCustomersServer(gRPCServer, handler)

	logger.Log("message", "server started")
	gRPCServer.Serve(addr)
}
