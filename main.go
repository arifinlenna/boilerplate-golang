package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/lenna-ai/azureOneSmile.git/helpers"
)

func main()  {
	defer helpers.RecoverPanicContext(&fiber.Ctx{})
	fmt.Println("main")
}