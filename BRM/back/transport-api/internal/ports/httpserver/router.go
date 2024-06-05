package httpserver

import (
	"github.com/gin-gonic/gin"
	_ "transport-api/docs"
	"transport-api/internal/app"
	"transport-api/internal/ports/httpserver/ads"
	"transport-api/internal/ports/httpserver/core/companies"
	"transport-api/internal/ports/httpserver/core/contacts"
	"transport-api/internal/ports/httpserver/core/employees"
	"transport-api/internal/ports/httpserver/leads"
	"transport-api/internal/ports/httpserver/middleware"
	"transport-api/internal/ports/httpserver/notifications"
	"transport-api/pkg/logger"
	"transport-api/pkg/tokenizer"
)

func appRouter(r *gin.RouterGroup, a app.App, tkn tokenizer.Tokenizer, logs logger.Logger) {
	r.Use(middleware.Panic(logs))
	r.Use(middleware.Log(logs))
	r.Use(middleware.Auth(tkn))

	r.GET("/companies/:id", companies.GetCompany(a))
	r.GET("/companies/:id/mainpage", companies.GetCompanyMainPage(a))
	r.PUT("/companies/:id", companies.UpdateCompany(a))
	r.GET("/companies/industries", companies.GetIndustriesMap(a))

	r.POST("/employees", employees.AddEmployee(a))
	r.GET("/employees/:id", employees.GetEmployee(a))
	r.GET("/employees", employees.GetEmployeesList(a))
	r.PUT("/employees/:id", employees.UpdateEmployee(a))
	r.DELETE("/employees/:id", employees.DeleteEmployee(a))

	r.POST("/contacts", contacts.AddContact(a))
	r.GET("/contacts/:id", contacts.GetContact(a))
	r.GET("/contacts", contacts.GetContactsList(a))
	r.PUT("/contacts/:id", contacts.UpdateContact(a))
	r.DELETE("/contacts/:id", contacts.DeleteContact(a))

	r.POST("/market", ads.AddAd(a))
	r.GET("/market/:id", ads.GetAd(a))
	r.GET("/market", ads.GetAdsList(a))
	r.PUT("/market/:id", ads.UpdateAd(a))
	r.DELETE("/market/:id", ads.DeleteAd(a))
	r.POST("/market/:id/response", ads.AddResponse(a))
	r.GET("/responses", ads.GetResponsesList(a))
	r.GET("/market/industries", ads.GetIndustriesMap(a))

	r.GET("/leads/:id", leads.GetLead(a))
	r.GET("/leads", leads.GetLeadsList(a))
	r.PUT("/leads/:id", leads.UpdateLead(a))
	r.GET("/leads/statuses", leads.GetStatuses(a))

	r.GET("/notifications", notifications.GetNotifications(a))
	r.GET("/notifications/:id", notifications.GetNotification(a))
	r.POST("/notifications/:id", notifications.SubmitClosedLead(a))
}
