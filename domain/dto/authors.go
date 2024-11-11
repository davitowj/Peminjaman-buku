package dto

import (
	"base-gin/domain"
	"base-gin/domain/dao"
	"time"
)

type AuthorCreateReq struct {
	Fullname  string             `json:"fullname" binding:"required,min=2,max=56"`
	Gender    *domain.TypeGender `json:"gender"`
	BirthDate *time.Time         `json:"birth_date"`
}

func (o *AuthorCreateReq) ToEntity() dao.Author {
	var item dao.Author
	item.Fullname = o.Fullname
	item.Gender = o.Gender
	item.BirthDate = o.BirthDate

	return item
}

type AuthorResp struct {
	ID        uint               `json:"id"`
	Fullname  string             `json:"fullname"`
	Gender    *domain.TypeGender `json:"gender"`
	BirthDate *time.Time         `json:"birth_date"`
}

func (o *AuthorResp) FromEntity(item *dao.Author) {
	o.ID = item.ID
	o.Fullname = item.Fullname
	o.Gender = item.Gender
	o.BirthDate = item.BirthDate
}

type AuthorUpdateReq struct {
	ID        uint               `json:"-"`
	Fullname  string             `json:"fullname" binding:"required,min=2,max=56"`
	Gender    *domain.TypeGender `json:"gender"`
	BirthDate *time.Time         `json:"birth_date"`
}