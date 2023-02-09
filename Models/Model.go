package Models

import (
	_ "gorm.io/driver/sqlite"
)

type Song struct {
	Name       string
	Chapter    string
	Sequence   int
	Difficulty int
}
