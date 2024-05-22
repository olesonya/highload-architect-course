package app

import (
	"os"

	logger "github.com/sirupsen/logrus"

	"github.com/olesonya/highload-architect-course/homework.01/internal/api/user"
	"github.com/olesonya/highload-architect-course/homework.01/internal/config"
	cfg "github.com/olesonya/highload-architect-course/homework.01/internal/config"
	database "github.com/olesonya/highload-architect-course/homework.01/internal/database"
	metrics "github.com/olesonya/highload-architect-course/homework.01/internal/metrics"
	repository "github.com/olesonya/highload-architect-course/homework.01/internal/repository"
	userRepo "github.com/olesonya/highload-architect-course/homework.01/internal/repository/user"
	service "github.com/olesonya/highload-architect-course/homework.01/internal/service"
	userService "github.com/olesonya/highload-architect-course/homework.01/internal/service/user"
)

type appParams struct {
	cfg *cfg.Config

	db *database.Database

	metrics *metrics.Server

	userRepo repository.UserRepository

	userSvc service.UserService

	userImpln *user.Instance
}

func loggerConfig(cfg *cfg.Config) {
	logger.SetReportCaller(true)
	logger.SetOutput(os.Stdout)
	logger.SetLevel(cfg.GetLogLevel())
	logger.SetFormatter(
		&logger.TextFormatter{
			FullTimestamp: true,
		},
	)
}

func newServerConfig() (*appParams, error) {
	cfg, err := config.New()
	if err != nil {
		return nil, err
	}

	loggerConfig(cfg)

	db := database.NewDatabase(cfg)
	db.Connect()
	db.MigrateDB()

	metrics := metrics.NewServer(cfg)
	metrics.Start()

	return &appParams{
		cfg:     cfg,
		db:      db,
		metrics: metrics,
	}, nil
}

func (a *appParams) userRepository() repository.UserRepository {
	if a.userRepo == nil {
		a.userRepo = userRepo.NewRepository(a.db)
	}

	return a.userRepo
}

func (a *appParams) userService() service.UserService {
	if a.userSvc == nil {
		a.userSvc = userService.NewService(
			a.userRepository(),
		)
	}

	return a.userSvc
}

func (a *appParams) userImplementation() *user.Instance {
	if a.userImpln == nil {
		a.userImpln = user.NewImplementation(a.userService())
	}

	return a.userImpln
}
