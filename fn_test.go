package main

import (
	"github.com/limoxi/ghost"
	"gorm.io/gorm"
	"testing"
)

func TestFn(t *testing.T) {
	st := new(gorm.Statement)
	cns := st.BuildCondition(ghost.Map{
		"id__in":            []int{1, 2, 3},
		"created_at__range": []string{"2021-01-02", "2021-12-09"},
		"ss__not":           "sync",
		"age__gte":          21,
	})
	t.Log(cns)
}
