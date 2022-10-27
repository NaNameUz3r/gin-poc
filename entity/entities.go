package entity

type Video struct {
	Title       string `json:"title" binding:"min=2,max=20" validate:"no-bad-word"`
	Description string `json:"description" binding:"max=150"`
	URL         string `json:"url" binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       int8   `json:"age" binding:"gte=8,lte=130"`
	Email     string `json:"email" binding:"required,email"`
}
