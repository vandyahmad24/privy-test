package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"vandyahmad24/privy/app/db"
	"vandyahmad24/privy/app/router"
	"vandyahmad24/privy/app/tracing"

	"github.com/gofiber/fiber/v2"
	opentracing "github.com/opentracing/opentracing-go"
)

func main() {
	fmt.Println("Privy Test By Vandy Ahmad")
	tracer, closer := tracing.Init("Cake Service")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	app := fiber.New(fiber.Config{
		BodyLimit: 8 * 1024 * 1024, // this is the default limit of 4MB
	})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("welcome to Cake Service")
	})
	db := db.InitDb()
	router.CakeRouter(app, db)

	go func() {
		app.Listen(":8080")
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("Server Mati : %v\n", signal.String())

}
