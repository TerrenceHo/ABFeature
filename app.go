package ABFeature

import (
	"github.com/TerrenceHo/ABFeature/controllers"
	"github.com/TerrenceHo/ABFeature/loggers"
	"github.com/TerrenceHo/ABFeature/services"
	"github.com/TerrenceHo/ABFeature/stores"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Start(viper *viper.Viper) {
	var err error
	// initiate database
	db := stores.NewConnection(
		viper.GetString("DATABASE.ENGINE"),
		viper.GetString("DATABASE.USER"),
		viper.GetString("DATABASE.PASSWORD"),
		viper.GetString("DATABASE.DBNAME"),
		viper.GetString("DATABASE.PORT"),
		viper.GetString("DATABASE.HOST"),
	)
	err = db.Ping()
	must(err)

	// create logger
	loggerUnsugared, err := zap.NewDevelopment()
	must(err)
	defer loggerUnsugared.Sync()
	logger := loggers.NewLogger(loggerUnsugared.Sugar())

	// initiate stores, and migrate tables
	projectStore := stores.NewProjectStore(db)
	experimentStore := stores.NewExperimentStore(db)
	groupStore := stores.NewGroupStore(db)

	stores.CreateTables(
		projectStore,
		experimentStore,
		groupStore,
	)

	// initiate services, connecting to stores
	projectService := services.NewProjectService(projectStore, logger)
	experimentService := services.NewExperimentService(experimentStore, logger)
	groupService := services.NewGroupService(groupStore, logger)

	// initiate http controllers, interfacing with services
	pagesController := controllers.NewPagesController(logger)
	projectController := controllers.NewProjectController(projectService, experimentService, logger)
	experimentController := controllers.NewExperimentController(experimentService, logger)
	groupController := controllers.NewGroupController(groupService, logger)

	// Configuration for a new Echo Server
	app := setupApp(viper)

	// Mount routes for new server, according to their groupings
	pagesController.MountRoutes(app.Group(""))
	projectController.MountRoutes(app.Group("/projects"))
	experimentController.MountRoutes(app.Group("/experiments"))
	groupController.MountRoutes(app.Group("/groups"))

	app.Logger.Fatal(app.Start(":" + viper.GetString("PORT")))
}

func setupApp(viper *viper.Viper) *echo.Echo {
	app := echo.New()
	app.HideBanner = true
	app.Debug = viper.GetBool("DEBUG")
	app.Pre(middleware.RequestID())
	app.Use(middleware.Gzip())
	app.Use(middleware.Secure())
	app.Use(middleware.RemoveTrailingSlash())
	app.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339_nano} ${method} {id":"${id}","remote_ip":"${remote_ip}",` +
			`"uri":"${uri}","status":${status},"latency":${latency},` +
			`"latency_human":"${latency_human}","bytes_in":${bytes_in},` +
			`"bytes_out":${bytes_out}}` + "\n",
	}))

	return app
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
