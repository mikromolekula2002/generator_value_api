package models

import "time"

type GeneratedValues struct {
	Id          uint   `gorm:"primaryKey"`
	RequestId   string `gorm:"not null;unique"`
	RandomValue string `gorm:"not null;unique"`
	ValueType   string `gorm:"not null"`
	Length      int
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

type UsersRequests struct {
	Id           uint      `gorm:"primaryKey"`
	RequestId    string    `gorm:"not null"`
	RandomValue  string    `gorm:"not null"`
	UserAgent    string    `gorm:"not null"`
	Url          string    `gorm:"not null"`
	Method       string    `gorm:"not null"`
	RequestCount int       `gorm:"default 1"`
	CreatedAt    time.Time `gorm:"autoCreateTime"`
}

type JsonResponse struct {
	HttpCode       int
	RequestId      string
	GeneratedValue string
	TypeGenValue   string
	LengthGenValue string
	ErrorMsg       string
}
