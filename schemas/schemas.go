package schemas

import "gorm.io/gorm"

type Opening struct {
	gorm.Model // tells go that this struct has a bunch of props that are useful to gorm
	Role       string
	Company    string
	Location   string
	Remote     bool
	Link       string
	Salary     int64
}
