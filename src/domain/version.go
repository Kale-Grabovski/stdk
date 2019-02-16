package domain

import "time"

type Version struct {
	Id        int
	ModuleId  int
	IsActive  bool
	Settings  string `form:"settings" binding:"required,max=4000"`
	Filename  string
	Hash      string
	CreatedAt time.Time
}
