package app

import (
	"context"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	logger "github.com/sirupsen/logrus"

	pbUser "github.com/olesonya/highload-architect-course/homework.01/pkg/grpc/user/v1"
)

const (
	gracefulShutdownTimeout = time.Second * 10
)

type App struct {
	params     *appParams
	grpcServer *grpc.Server
}

func NewApp(ctx context.Context) (*App, error) {
	a := &App{}

	err := a.initDependencies(ctx)
	if err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() {
	go a.grpcServerStart()

	defer a.params.db.Disconnect()
	defer a.params.metrics.Stop()
	defer a.grpcServerStop()

	a.waitForStop()
}

func (a *App) grpcServerStart() {
	logger.Infof("GRPC server is running on %s", a.params.cfg.GetGRPCAddress())

	listener, err := net.Listen("tcp", a.params.cfg.GetGRPCAddress())
	if err != nil {
		logger.Fatalf("unable to create grpc listener: %v", err)
	}

	if err = a.grpcServer.Serve(listener); err != nil {
		logger.Fatalf("unable to start server: %v", err)
	}
}

func (a *App) grpcServerStop() {
	ch := make(chan struct{}, 1)

	go func() {
		a.grpcServer.GracefulStop()

		ch <- struct{}{}
	}()

	select {
	case <-ch:
		logger.Infoln("GRPC server graceful stop")
	case <-time.After(gracefulShutdownTimeout):
		a.grpcServer.Stop()

		logger.Infoln("GRPC server force stop")
	}
}

func (a *App) waitForStop() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	for sig := range sigs {
		logger.Infof("Caught %v", sig)

		if sig == syscall.SIGINT || sig == syscall.SIGTERM {
			return
		}
	}
}

func (a *App) initDependencies(ctx context.Context) error {
	inits := []func(context.Context) error{
		a.initServerConfig,
		a.initGRPCServer,
	}

	for _, f := range inits {
		err := f(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initServerConfig(_ context.Context) error {
	params, err := newServerConfig()
	if err != nil {
		return err
	}

	a.params = params

	return nil
}

func (a *App) initGRPCServer(_ context.Context) error {
	a.grpcServer = grpc.NewServer(grpc.Creds(insecure.NewCredentials()))

	reflection.Register(a.grpcServer)

	pbUser.RegisterUserServiceServer(a.grpcServer, a.params.userImplementation())

	return nil
}
