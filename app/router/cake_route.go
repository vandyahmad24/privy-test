package router

import (
	"database/sql"
	"vandyahmad24/privy/app/adapter/api"
	"vandyahmad24/privy/app/adapter/database"
	"vandyahmad24/privy/app/usecase/cakeusecase"

	"github.com/gofiber/fiber/v2"
)

func CakeRouter(router *fiber.App, db *sql.DB) {
	cakeRepository := database.NewCake(db)
	cakeService := cakeusecase.NewCakeUsecase(cakeRepository)
	roleHandler := api.NewCakeServiceService(cakeService)
	router.Post("/cakes", roleHandler.CreateCakeService)
	router.Get("/cakes", roleHandler.GetAllCakeService)
	router.Get("/cakes/:id", roleHandler.GetCakeService)
	router.Delete("/cakes/:id", roleHandler.DeleteCakeService)
	router.Patch("/cakes/:id", roleHandler.UpdateCakeService)

}
