package controllers

import (
	"backend/initialazers"
	"backend/models"
	"github.com/gin-gonic/gin"
	"log"
)

func BookGetAll(c *gin.Context) {
	var books []models.Book
	result := initialazers.DB.Find(&books)
	if result.Error != nil {
		log.Fatal("could not retrieve objects")
		return
	}

	c.JSON(200, books)

}

func BookGet(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	initialazers.DB.First(&book, id)
	c.JSON(200, book)
}

func BookCreate(c *gin.Context) {

	var body models.Book

	err := c.Bind(&body)
	if err != nil {
		log.Println("could not bind" + err.Error())
		return
	}
	err = models.Validate.Struct(body)
	if err != nil {
		log.Println("could not validate" + err.Error())
		return
	}

	result := initialazers.DB.Create(&body)

	if result.Error != nil {
		log.Println("could not create post")
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"book": body,
	})

}

func BookUpdate(c *gin.Context) {
	var book, body models.Book
	id := c.Param("id")
	if id == "" {
		log.Println("no id")
		return
	}
	dbErr := initialazers.DB.First(&book, id)
	if dbErr != nil {
		log.Println("could not validate")
		return
	}

	err := c.Bind(&body)
	if err != nil {
		log.Println("could not bind")
		return
	}

	err = models.Validate.Struct(body)
	if err != nil {
		log.Println("could not validate")
		return
	}
	dbErr = initialazers.DB.Model(&book).Updates(body).Where("book_id = ?", id)
	if dbErr != nil {
		log.Println("could not update")
		return
	}

	c.JSON(200, gin.H{
		"book": book,
	})

}

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
