package dto

import "base-gin/domain/dao"

type BookCreateReq struct {
	Title       string `json:"title" binding:"required,min=2,max=56"`
	Subtitle    string `json:"subtitle,omitempty" binding:"max=46"`
	AuthorID    uint   `json:"author_id" binding:"required"`
	PublisherID uint   `json:"publisher_id" binding:"required"`
}

func (o *BookCreateReq) ToEntity() dao.Book {
	var item dao.Book
	item.Title = o.Title
	item.Subtitle = &o.Subtitle
	item.AuthorID = o.AuthorID
	item.PublisherID = o.PublisherID

	return item
}

type BookResp struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Subtitle    string `json:"subtitle,omitempty"`
	AuthorID    uint   `json:"author_id"`
	PublisherID uint   `json:"publisher_id"`
}

func (o *BookResp) FromEntity(item *dao.Book) {
	o.ID = item.ID
	o.Title = item.Title
	o.Subtitle = *item.Subtitle
	o.AuthorID = item.AuthorID
	o.PublisherID = item.PublisherID
}

type BookUpdateReq struct {
	ID          uint   `json:"-"`
	Title       string `json:"title" binding:"required,min=2,max=56"`
	Subtitle    string `json:"subtitle,omitempty" binding:"max=46"`
	AuthorID    uint   `json:"author_id" binding:"required"`
	PublisherID uint   `json:"publisher_id" binding:"required"`
}
