package main

import (
	"database/sql"
	"os"

	"github.com/Kale-Grabovski/stdk/src/controller/api"
	"github.com/Kale-Grabovski/stdk/src/repository"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	if len(os.Args) < 2 {
		panic("You must pass correct config path as a parameter")
		return
	}

	viper.SetConfigFile(os.Args[1])
	err := viper.ReadInConfig()
	if err != nil {
		panic("Error occurred while reading config file, run: cp .bin/config.dist.json .bin/config.json\n")
	}

	r := gin.Default()
	d := initApiDI()

	r.POST("/api/v1/modules", d.Get("api.action.installed_modules").(*api.InstalledModules).Handle)
	r.GET("/api/v1/modules/:id", d.Get("api.action.get_module").(*api.GetModule).Handle)

	err = r.Run(":8081")
	if err != nil {
		panic(err)
	}
}

func initApiDI() di.Container {
	builder, _ := di.NewBuilder()

	err := builder.Add([]di.Def{
		{
			Name:  "db",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				d := viper.GetStringMapString("db")
				db, err := sql.Open(d["driver"], d["dsn"])
				if err != nil {
					panic("Could not connect to database: " + err.Error())
				}

				return db, err
			},
			Close: func(obj interface{}) error {
				return obj.(*sql.DB).Close()
			},
		},
		{
			Name:  "logger",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				return zap.NewProduction()
			},
			Close: func(obj interface{}) error {
				return obj.(*zap.Logger).Sync()
			},
		},

		// Repos
		{
			Name:  "repo.version",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				db := ctx.Get("db").(*sql.DB)
				return repository.NewVersionRepository(db), nil
			},
		},

		// Actions
		{
			Name:  "api.action.installed_modules",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				repo := ctx.Get("repo.version").(*repository.VersionRepository)
				logger := ctx.Get("logger").(*zap.Logger)
				return api.NewInstalledModules(logger, repo), nil
			},
		},
		{
			Name:  "api.action.get_module",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				repo := ctx.Get("repo.version").(*repository.VersionRepository)
				logger := ctx.Get("logger").(*zap.Logger)
				return api.NewGetModule(logger, repo), nil
			},
		},
	}...)
	if err != nil {
		panic("Unable to build DI containers")
	}

	return builder.Build()
}
