package model

import (
	"gorm.io/gorm"
)

// Example struct
type Example struct {
	gorm.Model
	Name string
	Age  string
}
