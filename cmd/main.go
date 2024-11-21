package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/configs"
	membershipsHdlr "github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/handlers/memberships"
	postsHdlr "github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/handlers/posts"
	memberhipRepo "github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/repository/memberships"
	postsRepo "github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/repository/posts"
	memberhsipSvc "github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/service/memberships"
	postsSvc "github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/internal/service/posts"

	"github.com/joshuatheokurniawansiregar/FastCampus-Proyek2-Situs-Forum/pkg/internalsql"
)

func main() {
	r := gin.Default()
	var (
		cfg *configs.Config
	)

	var err error = configs.Init(
		configs.WithConfigFolder([]string{
			"./internal/configs/",
		}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	cfg = configs.GetConfig()

	if err != nil {
		log.Fatal("Gagal Inisiasi Config", err)
	}

	database, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi DB", err)
	}
	defer database.Close()
	

	membershipsRepository := memberhipRepo.NewRepository(database)

	membershipsService := memberhsipSvc.NewService(membershipsRepository, cfg)

	membershipsHandler := membershipsHdlr.NewHandler(r, membershipsService)
	
	membershipsHandler.RegisterRoute()

	postsRepository := postsRepo.NewRepository(database)
	postsServIce:= postsSvc.NewService(postsRepository, cfg)
	postsHanlder:= postsHdlr.NewHandler(r, postsServIce)
	postsHanlder.RegisterRoute()
	r.Run(cfg.Service.Port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
