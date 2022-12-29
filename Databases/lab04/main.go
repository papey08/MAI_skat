package main

import (
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
	"lab04/app/controller"
	"lab04/app/server"
	"log"
	"net/http"
)

func main() {
	err := server.OpenDB("user=user password=password dbname=dbname sslmode=disable")
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer func() {
		err = server.CloseDB()
		if err != nil {
			log.Fatal(err.Error())
		}
	}()
	r := httprouter.New()
	r.ServeFiles("/public/*filepath", http.Dir("public"))
	r.GET("/", controller.StartPage)
	r.GET("/driver", controller.GetDrivers)
	r.GET("/vehicle", controller.GetVehicles)
	r.GET("/voyage", controller.GetVoyages)
	r.POST("/driver/insert", controller.InsertDriver)
	r.POST("/driver/delete", controller.DeleteDriver)
	r.POST("/vehicle/insert", controller.InsertVehicle)
	r.POST("/vehicle/delete", controller.DeleteVehicle)
	r.POST("/voyage/insert", controller.InsertVoyage)
	r.POST("/voyage/delete", controller.DeleteVoyage)
	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	return
}
