package main

import (
	"context"
	"go-srv/internal/ports"
	"go-srv/internal/repo"
	"go-srv/internal/services"
	"log"

	"github.com/spf13/viper"
)

func InitConfig() error {
	viper.SetConfigFile("./config.yaml")
	return viper.ReadInConfig()
}

func main() {
	if err := InitConfig(); err != nil {
		log.Fatal(err.Error())
	}
	ctx := context.Background()

	repo, err := repo.NewRepo(ctx,
		viper.GetString("db.host"),
		viper.GetInt("db.port"),
		viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.dbname"),
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	adminService := services.NewAdminService(repo)
	adminPage := ports.NewAdminPage(adminService)

	log.Println("🎉server started🎉")
	adminPage.Run(viper.GetString("admin.addr"))
}
