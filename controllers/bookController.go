package controllers

import (
	"backend/initialazers"
	"backend/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BookGetAll(c *gin.Context) {
	var books []models.Book
	result := initialazers.DB.Find(&books)
	if result.Error != nil {
		initialazers.Logger.Print("could not retrieve objects")
		c.JSON(400, gin.H{
			"Error": "server might be down",
		})
		return
	}

	c.JSON(200, books)

}

func BookGet(c *gin.Context) {
	var book models.Book
	id := c.Param("id")
	initialazers.DB.First(&book, id)
	if book.BookId == 0 {
		c.JSON(400, gin.H{
			"Error": "book does not exist",
		})
		return
	}
	c.JSON(200, book)
}

func BookCreate(c *gin.Context) {

	var body models.Book
	var validationErrors string
	err := c.Bind(&body)
	if err != nil {
		initialazers.Logger.Print("could not bind" + err.Error())
		return
	}

	validationErrors = validateBook(body)
	if validationErrors != "" {
		c.JSON(400, gin.H{
			"Error": validationErrors,
		})
		return
	}
	result := initialazers.DB.Create(&body)

	if result.Error != nil {
		initialazers.Logger.Print("could not create post")
		c.Status(400)
	}
	c.JSON(200, gin.H{
		"book": body,
	})

}

func BookUpdate(c *gin.Context) {
	validationErrors := ""
	var book, body models.Book
	id := c.Param("id")
	if id == "" {
		initialazers.Logger.Print("no id")
		return
	}
	initialazers.DB.First(&book, id)
	if book.BookId == 0 {
		c.JSON(400, gin.H{
			"Error": "the book you are trying to edit does not exist",
		})
		return
	}

	err := c.Bind(&body)
	if err != nil {
		initialazers.Logger.Print("could not bind")
		return
	}

	validationErrors = validateBook(body)
	if validationErrors != "" {
		c.JSON(400, gin.H{
			"Error": validationErrors,
		})
		return
	}

	initialazers.DB.Model(&book).Updates(body).Where("book_id = ?", id)

	c.JSON(200, gin.H{
		"book": book,
	})

}

func BookDelete(c *gin.Context) {
	id := c.Param("id")
	var count int64
	err := initialazers.DB.Model(&models.Book{}).Where("book_id = ?", id).Count(&count).Error
	if err != nil {
		initialazers.Logger.Print("could no execute query" + err.Error())
	}
	if count == 0 {
		c.JSON(400, gin.H{
			"Error": "did not find book",
		})
		return
	}
	initialazers.DB.Delete(&models.Book{}, id)
	c.JSON(400, gin.H{
		"Success": "Book was deleted successfully",
	})
}

func validateBook(body models.Book) string {
	validationErrorsToReturn := ""
	err := initialazers.Validate.Struct(body)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, validationErr := range validationErrors {
			initialazers.Logger.Printf("%v", validationErr.Value())
			if validationErr.Tag() == "uniqueIsbn" {
				validationErrorsToReturn += fmt.Sprint("isbn must be unique. \n")
			} else if validationErr.Tag() == "tenLetters" {
				validationErrorsToReturn += fmt.Sprint("isbn must have 10 letters. \n")
			} else if validationErr.Tag() == "required" {
				validationErrorsToReturn += fmt.Sprint(validationErr.Field() + " is required. \n")
			} else if validationErr.Tag() == "min" {
				validationErrorsToReturn += "\n" + validationErr.Field() + " must be greater than 0. \n"
			}
		}
		fmt.Println(validationErrorsToReturn)
		return validationErrorsToReturn
	}

	return validationErrorsToReturn
}
