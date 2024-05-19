package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Book struct {
	BookId          uint    `gorm:"primaryKey"`
	Title           string  `gorm:"not null" validate:"required"`
	Author          string  `validate:"required"`
	PublicationYear PubYear `gorm:"type:jsonb"`
	Isbn            string  `gorm:"size:10;not null;unique" validate:"required,uniqueIsbn,tenLetters"`
	NumberInShelf   int     `validate:"min=0"`
	NumberBorrowed  int     `validate:"min=0"`
}

type PubYear struct {
	Year  uint
	Month string
}

// Implement the Scanner interface for PubYear
func (py *PubYear) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed to scan PubYear, invalid data type")
	}
	return json.Unmarshal(data, py)
}

// Implement the Valuer interface for PubYear
func (py PubYear) Value() (driver.Value, error) {
	if py.Year == 0 && py.Month == "" {
		return nil, nil
	}
	return json.Marshal(py)
}
