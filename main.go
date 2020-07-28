package main

import (
	"github.com/limoxi/ghost"
	"github.com/limoxi/ghost/utils/cron"
	_ "picasso/api"
	_ "picasso/cron"
	_ "picasso/db"
	_ "picasso/middleware"
)

func main(){
	cron.StartCronTasks()
	ghost.RunWebServer()
}