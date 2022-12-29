package main

import (
	"cp/app/controller"
	"cp/app/server"
	"fmt"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

// Routes matches methods with endpoints
func Routes(r *httprouter.Router) {
	r.ServeFiles("/public/*filepath", http.Dir("public"))

	r.GET("/", controller.StartPage)

	r.GET("/client", controller.SelectClients)
	r.POST("/client/insert", controller.InsertClient)
	r.POST("/client/update/subscription", controller.UpdateClientSubscription)
	r.POST("/client/update/height&weight", controller.UpdateHeightAndWeight)
	r.POST("/client/delete", controller.DeleteClient)

	r.GET("/trainer", controller.SelectTrainers)
	r.POST("/trainer/insert", controller.InsertTrainer)
	r.POST("/trainer/delete", controller.DeleteTrainer)

	r.GET("/group", controller.SelectGroups)
	r.POST("/group/insert", controller.InsertGroup)
	r.POST("/group/select", controller.SelectGroup)
	r.POST("/group/client/insert", controller.InsertClientIntoGroup)
	r.POST("/group/client/delete", controller.DeleteClientFromGroup)
	r.POST("/group/delete", controller.DeleteGroup)

	r.GET("/timetable", controller.SelectTimetable)
	r.POST("/timetable/insert", controller.InsertTimetable)
	r.POST("/timetable/select/by_group", controller.SelectTimetableByGroup)
	r.POST("/timetable/select/by_program", controller.SelectTimetableByProgram)
	r.POST("/timetable/select/by_trainer", controller.SelectTimetableByTrainer)
	r.POST("/timetable/delete", controller.DeleteTimetable)
}

// InitConfig initialises configuration file
func InitConfig() error {
	viper.SetConfigFile("config.yml")
	return viper.ReadInConfig()
}

func main() {
	if err := InitConfig(); err != nil {
		log.Fatal(err)
	}
	configStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		viper.GetString("db.host"), viper.GetString("db.port"),
		viper.GetString("db.user"), viper.GetString("db.dbname"),
		viper.GetString("db.password"), viper.GetString("db.sslmode"))

	if err := server.OpenDB(configStr); err != nil {
		log.Fatal(err)
		return
	}
	r := httprouter.New()
	Routes(r)
	if err := http.ListenAndServe(":"+viper.GetString("port"), r); err != nil {
		log.Fatal(err)
		return
	}
	if err := server.CloseDB(); err != nil {
		log.Fatal(err)
		return
	}
}
