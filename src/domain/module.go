package domain

type Module struct {
	Id   int    `uri:"id" binding:"max=435424233"`
	Name string `form:"name" binding:"required,max=10"`
}
