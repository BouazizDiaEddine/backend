package initialazers

import (
	"backend/models"
	"github.com/go-playground/validator/v10"
	"unicode/utf8"
)

var Validate = validator.New()

func uniqueIsbn(fl validator.FieldLevel) bool {
	isbnEntered := fl.Field().String()
	var count int64
	var bookId string
	err := DB.Model(&models.Book{}).Select("book_id").Where("isbn = ?", isbnEntered).Scan(&bookId).Error
	if err != nil {
		Logger.Print("could no execute query" + err.Error())
	}
	err = DB.Model(&models.Book{}).Where("isbn = ? and book_id != ?", isbnEntered, bookId).Count(&count).Error
	if err != nil {
		Logger.Print("could no execute query" + err.Error())
	}
	if count == 0 {
		return true
	} else {
		return false
	}
}

func tenLetters(fl validator.FieldLevel) bool {
	isbnEnteredLetters := utf8.RuneCountInString(fl.Field().String())
	if isbnEnteredLetters == 10 {
		return true
	} else {
		return false
	}
}

func InitValidator() {
	Validate.RegisterValidation("uniqueIsbn", uniqueIsbn)
	Validate.RegisterValidation("tenLetters", tenLetters)
}
