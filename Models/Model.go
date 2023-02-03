package Models

import (
	_ "gorm.io/driver/sqlite"
)

type Song struct {
	Chapter    string
	Sequence   int
	Name       string
	Difficulty int
}
