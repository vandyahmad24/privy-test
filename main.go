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
	"vandyahmad24/privy/app/util"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	opentracing "github.com/opentracing/opentracing-go"
)

func main() {
	fmt.Println("Privy Test By Vandy Ahmad")
	port := os.Getenv("PORT_GOLANG")
	if port == "" {
		log.Fatal("Port env is requeird")
	}
	tracer, closer := tracing.Init("Cake Service")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	app := fiber.New(fiber.Config{
		BodyLimit: 8 * 1024 * 1024, // this is the default limit of 4MB
	})
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON("welcome to Cake Service")
	})
	fmt.Println(util.GetEnvVariable("MYSQL_DBNAME"))
	app.Get("/panic", func(c *fiber.Ctx) error {
		panic("handle panic")
	})

	db := db.InitDb()
	router.CakeRouter(app, db)

	go func() {
		app.Listen(fmt.Sprintf(":%s", port))
	}()
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	signal := <-c
	log.Fatalf("Server Mati : %v\n", signal.String())

}
