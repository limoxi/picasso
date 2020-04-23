package main

import (
	"github.com/limoxi/ghost"
	_ "picasso/api/command"
	_ "picasso/api/op"
	_ "picasso/api/user"
	_ "picasso/common/db"
)

func main(){
	ghost.RunWebServer()
}