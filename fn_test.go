package main

import (
	"context"
	"github.com/limoxi/ghost"
	bs_media "picasso/business/app/file"
	db_space "picasso/db/user"
	"testing"
)

func TestFn(t *testing.T) {
	var dbModels []*db_space.Space
	db := ghost.GetDB().Model(&db_space.Space{}).Where(ghost.Map{
		"id__gt": 0,
	})
	paginator := ghost.NewPaginator(2, 3)
	db = paginator.Paginate(db)
	result := db.Find(&dbModels)
	if err := result.Error; err != nil {
		t.Log(err)
	}
	t.Log(paginator.ToMap())
}

type A struct {
	name string
}

func (this *A) Set(s string) {
	this.name = s
}

func (a A) get() string {
	return a.name
}

type B struct {
	*A
}

func TestFn2(t *testing.T) {

	t.Log(16 & 16)
}

func TestFn1(t *testing.T) {
	bs_media.NewMediaMetadataProcessor(context.Background()).ProcessImage(nil)
}
