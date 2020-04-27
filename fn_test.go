package main

import (
	"github.com/limoxi/ghost"
	m_space "picasso/common/db/space"
	"testing"
)

func TestFn (t *testing.T){
	var dbModels []*m_space.Space
	db := ghost.GetDB().Model(&m_space.Space{}).Where(ghost.Map{})
	db = ghost.NewPaginator(1, 3).Paginate(db)
	result := db.Find(&dbModels)
	if err := result.Error; err != nil{
		t.Log(err)
	}
	for _, dbModel := range dbModels{
		t.Log(dbModel.Id)
	}
}