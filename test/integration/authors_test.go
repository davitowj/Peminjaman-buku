package integration_test

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"time"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createAuthor() dao.Author {
	o := dao.Author{
		Fullname: util.RandomStringAlpha(6),
		Gender: nil,
		BirthDate: func() *time.Time {
			t := time.Now()
			return &t
		}(),
	}
	_ = authorRepo.Create(&o)

	return o
}

func TestAuthor_Create_Success(t *testing.T) {
	params := dto.AuthorCreateReq{
		Fullname: util.RandomStringAlpha(10),
		Gender:   nil, 
		BirthDate: func() *time.Time {
			t := time.Now()
			return &t
		}(),
	}

	w := doTest(
		"POST",
		server.RootAuthor,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestAuthor_Update_Success(t *testing.T) {
	
	o := createAuthor()

	params := dto.AuthorUpdateReq{
		Fullname: util.RandomStringAlpha(12),
		Gender:   nil,
		BirthDate: func() *time.Time {
			t := time.Now().AddDate(-1, 0, 0)
			return &t
		}(),
	}

	w := doTest(
		"PUT",
		fmt.Sprintf("%s/%d", server.RootAuthor, o.ID),
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)

	assert.Equal(t, 200, w.Code)

	item, _ := authorRepo.GetByID(o.ID)
	assert.Equal(t, params.Fullname, item.Fullname)
	assert.Equal(t, params.Gender, item.Gender)
	assert.Equal(t, false, item.DeletedAt.Valid)
}

func TestAuthor_Delete_Success(t *testing.T) {
	o := createAuthor()
	_ = authorRepo.Create(&o)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootAuthor, o.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := authorRepo.GetByID(o.ID)
	assert.Nil(t, item)
}

func TestAuthor_GetList_Success(t *testing.T) {
	o1 := createAuthor()
	_ = authorRepo.Create(&o1)

	o2 := createAuthor()
	_ = authorRepo.Create(&o2)

	w := doTest(
		"GET",
		server.RootAuthor,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, o1.Fullname)
	assert.Contains(t, body, o2.Fullname)

	w = doTest(
		"GET",
		server.RootAuthor+"?q="+o1.Fullname,
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body = w.Body.String()
	assert.Contains(t, body, o1.Fullname)
	assert.NotContains(t, body, o2.Fullname)
}

func TestAuthor_GetDetail_Success(t *testing.T) {
	o := createAuthor()
	_ = authorRepo.Create(&o)

	w := doTest(
		"GET",
		fmt.Sprintf("%s/%d", server.RootAuthor, o.ID),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, o.Fullname)
}