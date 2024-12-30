package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fajarherdian22/topo-api/controller"
	"github.com/fajarherdian22/topo-api/db"
	"github.com/fajarherdian22/topo-api/repository"
	"github.com/fajarherdian22/topo-api/service"
	"github.com/fajarherdian22/topo-api/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config :", err)
	}

	mysqlDB, err := db.ConDB(config.DBSource, "mysql")
	if err != nil {
		log.Fatalf("failed to connect to database:", err)
	}
	repo := repository.New(mysqlDB)

	postgisDB, err := db.ConDB(config.DBPostgis, "postgis")
	if err != nil {
		log.Fatalf("failed to connect to database:", err)
	}
	repoSpatial := repository.NewKabKotaRepository(postgisDB)

	validate := validator.New()
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("LevelName", util.LevelNameValidator)
	}

	ranService := service.NewRanService(repo)
	ranController := controller.NewRanController(ranService, validate)

	spatialService := service.NewKabKotaService(repoSpatial)
	spatialController := controller.NewSpatialController(spatialService, validate)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"POST", "GET"},
		AllowHeaders:    []string{"Content-Type", "Origin"},
		ExposeHeaders:   []string{"Content-Length"},
		MaxAge:          12 * time.Hour,
	}))

	r := router.Group("/api/")

	r.POST("/list/level", ranController.ListLevel)
	r.POST("/data/level", ranController.GetAllData)
	r.POST("/data/level/name", ranController.GetByLevel)
	r.POST("/data/level/reference", ranController.GetByReference)
	r.GET("/spatial/all", spatialController.GetAllSpatial)
	r.POST("/spatial/level/name", spatialController.GetSpatialByFilter)

	err = router.Run(config.HttpServerAddress)
	if err != nil {
		log.Fatalf("failed to start server:", err)
	}
	fmt.Printf("Runing server")
}
