package controllers

import (
	"backend/initialazers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
)

/*
	type bodyBook struct {
		Title           string
		Author          string
		PublicationYear models.Pubyear
		Isbn            string
		NumberInShelf   int
		NumberBorrowed  int
	}
*/
func BookGetAll(c *gin.Context) {
	var books []models.Book
	result := initialazers.DB.Find(&books)
	if result.Error != nil {
		log.Fatal("could not retrieve objects")
		return
	}

	c.JSON(200, gin.H{
		"Post": books,
	})

}

func BookGet(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	initialazers.DB.First(&book, id)
	c.JSON(200, gin.H{
		"Post": book,
	})
}

func BookCreate(c *gin.Context) {

	var body models.Book

	err := c.Bind(&body)
	if err != nil {
		log.Fatal("could not bind")
		return
	}
	book := models.Book{
		Title:           body.Title,
		Author:          body.Author,
		PublicationYear: body.PublicationYear,
		Isbn:            body.Isbn,
		NumberInShelf:   body.NumberInShelf,
		NumberBorrowed:  body.NumberBorrowed,
	}

	result := initialazers.DB.Create(&book)

	if result.Error != nil {
		log.Fatal("could not create post")
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"book": book,
	})

}

func BookUpdate(c *gin.Context) {
	var book, body models.Book

	id := c.Param("BookId")
	initialazers.DB.First(&book, id)

	err := c.Bind(&body)
	if err != nil {
		log.Fatal("could not bind")
		return
	}
	book.Title = body.Title
	book.Title = body.Title
	book.Author = body.Author
	book.PublicationYear = body.PublicationYear
	book.Isbn = body.Isbn
	book.NumberInShelf = body.NumberInShelf
	book.NumberBorrowed = body.NumberBorrowed
	initialazers.DB.Save(&book)
}

/*func BookUpdate2(c *gin.Context) {//todo : use this one instead
	var book models.Book
	var bodyBook struct {
		Title           string
		Author          string
		PublicationYear pgtype.Timestamp
		Isbn            string
		NumberInShelf   int
		NumberBorrowed  int
	}
	id := c.Param("id")
	initialazers.DB.First(&book, id)

	err := c.Bind(&bodyBook)

	initialazers.DB.Model(&book).Updates(models.Book{
		Title: bodyBook.Title,
		Author: bodyBook.Author,
		PublicationYear : bodyBook.PublicationYear,
		Isbn            : bodyBook.Isbn,
		NumberInShelf   : bodyBook.NumberInShelf,
		NumberBorrowed  : bodyBook.NumberBorrowed,
	})
	if err != nil {
		log.Fatal("could not bind")
		return
	}

}*/

/*func PostDelete(c *gin.Context) {
	var book models.Book
	id := c.Param("bookid")
	initialazers.DB.First(&book, id)
	initialazers.DB.Delete(&book)
}*/

func BookDelete(c *gin.Context) {
	id := c.Param("id")
	initialazers.DB.Delete(&models.Book{}, id)
}
