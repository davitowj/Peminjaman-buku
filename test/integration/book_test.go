package integration_test

import (
	"base-gin/domain/dao"
	"base-gin/domain/dto"
	"base-gin/server"
	"base-gin/util"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func createBook() dao.Book {
	book := dao.Book{
		Title:       util.RandomStringAlpha(10),
		Subtitle:    nil,
		AuthorID:    createAuthor().ID,     // Menggunakan Author ID dari fungsi createAuthor
		PublisherID: createPublisher().ID,  // Menggunakan Publisher ID dari fungsi createPublisher
	}
	_ = bookRepo.Create(&book)

	return book
}

func createPublisher() dao.Publisher {
	publisher := dao.Publisher{
		Name: util.RandomStringAlpha(8),
	}
	_ = publisherRepo.Create(&publisher)
	
	return publisher
}

func TestBook_Create_Success(t *testing.T) {
	params := dto.BookCreateReq{
		Title:       util.RandomStringAlpha(15),
		Subtitle:    util.RandomStringAlpha(10),
		AuthorID:    createAuthor().ID,
		PublisherID: createPublisher().ID,
	}

	w := doTest(
		"POST",
		server.RootBook,
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 201, w.Code)
}

func TestBook_Update_Success(t *testing.T) {
	b := createBook()

	params := dto.BookUpdateReq{
		ID:          b.ID,
		Title:       util.RandomStringAlpha(15),
		Subtitle:    util.RandomStringAlpha(10),
		AuthorID:    b.AuthorID,
		PublisherID: b.PublisherID,
	}

	w := doTest(
		"PUT",
		fmt.Sprintf("%s/%d", server.RootBook, b.ID),
		params,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := bookRepo.GetByID(b.ID)
	assert.Equal(t, params.Title, item.Title)
	assert.Equal(t, params.Subtitle, *item.Subtitle)
}

func TestBook_Delete_Success(t *testing.T) {
	b := createBook()
	_ = bookRepo.Create(&b)

	w := doTest(
		"DELETE",
		fmt.Sprintf("%s/%d", server.RootBook, b.ID),
		nil,
		createAuthAccessToken(dummyAdmin.Account.Username),
	)
	assert.Equal(t, 200, w.Code)

	item, _ := bookRepo.GetByID(b.ID)
	assert.Nil(t, item)
}

func TestBook_GetList_Success(t *testing.T) {
    b1 := createBook()
    _ = bookRepo.Create(&b1)

    b2 := createBook()
    _ = bookRepo.Create(&b2)

    w := doTest(
        "GET",
        server.RootBook,
        nil,
        "",
    )
    assert.Equal(t, 200, w.Code)

    body := w.Body.String()
    assert.Contains(t, body, b1.Title)
    if b1.Subtitle != nil {
        assert.Contains(t, body, *b1.Subtitle)
    }
    assert.Contains(t, body, b2.Title)
    if b2.Subtitle != nil {
        assert.Contains(t, body, *b2.Subtitle)
    }

    // Tes dengan filter pencarian (contoh menggunakan Title dari b1)
    w = doTest(
        "GET",
        server.RootBook+"?q="+b1.Title,
        nil,
        "",
    )
    assert.Equal(t, 200, w.Code)

    body = w.Body.String()
    assert.Contains(t, body, b1.Title)
    assert.NotContains(t, body, b2.Title)
}

func TestBook_GetDetail_Success(t *testing.T) {
	b := createBook()
	_ = bookRepo.Create(&b)

	w := doTest(
		"GET",
		fmt.Sprintf("%s/%d", server.RootBook, b.ID),
		nil,
		"",
	)
	assert.Equal(t, 200, w.Code)

	body := w.Body.String()
	assert.Contains(t, body, b.Title)
}
