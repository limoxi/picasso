package main

import (
	"github.com/limoxi/ghost"
	"testing"
)

func TestDb(t *testing.T){
	ghost.SyncDB()
}