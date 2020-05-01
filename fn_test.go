package main

import (
	"github.com/limoxi/ghost"
	db_space "picasso/db/space"
	"testing"
)

func TestFn (t *testing.T){
	var dbModels []*db_space.Space
	db := ghost.GetDB().Model(&db_space.Space{}).Where(ghost.Map{
		"id__gt": 0,
	})
	paginator := ghost.NewPaginator(2, 3)
	db = paginator.Paginate(db)
	result := db.Find(&dbModels)
	if err := result.Error; err != nil{
		t.Log(err)
	}
	t.Log(paginator.ToResultMap())
}