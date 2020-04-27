package main

import (
	"github.com/limoxi/ghost"
	_ "picasso/api/command"
	_ "picasso/api/op"
	_ "picasso/api/space"
	_ "picasso/api/user"
	_ "picasso/common/db"
	_ "picasso/middleware"
)

func main(){
	ghost.RunWebServer()
}