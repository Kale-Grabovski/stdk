package main

import (
	"database/sql"

	"github.com/Kale-Grabovski/stdk/src/controller/module"
	"github.com/Kale-Grabovski/stdk/src/controller/version"
	"github.com/Kale-Grabovski/stdk/src/repository"
	"github.com/Kale-Grabovski/stdk/src/service"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/sarulabs/di"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	viper.SetConfigFile(".bin/config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic("Error occurred while reading config file, run: cp .bin/config.dist.json .bin/config.json\n")
	}

	r := gin.Default()
	r.LoadHTMLGlob("src/templates/*")

	d := initCrmDI()

	r.GET("/", d.Get("action.modules_list").(*module.ModulesList).Handle)
	r.POST("/", d.Get("action.module_create").(*module.ModuleCreate).Handle)
	r.GET("/versions/:id", d.Get("action.versions_list").(*version.VersionsList).Handle)
	r.POST("/versions/:id", d.Get("action.version_create").(*version.VersionCreate).Handle)
	r.POST("/versions/:id/setActive", d.Get("action.version_set_active").(*version.SetActive).Handle)

	err = r.Run(":8080")
	if err != nil {
		panic(err)
	}
}

func initCrmDI() di.Container {
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
			Name:  "repo.module",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				db := ctx.Get("db").(*sql.DB)
				return repository.NewModuleRepository(db), nil
			},
		},
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
			Name:  "action.modules_list",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				repo := ctx.Get("repo.module").(*repository.ModuleRepository)
				logger := ctx.Get("logger").(*zap.Logger)
				return module.NewModulesList(logger, repo), nil
			},
		},
		{
			Name:  "action.module_create",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				repo := ctx.Get("repo.module").(*repository.ModuleRepository)
				logger := ctx.Get("logger").(*zap.Logger)
				return module.NewModuleCreate(logger, repo), nil
			},
		},
		{
			Name:  "action.versions_list",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				repo := ctx.Get("repo.version").(*repository.VersionRepository)
				logger := ctx.Get("logger").(*zap.Logger)
				return version.NewVersionsList(logger, repo), nil
			},
		},
		{
			Name:  "action.version_set_active",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				repo := ctx.Get("repo.version").(*repository.VersionRepository)
				logger := ctx.Get("logger").(*zap.Logger)
				return version.NewSetActive(logger, repo), nil
			},
		},
		{
			Name:  "action.version_create",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				repo := ctx.Get("repo.version").(*repository.VersionRepository)
				logger := ctx.Get("logger").(*zap.Logger)
				crypt := ctx.Get("service.crypt").(*service.CryptService)
				return version.NewVersionCreate(logger, repo, crypt), nil
			},
		},
		{
			Name:  "service.crypt",
			Scope: di.App,
			Build: func(ctx di.Container) (interface{}, error) {
				return service.NewCryptService(), nil
			},
		},
	}...)
	if err != nil {
		panic("Unable to build DI containers")
	}

	return builder.Build()
}
