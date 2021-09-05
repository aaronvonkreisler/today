package application

import (
	"github.com/todo-app/internal"
	"github.com/todo-app/internal/repositories"
	"github.com/todo-app/internal/services"
	"github.com/todo-app/pkg/config"
	"github.com/todo-app/pkg/logger"
)

type App struct {
	DataStore       *internal.DataStore
	Confg           *config.Confg
	UserRepository  repositories.UserRepositoryInterface
	IdentityService services.IdentityServiceInterface
}

func BootstrapApp() (*App, error) {
	cfg := config.Get()

	db, err := internal.GetDataStore(cfg.GetDBConnStr())
	if err != nil {
		return nil, err
	}

	logger.Info.Print("successfully connected to database")
	return &App{
		DataStore:       db, // TODO: take out -- db should only be accessed in repositories
		Confg:           cfg,
		UserRepository:  repositories.NewUserRepository(db.Client),
		IdentityService: services.NewIdentityService(db.Client),
	}, nil
}
