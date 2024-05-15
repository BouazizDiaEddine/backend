package models

type Operation string

const (
	Redirected Operation = "redirected"
	Canonical  Operation = "canonical"
	All        Operation = "all"
)

type Url struct {
	UrlId      uint      `gorm:"primaryKey"`
	UrlGiven   string    `gorm:"not null"`
	Operation  Operation `gorm:"not null"`
	UrlTreated string    `gorm:"not null"`
}
