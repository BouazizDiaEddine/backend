package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
)

type Book struct {
	BookId          uint   `gorm:"primaryKey"`
	Title           string `gorm:"not null"`
	Author          string
	PublicationYear Pubyear `gorm:"type:jsonb"`
	Isbn            string  `gorm:"size:10;not null;unique"`
	NumberInShelf   int     //`gorm:"size>0;not null"`
	NumberBorrowed  int     //`gorm:"size>0;not null"`
}

type Pubyear struct {
	Year  uint
	Month string
}

// Implement the Scanner interface for Pubyear
func (py *Pubyear) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed to scan Pubyear, invalid data type")
	}
	return json.Unmarshal(data, py)
}

// Implement the Valuer interface for Pubyear
func (py Pubyear) Value() (driver.Value, error) {
	if py.Year == 0 && py.Month == "" {
		return nil, nil
	}
	return json.Marshal(py)
}

func (py *Book) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	data, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("Failed to scan Pubyear, invalid data type")
	}
	return json.Unmarshal(data, py)
}
