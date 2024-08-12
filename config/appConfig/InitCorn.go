package appconfig

import (
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/lenna-ai/azureOneSmile.git/app/kernel/console"
	"github.com/robfig/cron/v3"
)

func initCornJob() {
	jakartaTime, _ := time.LoadLocation("Asia/Jakarta")
	scheduler := cron.New(cron.WithLocation(jakartaTime))
	var wg sync.WaitGroup

	go scheduler.AddFunc("59 23 * * *", func() {
		wg.Add(1)
		intTimeDay,_ := strconv.Atoi(os.Getenv("TIME_STORAGE_DAY"))
		console.RemoveFileStorage("./storage/logs", intTimeDay)
	})

	wg.Wait()
	go scheduler.Start()

	scheduler.Stop()

}
