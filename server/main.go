package main

import (
	"fmt"

	"server/actions"
	"server/controller"
	"server/models"
	v "server/variables"

	"github.com/labstack/gommon/log"
)

func init() {
	fmt.Println(v.Logo + v.TCR + v.Version + v.TCC)
	fmt.Println(v.Subtitle)
	fmt.Println(v.TCB + v.Repo + v.TCC)
}

func main() {

	initServerStatusCollector()

	app := actions.App()
	app.HideBanner = true
	app.Logger.SetLevel(log.INFO)
	app.Logger.SetHeader("${time_rfc3339} | ${level} | ${short_file}:${line}")
	app.Logger.Fatal(app.Start(v.ADDR + ":" + v.PORT))
}

func initServerStatusCollector() {
	go controller.CollectCpuStatus(models.Db)
<<<<<<< HEAD
	go controller.CollectMemStatus(models.Db)
=======
>>>>>>> 97e9e548f2ce395e9ae5a36f2f4c41359c1607b7
}
