package models

import (
	"github.com/jackc/pgx/v5/pgtype"
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	BookId          string `gorm:"not null"`
	Title           string `gorm:"not null"`
	Author          string
	PublicationYear pgtype.Timestamp
	Isbn            string `gorm:"size:10;not null"`
	NumberInShelf   int    //`gorm:"size>0;not null"`
	numberBorrowed  int    //`gorm:"size>0;not null"`
}
