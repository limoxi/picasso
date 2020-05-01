package main

import (
	"github.com/limoxi/ghost"
	_ "picasso/api"
	_ "picasso/db"
	_ "picasso/middleware"
)

func main(){
	ghost.RunWebServer()
}