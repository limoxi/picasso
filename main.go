package main

import (
	"flag"
	"github.com/limoxi/ghost"
	"github.com/limoxi/ghost/utils/cron"
	_ "picasso/api"
	_ "picasso/cron"
	_ "picasso/db"
	_ "picasso/middleware"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 && args[0] == "sync" {
		ghost.SyncDB()
		return
	}
	cron.StartCronTasks()
	ghost.RunWebServer()
}
