package main

import (
	"fmt"
	"sync"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/lenna-ai/azureOneSmile.git/config"
	appconfig "github.com/lenna-ai/azureOneSmile.git/config/appConfig"
	"github.com/lenna-ai/azureOneSmile.git/helpers"
)

func main()  {
	defer helpers.RecoverPanicContext(&fiber.Ctx{})
	appconfig.InitApplication()
	app := fiber.New()
	app.Use(cors.New())
	config.Logger(app)
	// routes.Router(app)

	if err := app.Listen(":3000"); err != nil {
		panic(err.Error())
	}

	// var swg sync.WaitGroup
	// channel := make(chan int)
	// for i := 0; i < 10; i++ {
	// 	swg.Add(1)
	// 	go increaseNumber(channel,i,&swg)
	// }

	// go func ()  {
	// 	swg.Wait()
	// }()
	// close(channel)
	// readNumber(channel)

}

func increaseNumber(channel chan<- int,number int,swg *sync.WaitGroup)  {
	defer swg.Done()
	channel <- number
}

func readNumber(channel <-chan int)  {
	for v := range channel {
		fmt.Println(v)
	}
}