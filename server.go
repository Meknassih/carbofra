package main

import (
	"fmt"
	"os"

	"github.com/Meknassih/carbofra/routers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main() {
	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
		ViewsLayout: "layouts/main",
	})
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	routers.Root(app)

	app.Listen(":" + port)
	fmt.Println("Server started on port " + port)
}
